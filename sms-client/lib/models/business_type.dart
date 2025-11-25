class BusinessType {
  final String name;
  final String code;

  BusinessType({required this.name, required this.code});

  factory BusinessType.fromJson(Map<String, dynamic> json) {
    return BusinessType(
      // 兼容两种字段名：business_name/name, business_code/code
      name: (json['business_name'] ?? json['name']) as String,
      code: (json['business_code'] ?? json['code']) as String,
    );
  }

  Map<String, dynamic> toJson() {
    return {'name': name, 'code': code};
  }

  @override
  String toString() => 'BusinessType(name: $name, code: $code)';

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;
    return other is BusinessType && other.code == code;
  }

  @override
  int get hashCode => code.hashCode;
}
