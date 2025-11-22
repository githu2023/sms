class Assignment {
  final String phone;
  final String businessType;
  final String cardType;
  final String? code;
  final double cost;
  final String status;
  final DateTime createdAt;

  Assignment({
    required this.phone,
    required this.businessType,
    required this.cardType,
    this.code,
    required this.cost,
    required this.status,
    required this.createdAt,
  });

  factory Assignment.fromJson(Map<String, dynamic> json) {
    return Assignment(
      phone: json['phone'] as String,
      businessType: json['business_type'] as String,
      cardType: json['card_type'] as String,
      code: json['code'] as String?,
      cost: (json['cost'] as num).toDouble(),
      status: json['status'] as String,
      createdAt: DateTime.parse(json['created_at'] as String),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'phone': phone,
      'business_type': businessType,
      'card_type': cardType,
      'code': code,
      'cost': cost,
      'status': status,
      'created_at': createdAt.toIso8601String(),
    };
  }

  @override
  String toString() {
    return 'Assignment(phone: $phone, businessType: $businessType, code: $code, status: $status)';
  }
}
