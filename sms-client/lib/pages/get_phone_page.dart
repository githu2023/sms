import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:intl/intl.dart';
import '../core/api_client.dart';
import '../l10n/app_localizations.dart';
import '../models/assigned_phone.dart';
import '../models/assignment.dart';
import '../models/business_type.dart';
import '../models/phone_assignment_result.dart';
import '../widgets/assignment_card.dart';

class GetPhonePage extends StatefulWidget {
  const GetPhonePage({super.key});

  @override
  State<GetPhonePage> createState() => _GetPhonePageState();
}

class _GetPhonePageState extends State<GetPhonePage> {
  final ApiClient _apiClient = ApiClient();
  final ScrollController _scrollController = ScrollController();

  List<BusinessType> _businessTypes = [];
  BusinessType? _selectedBusinessType;
  String _selectedCardType = 'physical';
  int _requestedCount = 1;
  List<PhoneAssignmentResult> _assignmentResults = []; // 页面内的分配结果
  List<Assignment> _recentAssignments = []; // 最近的历史记录
  bool _isLoading = false;
  bool _isLoadingTypes = false;
  bool _isLoadingHistory = false;
  bool _hasMoreHistory = true;
  int _historyPage = 1;
  String? _error;

  @override
  void initState() {
    super.initState();
    _loadBusinessTypes();
    _loadRecentAssignments(); // 加载最近的历史记录
    _scrollController.addListener(_onScroll);
  }

  @override
  void dispose() {
    _scrollController.dispose();
    super.dispose();
  }

  void _onScroll() {
    if (_scrollController.position.pixels >=
        _scrollController.position.maxScrollExtent - 200) {
      if (!_isLoadingHistory && _hasMoreHistory) {
        _loadMoreHistory();
      }
    }
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
            _error = context.mounted 
                ? AppLocalizations.of(context)!.noAvailableBusinessTypes
                : 'No business types available';
          }
        });
      } else {
        setState(() {
          _error = response.message ?? 
              (context.mounted 
                  ? AppLocalizations.of(context)!.getBusinessTypesFailed
                  : 'Failed to get business types');
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

  Future<void> _loadRecentAssignments({bool loadMore = false}) async {
    if (_isLoadingHistory) return;

    setState(() {
      _isLoadingHistory = true;
    });

    try {
      final page = loadMore ? _historyPage + 1 : 1;
      final response = await _apiClient.getAssignments(page: page, limit: 10);

      if (response.success && response.data != null) {
        setState(() {
          if (loadMore) {
            _recentAssignments.addAll(response.data!);
            _historyPage = page;
          } else {
            _recentAssignments = response.data!;
            _historyPage = 1;
          }
          _hasMoreHistory = response.data!.length >= 10;
        });
      }
    } catch (e) {
      debugPrint('Failed to load recent assignments: $e');
    } finally {
      setState(() {
        _isLoadingHistory = false;
      });
    }
  }

  Future<void> _loadMoreHistory() async {
    await _loadRecentAssignments(loadMore: true);
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
              : RefreshIndicator(
                onRefresh: () async {
                  await _loadRecentAssignments();
                },
                child: SingleChildScrollView(
                  controller: _scrollController,
                  padding: const EdgeInsets.all(16),
                  physics: const AlwaysScrollableScrollPhysics(),
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

                    // Card Type and Count Selection (合并到一行)
                    Card(
                      child: Padding(
                        padding: const EdgeInsets.all(16),
                        child: Row(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            // 左边：卡类型选择
                            Expanded(
                              flex: 3,
                              child: Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(
                                    l10n.selectCardType,
                                    style: Theme.of(context).textTheme.titleSmall,
                                  ),
                                  const SizedBox(height: 8),
                                  Row(
                                    children: [
                                      Expanded(
                                        child: InkWell(
                                          onTap: () {
                                            setState(() {
                                              _selectedCardType = 'physical';
                                            });
                                          },
                                          child: Container(
                                            padding: const EdgeInsets.symmetric(vertical: 8),
                                            child: Row(
                                              mainAxisSize: MainAxisSize.min,
                                              children: [
                                                Radio<String>(
                                                  value: 'physical',
                                                  groupValue: _selectedCardType,
                                                  onChanged: (value) {
                                                    setState(() {
                                                      _selectedCardType = value!;
                                                    });
                                                  },
                                                  materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
                                                ),
                                                Flexible(
                                                  child: Text(
                                                    l10n.physicalCard,
                                                    style: const TextStyle(fontSize: 14),
                                                  ),
                                                ),
                                              ],
                                            ),
                                          ),
                                        ),
                                      ),
                                      const SizedBox(width: 8),
                                      Expanded(
                                        child: InkWell(
                                          onTap: () {
                                            setState(() {
                                              _selectedCardType = 'virtual';
                                            });
                                          },
                                          child: Container(
                                            padding: const EdgeInsets.symmetric(vertical: 8),
                                            child: Row(
                                              mainAxisSize: MainAxisSize.min,
                                              children: [
                                                Radio<String>(
                                                  value: 'virtual',
                                                  groupValue: _selectedCardType,
                                                  onChanged: (value) {
                                                    setState(() {
                                                      _selectedCardType = value!;
                                                    });
                                                  },
                                                  materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
                                                ),
                                                Flexible(
                                                  child: Text(
                                                    l10n.virtualCard,
                                                    style: const TextStyle(fontSize: 14),
                                                  ),
                                                ),
                                              ],
                                            ),
                                          ),
                                        ),
                                      ),
                                    ],
                                  ),
                                ],
                              ),
                            ),
                            const SizedBox(width: 16),
                            // 右边：数量选择
                            Expanded(
                              flex: 2,
                              child: Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(
                                    l10n.requestCount,
                                    style: Theme.of(context).textTheme.titleSmall,
                                  ),
                                  const SizedBox(height: 8),
                                  DropdownButtonFormField<int>(
                                    value: _requestedCount,
                                    decoration: const InputDecoration(
                                      border: OutlineInputBorder(),
                                      contentPadding: EdgeInsets.symmetric(
                                        horizontal: 12,
                                        vertical: 12,
                                      ),
                                      isDense: true,
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
                                  const SizedBox(height: 4),
                                  Text(
                                    l10n.requestCountHelper,
                                    style: Theme.of(context).textTheme.bodySmall
                                        ?.copyWith(color: Colors.grey, fontSize: 10),
                                  ),
                                ],
                              ),
                            ),
                          ],
                        ),
                      ),
                    ),
                    const SizedBox(height: 16),

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

                    // 显示所有分配结果（最新的在上面）- 优化：直接显示手机号，简化摘要信息
                    if (_assignmentResults.isNotEmpty)
                      ...(_assignmentResults.map((result) => Column(
                        children: [
                          // 简化摘要卡片，只显示关键信息
                          if (result.phones.isNotEmpty)
                            _buildCompactSummaryCard(context, l10n, result),
                          const SizedBox(height: 12),
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
                          const SizedBox(height: 16),
                          const Divider(thickness: 1),
                          const SizedBox(height: 12),
                        ],
                      ))),

                    // 分隔线
                    const SizedBox(height: 32),
                    const Divider(thickness: 2, height: 2),
                    const SizedBox(height: 24),

                    // 最近的历史记录 - 固定在底部
                    Container(
                      width: double.infinity,
                      padding: const EdgeInsets.symmetric(vertical: 16),
                      decoration: BoxDecoration(
                        color: Theme.of(context).colorScheme.surfaceContainerHighest,
                        borderRadius: BorderRadius.circular(12),
                      ),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Padding(
                            padding: const EdgeInsets.symmetric(horizontal: 16),
                            child: Row(
                              children: [
                                Icon(
                                  Icons.history,
                                  color: Theme.of(context).primaryColor,
                                  size: 24,
                                ),
                                const SizedBox(width: 8),
                                Text(
                                  l10n.recentRecords,
                                  style: Theme.of(context).textTheme.titleLarge?.copyWith(
                                    fontWeight: FontWeight.bold,
                                  ),
                                ),
                              ],
                            ),
                          ),
                          const SizedBox(height: 8),
                          Padding(
                            padding: const EdgeInsets.symmetric(horizontal: 16),
                            child: Text(
                              l10n.pullDownToRefresh,
                              style: Theme.of(context).textTheme.bodySmall?.copyWith(
                                color: Colors.grey,
                              ),
                            ),
                          ),
                          const SizedBox(height: 16),
                          if (_recentAssignments.isEmpty && !_isLoadingHistory)
                            Padding(
                              padding: const EdgeInsets.all(16),
                              child: Center(
                                child: Text(
                                  l10n.noData,
                                  style: TextStyle(color: Colors.grey),
                                ),
                              ),
                            )
                          else
                            ...(_recentAssignments.map((assignment) =>
                                AssignmentCard(
                                  assignment: assignment,
                                  onRefresh: _loadRecentAssignments,
                                ))),
                          if (_isLoadingHistory)
                            const Padding(
                              padding: EdgeInsets.all(16),
                              child: Center(child: CircularProgressIndicator()),
                            ),
                          if (!_hasMoreHistory && _recentAssignments.isNotEmpty)
                            Padding(
                              padding: const EdgeInsets.all(16),
                              child: Center(
                                child: Text(
                                  l10n.noMoreData,
                                  style: TextStyle(color: Colors.grey),
                                ),
                              ),
                            ),
                        ],
                      ),
                    ),
                    const SizedBox(height: 16),
                  ],
                ),
              ),
            ),
    );
  }


  // 紧凑的摘要卡片 - 只显示关键信息，节省空间
  Widget _buildCompactSummaryCard(
    BuildContext context,
    AppLocalizations l10n,
    PhoneAssignmentResult result,
  ) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
      decoration: BoxDecoration(
        color: Theme.of(context).colorScheme.primaryContainer.withOpacity(0.3),
        borderRadius: BorderRadius.circular(8),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: [
          _buildCompactInfoItem(
            Icons.attach_money,
            _formatAmount(result.totalCost),
            l10n.totalCostLabel,
          ),
          Container(width: 1, height: 30, color: Colors.grey[300]),
          _buildCompactInfoItem(
            Icons.account_balance_wallet,
            _formatAmount(result.remainingBalance),
            l10n.remainingBalanceLabel,
          ),
          Container(width: 1, height: 30, color: Colors.grey[300]),
          _buildCompactInfoItem(
            Icons.check_circle,
            '${result.successCount}/${result.successCount + result.failedCount}',
            '成功',
          ),
        ],
      ),
    );
  }

  Widget _buildCompactInfoItem(IconData icon, String value, String label) {
    return Expanded(
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(icon, size: 14, color: Colors.grey[700]),
              const SizedBox(width: 4),
              Flexible(
                child: Text(
                  value,
                  style: const TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 13,
                  ),
                  overflow: TextOverflow.ellipsis,
                ),
              ),
            ],
          ),
          const SizedBox(height: 2),
          Text(
            label,
            style: TextStyle(
              fontSize: 10,
              color: Colors.grey[600],
            ),
            textAlign: TextAlign.center,
          ),
        ],
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
    // 查找该手机号的验证码（从历史记录中）
    String? verificationCode;
    for (var assignment in _recentAssignments) {
      if (assignment.phone == phone.phoneNumber && 
          assignment.code != null && 
          assignment.code!.isNotEmpty) {
        verificationCode = assignment.code;
        break;
      }
    }

    return Card(
      margin: const EdgeInsets.symmetric(vertical: 6),
      elevation: 2,
      child: Padding(
        padding: const EdgeInsets.all(12),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // 手机号和复制按钮 - 更紧凑
            Row(
              children: [
                Icon(
                  Icons.phone_android,
                  color: Theme.of(context).primaryColor,
                  size: 24,
                ),
                const SizedBox(width: 8),
                Expanded(
                  child: Text(
                    phone.phoneNumber,
                    style: const TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 18,
                      letterSpacing: 1.0,
                    ),
                  ),
                ),
                IconButton(
                  onPressed: () => _copyPhone(phone.phoneNumber),
                  icon: const Icon(Icons.copy, size: 18),
                  tooltip: l10n.copyPhone,
                  padding: EdgeInsets.zero,
                  constraints: const BoxConstraints(),
                ),
              ],
            ),
            const SizedBox(height: 10),
            
            // 信息网格 - 更紧凑
            Row(
              children: [
                Expanded(
                  child: _buildCompactPhoneInfoItem(
                    Icons.public,
                    phone.countryCode,
                    l10n.countryCode,
                  ),
                ),
                Expanded(
                  child: _buildCompactPhoneInfoItem(
                    Icons.store,
                    phone.providerId,
                    l10n.providerLabel,
                  ),
                ),
              ],
            ),
            const SizedBox(height: 8),
            Row(
              children: [
                Expanded(
                  child: _buildCompactPhoneInfoItem(
                    Icons.schedule,
                    _formatDate(phone.validUntil),
                    l10n.validUntilLabel,
                  ),
                ),
                Expanded(
                  child: _buildCompactPhoneInfoItem(
                    Icons.attach_money,
                    _formatAmount(phone.cost),
                    l10n.cost,
                  ),
                ),
              ],
            ),
            
            // 验证码显示或获取按钮 - 更紧凑
            if (verificationCode != null) ...[
              const SizedBox(height: 10),
              Container(
                width: double.infinity,
                padding: const EdgeInsets.all(10),
                decoration: BoxDecoration(
                  color: Colors.green.withOpacity(0.1),
                  borderRadius: BorderRadius.circular(8),
                  border: Border.all(color: Colors.green, width: 1.5),
                ),
                child: Row(
                  children: [
                    const Icon(Icons.verified, color: Colors.green, size: 20),
                    const SizedBox(width: 8),
                    Expanded(
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(
                            verificationCode,
                            style: const TextStyle(
                              fontSize: 20,
                              fontWeight: FontWeight.bold,
                              color: Colors.green,
                              letterSpacing: 1.5,
                            ),
                          ),
                          Text(
                            l10n.code,
                            style: const TextStyle(
                              fontSize: 10,
                              color: Colors.grey,
                            ),
                          ),
                        ],
                      ),
                    ),
                    IconButton(
                      onPressed: () {
                        Clipboard.setData(ClipboardData(text: verificationCode!));
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(
                            content: Text(l10n.copied),
                            duration: const Duration(seconds: 1),
                          ),
                        );
                      },
                      icon: const Icon(Icons.copy, size: 18),
                      padding: EdgeInsets.zero,
                      constraints: const BoxConstraints(),
                      style: IconButton.styleFrom(
                        backgroundColor: Colors.green,
                        foregroundColor: Colors.white,
                      ),
                    ),
                  ],
                ),
              ),
            ] else ...[
              const SizedBox(height: 12),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: () => _fetchCodeForNewPhone(phone.phoneNumber),
                  icon: const Icon(Icons.sms, size: 18),
                  label: Text(
                    l10n.getCode,
                    style: const TextStyle(fontSize: 14, fontWeight: FontWeight.bold),
                  ),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.orange,
                    foregroundColor: Colors.white,
                    padding: const EdgeInsets.symmetric(vertical: 12),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
              ),
            ],
          ],
        ),
      ),
    );
  }

  // 紧凑版手机号信息项 - 值在上，标签在下
  Widget _buildCompactPhoneInfoItem(IconData icon, String value, String label) {
    return Row(
      children: [
        Icon(icon, size: 14, color: Colors.grey[600]),
        const SizedBox(width: 4),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                value,
                style: const TextStyle(
                  fontSize: 12,
                  fontWeight: FontWeight.w600,
                ),
                overflow: TextOverflow.ellipsis,
              ),
              Text(
                label,
                style: TextStyle(
                  fontSize: 10,
                  color: Colors.grey[600],
                ),
                overflow: TextOverflow.ellipsis,
              ),
            ],
          ),
        ),
      ],
    );
  }

  Widget _buildInfoItem(IconData icon, String label, String value) {
    return Row(
      children: [
        Icon(icon, size: 16, color: Colors.grey[600]),
        const SizedBox(width: 6),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                label,
                style: TextStyle(
                  fontSize: 11,
                  color: Colors.grey[600],
                ),
              ),
              Text(
                value,
                style: const TextStyle(
                  fontSize: 13,
                  fontWeight: FontWeight.w600,
                ),
                overflow: TextOverflow.ellipsis,
              ),
            ],
          ),
        ),
      ],
    );
  }

  Future<void> _fetchCodeForNewPhone(String phoneNumber) async {
    final l10n = AppLocalizations.of(context)!;
    
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
                  content: Text('${l10n.code}: ${codeEntry.code}'),
                  backgroundColor: Colors.green,
                  duration: const Duration(seconds: 5),
                  action: SnackBarAction(
                    label: l10n.copy,
                    textColor: Colors.white,
                    onPressed: () {
                      Clipboard.setData(ClipboardData(text: codeEntry.code));
                    },
                  ),
                ),
              );
              // 刷新历史记录
              _loadRecentAssignments();
            } else if (codeEntry.status == 'pending') {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text(l10n.waitingForCode),
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
              content: Text(response.message ?? l10n.getFailed),
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

  String _formatAmount(double value) => value.toStringAsFixed(4);

  String _formatDate(DateTime? date) {
    if (date == null) return '--';
    return DateFormat('yyyy-MM-dd HH:mm:ss').format(date.toLocal());
  }
}
