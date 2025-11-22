class BusinessType {
  final String name;
  final String code;

  BusinessType({
    required this.name,
    required this.code,
  });

  factory BusinessType.fromJson(Map<String, dynamic> json) {
    return BusinessType(
      name: json['name'] as String,
      code: json['code'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'code': code,
    };
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
