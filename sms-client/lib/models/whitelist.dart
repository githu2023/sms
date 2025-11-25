class Whitelist {
  final String id;
  final String ipAddress;
  final String? notes;
  final DateTime? createdAt;

  const Whitelist({
    required this.id,
    required this.ipAddress,
    this.notes,
    this.createdAt,
  });

  factory Whitelist.fromJson(Map<String, dynamic> json) {
    return Whitelist(
      id: (json['id'] ?? '').toString(),
      ipAddress: json['ip_address'] as String? ?? '',
      notes: json['notes'] as String?,
      createdAt:
          json['created_at'] != null
              ? DateTime.tryParse(json['created_at'] as String)
              : null,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'ip_address': ipAddress,
      'notes': notes,
      'created_at': createdAt?.toIso8601String(),
    };
  }
}
