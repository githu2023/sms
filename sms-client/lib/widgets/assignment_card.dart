import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:intl/intl.dart';
import '../core/api_client.dart';
import '../l10n/app_localizations.dart';
import '../models/assignment.dart';

/// 统一的手机号分配记录卡片组件
/// 包含：手机号、业务类型、验证码、费用、状态等信息
/// 以及"获取验证码"按钮（根据状态显示）
class AssignmentCard extends StatelessWidget {
  final Assignment assignment;
  final VoidCallback? onRefresh; // 获取验证码后的回调

  const AssignmentCard({
    super.key,
    required this.assignment,
    this.onRefresh,
  });

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    // 只有 pending 状态且没有验证码时才显示"获取验证码"按钮
    final canFetchCode = assignment.code == null && assignment.status == 'pending';

    return Card(
      margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
      child: Padding(
        padding: const EdgeInsets.all(12),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Icon(
                  Icons.phone_android,
                  color: Theme.of(context).primaryColor,
                ),
                const SizedBox(width: 12),
                Expanded(
                  child: Text(
                    assignment.phone,
                    style: const TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                    ),
                  ),
                ),
                _buildStatusChip(assignment.status, l10n, context),
              ],
            ),
            const SizedBox(height: 12),
            Text('${l10n.businessType}: ${assignment.businessType}'),
            const SizedBox(height: 4),
            Row(
              children: [
                Expanded(
                  child: Text(
                    '${l10n.code}: ${_getCodeDisplay(assignment, l10n)}',
                    style: TextStyle(
                      color: _getCodeColor(assignment),
                      fontWeight:
                          assignment.code != null ? FontWeight.bold : FontWeight.normal,
                      fontSize: assignment.code != null ? 18 : 14,
                    ),
                  ),
                ),
                if (assignment.code != null && assignment.code!.isNotEmpty)
                  IconButton(
                    onPressed: () {
                      Clipboard.setData(ClipboardData(text: assignment.code!));
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          content: Text(l10n.copied),
                          duration: const Duration(seconds: 1),
                        ),
                      );
                    },
                    icon: const Icon(Icons.copy, size: 18),
                    tooltip: l10n.copy,
                  ),
              ],
            ),
            const SizedBox(height: 4),
            Text('${l10n.cost}: ¥${assignment.cost.toStringAsFixed(4)}'),
            const SizedBox(height: 4),
            Text(
              '${l10n.time}: ${DateFormat('yyyy-MM-dd HH:mm').format(assignment.createdAt)}',
              style: const TextStyle(fontSize: 12, color: Colors.grey),
            ),
            if (canFetchCode) ...[
              const SizedBox(height: 12),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: () => _fetchCodeForPhone(context, assignment.phone),
                  icon: const Icon(Icons.sms, size: 18),
                  label: Text(l10n.getCode),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.orange,
                    foregroundColor: Colors.white,
                  ),
                ),
              ),
            ],
          ],
        ),
      ),
    );
  }

  String _getCodeDisplay(Assignment assignment, AppLocalizations l10n) {
    if (assignment.code != null && assignment.code!.isNotEmpty) {
      return assignment.code!;
    }

    switch (assignment.status) {
      case 'pending':
        return l10n.waitingForCode;
      case 'completed':
        return l10n.noCodeReceived;
      case 'expired':
        return l10n.expired;
      case 'failed':
        return l10n.statusFailed;
      default:
        return '--';
    }
  }

  Color _getCodeColor(Assignment assignment) {
    if (assignment.code != null && assignment.code!.isNotEmpty) {
      return Colors.green;
    }

    switch (assignment.status) {
      case 'pending':
        return Colors.orange;
      case 'completed':
        return Colors.grey;
      case 'expired':
        return Colors.red;
      case 'failed':
        return Colors.red;
      default:
        return Colors.grey;
    }
  }

  Widget _buildStatusChip(String status, AppLocalizations l10n, BuildContext context) {
    Color color;
    String label;
    switch (status) {
      case 'completed':
        color = Colors.green;
        label = l10n.completed;
        break;
      case 'expired':
        color = Colors.red;
        label = l10n.expired;
        break;
      case 'pending':
        color = Colors.orange;
        label = l10n.statusPending;
        break;
      case 'failed':
        color = Colors.red;
        label = l10n.statusFailed;
        break;
      default:
        color = Colors.grey;
        label = status;
    }

    return Chip(
      label: Text(label, style: const TextStyle(fontSize: 12)),
      backgroundColor: color.withOpacity(0.2),
      labelStyle: TextStyle(color: color),
    );
  }

  Future<void> _fetchCodeForPhone(BuildContext context, String phoneNumber) async {
    final apiClient = ApiClient();
    final l10n = AppLocalizations.of(context)!;

    try {
      final response = await apiClient.getVerificationCodes(
        phoneNumbers: [phoneNumber],
      );

      if (response.success && response.data != null) {
        final result = response.data!;

        if (result.codes.isNotEmpty) {
          final codeEntry = result.codes.first;

          if (context.mounted) {
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
              // 调用刷新回调
              onRefresh?.call();
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
        if (context.mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(response.message ?? '${l10n.getFailed}'),
              backgroundColor: Colors.red,
            ),
          );
        }
      }
    } catch (e) {
      if (context.mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('网络错误: $e'),
            backgroundColor: Colors.red,
          ),
        );
      }
    }
  }
}

