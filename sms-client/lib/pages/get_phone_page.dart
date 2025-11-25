import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:intl/intl.dart';
import '../core/api_client.dart';
import '../l10n/app_localizations.dart';
import '../models/assigned_phone.dart';
import '../models/business_type.dart';
import '../models/phone_assignment_result.dart';

class GetPhonePage extends StatefulWidget {
  const GetPhonePage({super.key});

  @override
  State<GetPhonePage> createState() => _GetPhonePageState();
}

class _GetPhonePageState extends State<GetPhonePage> {
  final ApiClient _apiClient = ApiClient();

  List<BusinessType> _businessTypes = [];
  BusinessType? _selectedBusinessType;
  String _selectedCardType = 'physical';
  int _requestedCount = 1;
  List<PhoneAssignmentResult> _assignmentResults = []; // 改为列表，保存所有分配结果
  bool _isLoading = false;
  bool _isLoadingTypes = false;
  String? _error;

  @override
  void initState() {
    super.initState();
    _loadBusinessTypes();
  }

  Future<void> _loadBusinessTypes() async {
    setState(() {
      _isLoadingTypes = true;
      _error = null;
    });

    try {
      final response = await _apiClient.getBusinessTypes();
      if (response.success && response.data != null) {
        setState(() {
          _businessTypes = response.data!;
          if (_businessTypes.isNotEmpty) {
            _selectedBusinessType = _businessTypes.first;
          } else {
            // 数据为空时的友好提示
            _error = '暂无可用的业务类型，请联系管理员配置';
          }
        });
      } else {
        setState(() {
          _error = response.message ?? '获取业务类型失败';
        });
      }
    } catch (e) {
      setState(() {
        _error = '网络错误: $e';
      });
    } finally {
      setState(() {
        _isLoadingTypes = false;
      });
    }
  }

  Future<void> _assignPhone() async {
    if (_selectedBusinessType == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text(AppLocalizations.of(context)!.businessTypeRequired),
        ),
      );
      return;
    }

    setState(() {
      _isLoading = true;
      _error = null;
    });

    try {
      final response = await _apiClient.assignPhone(
        businessType: _selectedBusinessType!.code,
        cardType: _selectedCardType,
        count: _requestedCount,
      );

      if (response.success && response.data != null) {
        setState(() {
          // 将新结果添加到列表顶部，而不是替换
          _assignmentResults.insert(0, response.data!);
        });
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(AppLocalizations.of(context)!.assignSuccess),
              backgroundColor: Colors.green,
            ),
          );
        }
      } else {
        setState(() {
          _error = response.message;
        });
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(response.message ?? 'Error'),
              backgroundColor: Colors.red,
            ),
          );
        }
      }
    } catch (e) {
      setState(() {
        _error = 'Network error: $e';
      });
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('Network error: $e'),
            backgroundColor: Colors.red,
          ),
        );
      }
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  void _copyPhone(String phoneNumber) {
    Clipboard.setData(ClipboardData(text: phoneNumber));
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(
        content: Text(AppLocalizations.of(context)!.copied),
        duration: const Duration(seconds: 1),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    return Scaffold(
      appBar: AppBar(title: Text(l10n.getPhoneTitle)),
      body:
          _isLoadingTypes
              ? Center(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    const CircularProgressIndicator(),
                    const SizedBox(height: 16),
                    Text(l10n.loadingBusinessTypes),
                  ],
                ),
              )
              : _error != null && _businessTypes.isEmpty
              ? Center(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Icon(Icons.error_outline, size: 64, color: Colors.grey),
                    const SizedBox(height: 16),
                    Text(_error!, style: TextStyle(color: Colors.grey)),
                    const SizedBox(height: 16),
                    ElevatedButton.icon(
                      onPressed: _loadBusinessTypes,
                      icon: const Icon(Icons.refresh),
                      label: Text(l10n.retryGetCode),
                    ),
                  ],
                ),
              )
              : SingleChildScrollView(
                padding: const EdgeInsets.all(16),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.stretch,
                  children: [
                    // Business Type Selection
                    Card(
                      child: Padding(
                        padding: const EdgeInsets.all(16),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              l10n.selectBusinessType,
                              style: Theme.of(context).textTheme.titleMedium,
                            ),
                            const SizedBox(height: 12),
                            if (_businessTypes.isEmpty)
                              Text(
                                l10n.noBusinessTypes,
                                style: TextStyle(color: Colors.grey),
                              )
                            else
                              DropdownButtonFormField<BusinessType>(
                                value: _selectedBusinessType,
                                decoration: InputDecoration(
                                  border: OutlineInputBorder(),
                                  contentPadding: const EdgeInsets.symmetric(
                                    horizontal: 12,
                                    vertical: 8,
                                  ),
                                ),
                                items:
                                    _businessTypes.map((type) {
                                      return DropdownMenuItem(
                                        value: type,
                                        child: Text(type.name),
                                      );
                                    }).toList(),
                                onChanged: (value) {
                                  setState(() {
                                    _selectedBusinessType = value;
                                  });
                                },
                              ),
                          ],
                        ),
                      ),
                    ),
                    const SizedBox(height: 16),

                    // Card Type Selection
                    Card(
                      child: Padding(
                        padding: const EdgeInsets.all(16),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              l10n.selectCardType,
                              style: Theme.of(context).textTheme.titleMedium,
                            ),
                            const SizedBox(height: 12),
                            Row(
                              children: [
                                Expanded(
                                  child: RadioListTile<String>(
                                    title: Text(l10n.physicalCard),
                                    value: 'physical',
                                    groupValue: _selectedCardType,
                                    onChanged: (value) {
                                      setState(() {
                                        _selectedCardType = value!;
                                      });
                                    },
                                    contentPadding: EdgeInsets.zero,
                                  ),
                                ),
                                Expanded(
                                  child: RadioListTile<String>(
                                    title: Text(l10n.virtualCard),
                                    value: 'virtual',
                                    groupValue: _selectedCardType,
                                    onChanged: (value) {
                                      setState(() {
                                        _selectedCardType = value!;
                                      });
                                    },
                                    contentPadding: EdgeInsets.zero,
                                  ),
                                ),
                              ],
                            ),
                          ],
                        ),
                      ),
                    ),
                    const SizedBox(height: 16),

                    // Count Selection
                    Card(
                      child: Padding(
                        padding: const EdgeInsets.all(16),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              l10n.requestCount,
                              style: Theme.of(context).textTheme.titleMedium,
                            ),
                            const SizedBox(height: 12),
                            DropdownButtonFormField<int>(
                              value: _requestedCount,
                              decoration: const InputDecoration(
                                border: OutlineInputBorder(),
                              ),
                              items: List.generate(
                                10,
                                (index) => DropdownMenuItem(
                                  value: index + 1,
                                  child: Text('${index + 1}'),
                                ),
                              ),
                              onChanged: (value) {
                                if (value == null) return;
                                setState(() {
                                  _requestedCount = value;
                                });
                              },
                            ),
                            const SizedBox(height: 8),
                            Text(
                              l10n.requestCountHelper,
                              style: Theme.of(context).textTheme.bodySmall
                                  ?.copyWith(color: Colors.grey),
                            ),
                          ],
                        ),
                      ),
                    ),
                    const SizedBox(height: 24),

                    // Assign Button
                    FilledButton.icon(
                      onPressed:
                          _isLoading || _businessTypes.isEmpty
                              ? null
                              : _assignPhone,
                      icon:
                          _isLoading
                              ? const SizedBox(
                                width: 20,
                                height: 20,
                                child: CircularProgressIndicator(
                                  strokeWidth: 2,
                                  color: Colors.white,
                                ),
                              )
                              : const Icon(Icons.phone_android),
                      label: Text(
                        _isLoading ? l10n.loading : l10n.assignNow,
                        style: const TextStyle(fontSize: 16),
                      ),
                      style: FilledButton.styleFrom(
                        padding: const EdgeInsets.symmetric(vertical: 16),
                      ),
                    ),
                    const SizedBox(height: 24),

                    // 显示所有分配结果（最新的在上面）
                    if (_assignmentResults.isNotEmpty)
                      ...(_assignmentResults.map((result) => Column(
                        children: [
                          _buildSummaryCard(context, l10n, result),
                          const SizedBox(height: 16),
                          if (result.phones.isEmpty)
                            Card(
                              child: Padding(
                                padding: const EdgeInsets.all(16),
                                child: Row(
                                  children: [
                                    const Icon(
                                      Icons.info_outline,
                                      color: Colors.grey,
                                    ),
                                    const SizedBox(width: 8),
                                    Expanded(
                                      child: Text(
                                        l10n.noData,
                                        style: const TextStyle(color: Colors.grey),
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                            )
                          else
                            Column(
                              children:
                                  result.phones
                                      .map(
                                        (phone) =>
                                            _buildPhoneCard(context, phone, l10n),
                                      )
                                      .toList(),
                            ),
                          const SizedBox(height: 24),
                          const Divider(thickness: 2),
                          const SizedBox(height: 16),
                        ],
                      ))),
                  ],
                ),
              ),
    );
  }

  Widget _buildSummaryCard(
    BuildContext context,
    AppLocalizations l10n,
    PhoneAssignmentResult result,
  ) {
    return Card(
      color: Theme.of(context).colorScheme.primaryContainer,
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Icon(
                  Icons.check_circle,
                  color: Theme.of(context).colorScheme.primary,
                ),
                const SizedBox(width: 8),
                Text(
                  l10n.assignmentResult,
                  style: Theme.of(context).textTheme.titleMedium?.copyWith(
                    color: Theme.of(context).colorScheme.primary,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
            const SizedBox(height: 16),
            _buildSummaryRow(
              l10n.totalCostLabel,
              _formatAmount(result.totalCost),
            ),
            _buildSummaryRow(
              l10n.remainingBalanceLabel,
              _formatAmount(result.remainingBalance),
            ),
            const Divider(height: 24),
            Row(
              children: [
                Expanded(
                  child: _buildSummaryRow(
                    l10n.successCountLabel,
                    result.successCount.toString(),
                  ),
                ),
                Expanded(
                  child: _buildSummaryRow(
                    l10n.failedCountLabel,
                    result.failedCount.toString(),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildSummaryRow(String label, String value) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 4),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(label, style: const TextStyle(color: Colors.black54)),
          Text(value, style: const TextStyle(fontWeight: FontWeight.bold)),
        ],
      ),
    );
  }

  Widget _buildPhoneCard(
    BuildContext context,
    AssignedPhone phone,
    AppLocalizations l10n,
  ) {
    return Card(
      margin: const EdgeInsets.symmetric(vertical: 6),
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Expanded(
                  child: Text(
                    phone.phoneNumber,
                    style: Theme.of(context).textTheme.titleLarge?.copyWith(
                      fontWeight: FontWeight.bold,
                      letterSpacing: 1.5,
                    ),
                  ),
                ),
                IconButton.filled(
                  onPressed: () => _copyPhone(phone.phoneNumber),
                  icon: const Icon(Icons.copy),
                  tooltip: l10n.copyPhone,
                ),
              ],
            ),
            const SizedBox(height: 12),
            Wrap(
              spacing: 12,
              runSpacing: 8,
              children: [
                _buildInfoChip(l10n.countryCode, phone.countryCode),
                _buildInfoChip(l10n.providerLabel, phone.providerId),
                _buildInfoChip(
                  l10n.validUntilLabel,
                  _formatDate(phone.validUntil),
                ),
                _buildInfoChip(l10n.cost, _formatAmount(phone.cost)),
              ],
            ),
            const SizedBox(height: 12),
            // 获取验证码按钮
            SizedBox(
              width: double.infinity,
              child: ElevatedButton.icon(
                onPressed: () => _fetchCodeForPhone(phone.phoneNumber),
                icon: const Icon(Icons.sms, size: 18),
                label: const Text('获取验证码'),
                style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.orange,
                  foregroundColor: Colors.white,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  Future<void> _fetchCodeForPhone(String phoneNumber) async {
    try {
      final response = await _apiClient.getVerificationCodes(
        phoneNumbers: [phoneNumber],
      );

      if (response.success && response.data != null) {
        final result = response.data!;

        if (result.codes.isNotEmpty) {
          final codeEntry = result.codes.first;

          if (mounted) {
            if (codeEntry.status == 'success') {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text('验证码: ${codeEntry.code}'),
                  backgroundColor: Colors.green,
                  duration: const Duration(seconds: 5),
                  action: SnackBarAction(
                    label: '复制',
                    textColor: Colors.white,
                    onPressed: () {
                      Clipboard.setData(ClipboardData(text: codeEntry.code));
                    },
                  ),
                ),
              );
            } else if (codeEntry.status == 'pending') {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text('验证码还未收到，请稍后再试'),
                  backgroundColor: Colors.orange,
                ),
              );
            } else {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text(codeEntry.message),
                  backgroundColor: Colors.red,
                ),
              );
            }
          }
        }
      } else {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(response.message ?? '获取失败'),
              backgroundColor: Colors.red,
            ),
          );
        }
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('网络错误: $e'),
            backgroundColor: Colors.red,
          ),
        );
      }
    }
  }

  Widget _buildInfoChip(String label, String value) {
    return Chip(
      label: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            label,
            style: const TextStyle(fontSize: 12, color: Colors.black54),
          ),
          Text(
            value.isEmpty ? '--' : value,
            style: const TextStyle(fontWeight: FontWeight.bold),
          ),
        ],
      ),
      padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 12),
    );
  }

  String _formatAmount(double value) => value.toStringAsFixed(4);

  String _formatDate(DateTime? date) {
    if (date == null) return '--';
    return DateFormat('yyyy-MM-dd HH:mm:ss').format(date.toLocal());
  }
}
