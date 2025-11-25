class AssignedPhone {
  final String phoneNumber;
  final String countryCode;
  final double cost;
  final DateTime? validUntil;
  final String providerId;

  const AssignedPhone({
    required this.phoneNumber,
    required this.countryCode,
    required this.cost,
    required this.providerId,
    this.validUntil,
  });

  factory AssignedPhone.fromJson(Map<String, dynamic> json) {
    return AssignedPhone(
      phoneNumber: json['phone_number'] as String,
      countryCode: json['country_code'] as String? ?? '',
      cost: (json['cost'] as num?)?.toDouble() ?? 0.0,
      providerId: json['provider_id'] as String? ?? '',
      validUntil:
          json['valid_until'] != null
              ? DateTime.tryParse(json['valid_until'] as String)
              : null,
    );
  }
}
