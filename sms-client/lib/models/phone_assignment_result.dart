import 'assigned_phone.dart';

class PhoneAssignmentResult {
  final List<AssignedPhone> phones;
  final double totalCost;
  final double remainingBalance;
  final int successCount;
  final int failedCount;

  const PhoneAssignmentResult({
    required this.phones,
    required this.totalCost,
    required this.remainingBalance,
    required this.successCount,
    required this.failedCount,
  });

  factory PhoneAssignmentResult.fromJson(Map<String, dynamic> json) {
    final phonesJson = json['phones'] as List<dynamic>? ?? [];
    return PhoneAssignmentResult(
      phones:
          phonesJson
              .map(
                (item) => AssignedPhone.fromJson(item as Map<String, dynamic>),
              )
              .toList(),
      totalCost: (json['total_cost'] as num?)?.toDouble() ?? 0.0,
      remainingBalance: (json['remaining_balance'] as num?)?.toDouble() ?? 0.0,
      successCount: json['success_count'] as int? ?? 0,
      failedCount: json['failed_count'] as int? ?? 0,
    );
  }
}
