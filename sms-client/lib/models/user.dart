class User {
  final String id;
  final String username;
  final String email;
  final double balance;
  final String? apiKey;
  final DateTime createdAt;
  final DateTime? updatedAt;

  User({
    required this.id,
    required this.username,
    required this.email,
    required this.balance,
    this.apiKey,
    required this.createdAt,
    this.updatedAt,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: (json['user_id'] ?? json['id']).toString(),
      username: json['username'] as String,
      email: json['email'] as String,
      balance: (json['balance'] as num).toDouble(),
      apiKey: json['api_secret_key'] as String?,
      createdAt: json['created_at'] != null
          ? DateTime.parse(json['created_at'] as String)
          : DateTime.now(),
      updatedAt: json['updated_at'] != null
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
    DateTime? createdAt,
    DateTime? updatedAt,
  }) {
    return User(
      id: id ?? this.id,
      username: username ?? this.username,
      email: email ?? this.email,
      balance: balance ?? this.balance,
      apiKey: apiKey ?? this.apiKey,
      createdAt: createdAt ?? this.createdAt,
      updatedAt: updatedAt ?? this.updatedAt,
    );
  }
}
