import 'verification_code_entry.dart';

class VerificationCodeResult {
  final List<VerificationCodeEntry> codes;
  final int successCount;
  final int pendingCount;
  final int failedCount;

  const VerificationCodeResult({
    required this.codes,
    required this.successCount,
    required this.pendingCount,
    required this.failedCount,
  });

  factory VerificationCodeResult.fromJson(Map<String, dynamic> json) {
    final codesJson = json['codes'] as List<dynamic>? ?? [];
    return VerificationCodeResult(
      codes:
          codesJson
              .map(
                (item) => VerificationCodeEntry.fromJson(
                  item as Map<String, dynamic>,
                ),
              )
              .toList(),
      successCount: json['success_count'] as int? ?? 0,
      pendingCount: json['pending_count'] as int? ?? 0,
      failedCount: json['failed_count'] as int? ?? 0,
    );
  }
}
