class User {
  final String id;
  final String username;
  final String email;
  final double balance;
  final String? apiKey;
  final String? registrationIp;
  final DateTime? lastLoginAt;
  final DateTime createdAt;
  final DateTime? updatedAt;

  User({
    required this.id,
    required this.username,
    required this.email,
    required this.balance,
    this.apiKey,
    this.registrationIp,
    this.lastLoginAt,
    required this.createdAt,
    this.updatedAt,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: (json['user_id'] ?? json['id']).toString(),
      username: json['username'] as String? ?? '',
      email: json['email'] as String? ?? '',
      balance: (json['balance'] as num?)?.toDouble() ?? 0.0,
      apiKey: json['api_secret_key'] as String?,
      registrationIp: json['registration_ip'] as String?,
      lastLoginAt:
          json['last_login_at'] != null
              ? DateTime.tryParse(json['last_login_at'] as String)
              : null,
      createdAt:
          json['created_at'] != null
              ? DateTime.parse(json['created_at'] as String)
              : DateTime.now(),
      updatedAt:
          json['updated_at'] != null
              ? DateTime.parse(json['updated_at'] as String)
              : null,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'user_id': id,
      'username': username,
      'email': email,
      'balance': balance,
      'api_secret_key': apiKey,
      'registration_ip': registrationIp,
      'last_login_at': lastLoginAt?.toIso8601String(),
      'created_at': createdAt.toIso8601String(),
      'updated_at': updatedAt?.toIso8601String(),
    };
  }

  User copyWith({
    String? id,
    String? username,
    String? email,
    double? balance,
    String? apiKey,
    String? registrationIp,
    DateTime? lastLoginAt,
    DateTime? createdAt,
    DateTime? updatedAt,
  }) {
    return User(
      id: id ?? this.id,
      username: username ?? this.username,
      email: email ?? this.email,
      balance: balance ?? this.balance,
      apiKey: apiKey ?? this.apiKey,
      registrationIp: registrationIp ?? this.registrationIp,
      lastLoginAt: lastLoginAt ?? this.lastLoginAt,
      createdAt: createdAt ?? this.createdAt,
      updatedAt: updatedAt ?? this.updatedAt,
    );
  }
}
