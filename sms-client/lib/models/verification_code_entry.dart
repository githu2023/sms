class VerificationCodeEntry {
  final String phoneNumber;
  final String code;
  final String message;
  final DateTime? receivedAt;
  final String providerId;
  final String status;

  const VerificationCodeEntry({
    required this.phoneNumber,
    required this.code,
    required this.message,
    required this.providerId,
    required this.status,
    this.receivedAt,
  });

  factory VerificationCodeEntry.fromJson(Map<String, dynamic> json) {
    return VerificationCodeEntry(
      phoneNumber: json['phone_number'] as String? ?? '',
      code: json['code'] as String? ?? '',
      message: json['message'] as String? ?? '',
      providerId: json['provider_id'] as String? ?? '',
      status: json['status'] as String? ?? '',
      receivedAt:
          json['received_at'] != null
              ? DateTime.tryParse(json['received_at'] as String)
              : null,
    );
  }
}
