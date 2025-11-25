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
    // 状态映射：数字状态转为字符串
    // 根据 API 文档: 1=pending、2=completed、3=expired、4=failed
    String statusFromInt(int statusCode) {
      switch (statusCode) {
        case 1:
          return 'pending'; // 等待中
        case 2:
          return 'completed'; // 已完成
        case 3:
          return 'expired'; // 已过期
        case 4:
          return 'failed'; // 失败
        default:
          return 'unknown';
      }
    }

    // 处理 code 字段：支持 verification_code 和 code 两种字段名，空字符串转为 null
    String? codeValue = (json['verification_code'] ?? json['code']) as String?;
    if (codeValue != null && codeValue.isEmpty) {
      codeValue = null;
    }

    return Assignment(
      phone: (json['phone_number'] ?? json['phone']) as String,
      businessType: json['business_type'] as String,
      cardType: json['card_type'] as String,
      code: codeValue,
      cost: (json['cost'] as num).toDouble(),
      status: json['status'] is int 
          ? statusFromInt(json['status'] as int)
          : json['status'] as String,
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
