import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:intl/intl.dart';
import '../core/api_client.dart';
import '../l10n/app_localizations.dart';
import '../models/verification_code_entry.dart';
import '../models/verification_code_result.dart';

class GetCodePage extends StatefulWidget {
  const GetCodePage({super.key});

  @override
  State<GetCodePage> createState() => _GetCodePageState();
}

class _GetCodePageState extends State<GetCodePage> with RouteAware {
  final ApiClient _apiClient = ApiClient();
  final TextEditingController _phonesController = TextEditingController();
  VerificationCodeResult? _result;
  String? _error;
  bool _isLoading = false;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    // 检查验证码是否过期
    _checkExpired();
  }

  void _checkExpired() {
    if (_result != null && _result!.codes.isNotEmpty) {
      bool hasExpired = false;
      
      for (var entry in _result!.codes) {
        // 如果状态是 failed，说明已过期或失败，需要刷新状态
        if (entry.status == 'failed') {
          hasExpired = true;
          break;
        }
      }
      
      if (hasExpired) {
        // 清除过期结果，避免再次请求
        setState(() {
          _result = null;
        });
      }
    }
  }

  @override
  void dispose() {
    _phonesController.dispose();
    super.dispose();
  }

  Future<void> _getCodes() async {
    final phones = _parsePhoneNumbers(_phonesController.text);
    if (phones.isEmpty || phones.length > 10) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text(AppLocalizations.of(context)!.invalidPhoneInput),
        ),
      );
      return;
    }

    // 如果之前的结果中有过期状态，先清除
    if (_result != null) {
      bool hasExpired = _result!.codes.any((entry) => entry.status == 'failed');
      if (hasExpired) {
        setState(() {
          _result = null;
        });
      }
    }

    setState(() {
      _isLoading = true;
      _error = null;
    });

    try {
      final response = await _apiClient.getVerificationCodes(
        phoneNumbers: phones,
      );

      if (response.success && response.data != null) {
        setState(() {
          _result = response.data;
        });
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(AppLocalizations.of(context)!.codeReceived),
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

  void _copyCode(String code) {
    Clipboard.setData(ClipboardData(text: code));
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
      appBar: AppBar(title: Text(l10n.getCodeTitle)),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            // Info Card
            Card(
              color: Theme.of(context).colorScheme.surfaceContainerHighest,
              child: Padding(
                padding: const EdgeInsets.all(16),
                child: Row(
                  children: [
                    Icon(
                      Icons.info_outline,
                      color: Theme.of(context).colorScheme.primary,
                    ),
                    const SizedBox(width: 12),
                    Expanded(
                      child: Text(
                        l10n.getCodeInstructions,
                        style: Theme.of(context).textTheme.bodyMedium,
                      ),
                    ),
                  ],
                ),
              ),
            ),
            const SizedBox(height: 24),

            // Phone Input
            Card(
              child: Padding(
                padding: const EdgeInsets.all(16),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      l10n.phoneNumber,
                      style: Theme.of(context).textTheme.titleMedium,
                    ),
                    const SizedBox(height: 12),
                    TextField(
                      controller: _phonesController,
                      keyboardType: TextInputType.text,
                      minLines: 3,
                      maxLines: 5,
                      decoration: InputDecoration(
                        hintText: l10n.enterPhonesHint,
                        helperText: l10n.enterPhonesHelper,
                        alignLabelWithHint: true,
                        border: const OutlineInputBorder(),
                        enabled: !_isLoading,
                      ),
                      onSubmitted: (_) => _getCodes(),
                    ),
                  ],
                ),
              ),
            ),
            const SizedBox(height: 24),

            // Get Code Button
            FilledButton.icon(
              onPressed: _isLoading ? null : _getCodes,
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
                      : const Icon(Icons.sms),
              label: Text(
                _isLoading ? l10n.waitingForCode : l10n.getCodeNow,
                style: const TextStyle(fontSize: 16),
              ),
              style: FilledButton.styleFrom(
                padding: const EdgeInsets.symmetric(vertical: 16),
              ),
            ),

            // Loading Message
            if (_isLoading)
              Padding(
                padding: const EdgeInsets.only(top: 16),
                child: Card(
                  color: Theme.of(context).colorScheme.secondaryContainer,
                  child: Padding(
                    padding: const EdgeInsets.all(16),
                    child: Row(
                      children: [
                        const SizedBox(
                          width: 20,
                          height: 20,
                          child: CircularProgressIndicator(strokeWidth: 2),
                        ),
                        const SizedBox(width: 12),
                        Expanded(
                          child: Text(
                            l10n.waitingForCode,
                            style: Theme.of(context).textTheme.bodyMedium,
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ),

            if (_error != null)
              Padding(
                padding: const EdgeInsets.only(top: 16),
                child: Text(_error!, style: const TextStyle(color: Colors.red)),
              ),

            const SizedBox(height: 24),

            // Result Display
            if (_result != null) ...[
              _buildCodesSummaryCard(context, l10n),
              const SizedBox(height: 16),
              if (_result!.codes.isEmpty)
                Card(
                  child: Padding(
                    padding: const EdgeInsets.all(16),
                    child: Text(
                      l10n.noCodesYet,
                      style: const TextStyle(color: Colors.grey),
                    ),
                  ),
                )
              else
                Column(
                  children:
                      _result!.codes
                          .map((entry) => _buildCodeCard(context, entry, l10n))
                          .toList(),
                ),
            ],
          ],
        ),
      ),
    );
  }

  List<String> _parsePhoneNumbers(String raw) {
    return raw
        .split(RegExp(r'[,\n\r\s]+'))
        .map((e) => e.trim())
        .where((element) => element.isNotEmpty)
        .toSet()
        .toList();
  }

  Widget _buildCodesSummaryCard(BuildContext context, AppLocalizations l10n) {
    final data = _result!;
    return Card(
      color: Theme.of(context).colorScheme.primaryContainer,
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Icon(Icons.sms, color: Theme.of(context).colorScheme.primary),
                const SizedBox(width: 8),
                Text(
                  l10n.codesSummary,
                  style: Theme.of(context).textTheme.titleMedium?.copyWith(
                    color: Theme.of(context).colorScheme.primary,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
            const SizedBox(height: 12),
            _buildSummaryRow(
              l10n.successCountLabel,
              data.successCount.toString(),
            ),
            _buildSummaryRow(
              l10n.pendingCountLabel,
              data.pendingCount.toString(),
            ),
            _buildSummaryRow(
              l10n.failedCountLabel,
              data.failedCount.toString(),
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

  Widget _buildCodeCard(
    BuildContext context,
    VerificationCodeEntry entry,
    AppLocalizations l10n,
  ) {
    final codeAvailable = entry.code.isNotEmpty;
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
                    entry.phoneNumber,
                    style: Theme.of(context).textTheme.titleMedium?.copyWith(
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
                _buildStatusChip(entry.status, l10n),
              ],
            ),
            if (codeAvailable) ...[
              const SizedBox(height: 12),
              Row(
                children: [
                  Expanded(
                    child: Text(
                      entry.code,
                      style: Theme.of(
                        context,
                      ).textTheme.headlineSmall?.copyWith(
                        fontWeight: FontWeight.bold,
                        letterSpacing: 4,
                      ),
                    ),
                  ),
                  IconButton.filled(
                    onPressed: () => _copyCode(entry.code),
                    icon: const Icon(Icons.copy),
                    tooltip: l10n.copyCode,
                  ),
                ],
              ),
            ],
            const SizedBox(height: 12),
            Text(
              '${l10n.messageLabel}: ${entry.message.isEmpty ? '--' : entry.message}',
            ),
            const SizedBox(height: 4),
            Text('${l10n.receivedAtLabel}: ${_formatDate(entry.receivedAt)}'),
          ],
        ),
      ),
    );
  }

  Widget _buildStatusChip(String status, AppLocalizations l10n) {
    String label;
    switch (status) {
      case 'success':
        label = l10n.statusSuccess;
        break;
      case 'failed':
        label = l10n.statusFailed;
        break;
      default:
        label = l10n.statusPending;
        break;
    }

    final statusColor = _statusColor(status);
    return Chip(
      label: Text(label),
      backgroundColor: statusColor.withOpacity(0.15),
      labelStyle: TextStyle(
        color: statusColor,
        fontWeight: FontWeight.bold,
      ),
    );
  }

  Color _statusColor(String status) {
    switch (status) {
      case 'success':
        return Colors.green;
      case 'failed':
        return Colors.red;
      default:
        return Colors.orange;
    }
  }

  String _formatDate(DateTime? date) {
    if (date == null) return '--';
    return DateFormat('yyyy-MM-dd HH:mm:ss').format(date.toLocal());
  }
}
