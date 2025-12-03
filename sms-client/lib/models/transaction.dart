class Transaction {
  final int id;
  final double amount;
  final double balanceBefore;
  final double balanceAfter;
  final double frozenBefore;
  final double frozenAfter;
  final String type;
  final int? referenceId;
  final String? notes;
  final DateTime createdAt;

  Transaction({
    required this.id,
    required this.amount,
    required this.balanceBefore,
    required this.balanceAfter,
    required this.frozenBefore,
    required this.frozenAfter,
    required this.type,
    this.referenceId,
    this.notes,
    required this.createdAt,
  });

  factory Transaction.fromJson(Map<String, dynamic> json) {
    return Transaction(
      id: json['id'] as int,
      amount: (json['amount'] as num).toDouble(),
      balanceBefore: (json['balance_before'] as num).toDouble(),
      balanceAfter: (json['balance_after'] as num).toDouble(),
      frozenBefore: (json['frozen_before'] as num?)?.toDouble() ?? 0.0,
      frozenAfter: (json['frozen_after'] as num?)?.toDouble() ?? 0.0,
      type: json['type'].toString(),
      referenceId: json['reference_id'] as int?,
      notes: json['notes'] as String?,
      createdAt: DateTime.parse(json['created_at'] as String),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'amount': amount,
      'balance_before': balanceBefore,
      'balance_after': balanceAfter,
      'frozen_before': frozenBefore,
      'frozen_after': frozenAfter,
      'type': type,
      'reference_id': referenceId,
      'notes': notes,
      'created_at': createdAt.toIso8601String(),
    };
  }

  // 获取交易类型的中文名称
  String getTypeName() {
    switch (type) {
      case '1':
        return '充值/入账';
      case '2':
        return '拉号消费';
      case '3':
        return '拉号回退';
      case '4':
        return '上分';
      case '5':
        return '下分';
      case '6':
        return '预冻结';
      case '7':
        return '解冻';
      case '8':
        return '冻结转实扣';
      default:
        return '未知类型';
    }
  }

  // 判断是否为收入（正数）
  bool get isIncome => amount > 0;

  @override
  String toString() {
    return 'Transaction(id: $id, amount: $amount, type: $type, createdAt: $createdAt)';
  }
}

class TransactionListResponse {
  final int total;
  final int limit;
  final int offset;
  final List<Transaction> transactions;

  TransactionListResponse({
    required this.total,
    required this.limit,
    required this.offset,
    required this.transactions,
  });

  factory TransactionListResponse.fromJson(Map<String, dynamic> json) {
    final transactionsList = json['transactions'] as List<dynamic>;
    return TransactionListResponse(
      total: json['total'] as int,
      limit: json['limit'] as int,
      offset: json['offset'] as int,
      transactions: transactionsList
          .map((e) => Transaction.fromJson(e as Map<String, dynamic>))
          .toList(),
    );
  }
}

