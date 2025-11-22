class Whitelist {
  final String id;
  final String ipAddress;
  final String? notes;
  final DateTime createdAt;

  Whitelist({
    required this.id,
    required this.ipAddress,
    this.notes,
    required this.createdAt,
  });

  factory Whitelist.fromJson(Map<String, dynamic> json) {
    return Whitelist(
      id: json['id'] as String,
      ipAddress: json['ip_address'] as String,
      notes: json['notes'] as String?,
      createdAt: DateTime.parse(json['created_at'] as String),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'ip_address': ipAddress,
      'notes': notes,
      'created_at': createdAt.toIso8601String(),
    };
  }
}
