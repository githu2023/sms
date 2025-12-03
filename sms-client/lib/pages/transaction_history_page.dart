import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import '../core/api_client.dart';
import '../models/transaction.dart';
import '../models/api_response.dart';
import '../l10n/app_localizations.dart';

class TransactionHistoryPage extends StatefulWidget {
  const TransactionHistoryPage({super.key});

  @override
  State<TransactionHistoryPage> createState() => _TransactionHistoryPageState();
}

class _TransactionHistoryPageState extends State<TransactionHistoryPage> {
  final ApiClient _apiClient = ApiClient();
  final ScrollController _scrollController = ScrollController();
  
  List<Transaction> _transactions = [];
  bool _isLoading = false;
  bool _hasMore = true;
  int _offset = 0;
  final int _limit = 20;
  int? _filterType; // null=全部, 1=充值, 2=消费
  String? _startDate;
  String? _endDate;

  @override
  void initState() {
    super.initState();
    _scrollController.addListener(_onScroll);
    _loadTransactions();
  }

  @override
  void dispose() {
    _scrollController.dispose();
    super.dispose();
  }

  void _onScroll() {
    if (_scrollController.position.pixels >=
            _scrollController.position.maxScrollExtent * 0.8 &&
        !_isLoading &&
        _hasMore) {
      _loadMoreTransactions();
    }
  }

  Future<void> _loadTransactions({bool loadMore = false}) async {
    if (_isLoading) return;

    setState(() {
      _isLoading = true;
    });

    try {
      final offset = loadMore ? _offset : 0;
      ApiResponse<TransactionListResponse> response;

      if (_filterType != null) {
        // 按类型筛选
        response = await _apiClient.getTransactionsByType(
          type: _filterType!,
          limit: _limit,
          offset: offset,
        );
      } else if (_startDate != null && _endDate != null) {
        // 按日期范围筛选
        response = await _apiClient.getTransactionsByDate(
          startDate: _startDate!,
          endDate: _endDate!,
          limit: _limit,
          offset: offset,
        );
      } else {
        // 获取全部
        response = await _apiClient.getTransactions(
          limit: _limit,
          offset: offset,
        );
      }

      debugPrint('Transaction response: success=${response.success}, data=${response.data}');
      
      if (response.success && response.data != null) {
        debugPrint('Transactions count: ${response.data!.transactions.length}');
        setState(() {
          if (loadMore) {
            _transactions.addAll(response.data!.transactions);
            _offset = response.data!.offset + response.data!.transactions.length;
          } else {
            _transactions = response.data!.transactions;
            _offset = response.data!.offset + response.data!.transactions.length;
          }
          _hasMore = _transactions.length < response.data!.total;
        });
        debugPrint('Loaded ${_transactions.length} transactions, hasMore=$_hasMore');
      } else {
        debugPrint('Failed to load transactions: ${response.message}, code=${response.code}');
        if (mounted) {
          final l10n = AppLocalizations.of(context)!;
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(response.message ?? l10n.loadingFailed),
              backgroundColor: Colors.red,
            ),
          );
        }
      }
    } catch (e) {
      debugPrint('Error loading transactions: $e');
      if (mounted) {
        final l10n = AppLocalizations.of(context)!;
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('${l10n.loadingFailed}: $e'),
            backgroundColor: Colors.red,
          ),
        );
      }
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  Future<void> _loadMoreTransactions() async {
    await _loadTransactions(loadMore: true);
  }

  Future<void> _refreshTransactions() async {
    setState(() {
      _offset = 0;
      _hasMore = true;
    });
    await _loadTransactions(loadMore: false);
  }

  Future<void> _selectDateRange() async {
    final DateTimeRange? picked = await showDateRangePicker(
      context: context,
      firstDate: DateTime(2020),
      lastDate: DateTime.now(),
      initialDateRange: _startDate != null && _endDate != null
          ? DateTimeRange(
              start: DateTime.parse(_startDate!),
              end: DateTime.parse(_endDate!),
            )
          : null,
    );

    if (picked != null) {
      setState(() {
        _startDate = DateFormat('yyyy-MM-dd').format(picked.start);
        _endDate = DateFormat('yyyy-MM-dd').format(picked.end);
        _filterType = null; // 清除类型筛选
        _offset = 0;
        _hasMore = true;
      });
      await _loadTransactions(loadMore: false);
    }
  }

  void _showFilterDialog() {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('筛选'),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            ListTile(
              title: const Text('全部'),
              leading: Radio<int?>(
                value: null,
                groupValue: _filterType,
                onChanged: (value) {
                  setState(() {
                    _filterType = value;
                    _startDate = null;
                    _endDate = null;
                  });
                  Navigator.pop(context);
                  _refreshTransactions();
                },
              ),
            ),
            ListTile(
              title: const Text('充值/入账'),
              leading: Radio<int?>(
                value: 1,
                groupValue: _filterType,
                onChanged: (value) {
                  setState(() {
                    _filterType = value;
                    _startDate = null;
                    _endDate = null;
                  });
                  Navigator.pop(context);
                  _refreshTransactions();
                },
              ),
            ),
            ListTile(
              title: const Text('消费'),
              leading: Radio<int?>(
                value: 2,
                groupValue: _filterType,
                onChanged: (value) {
                  setState(() {
                    _filterType = value;
                    _startDate = null;
                    _endDate = null;
                  });
                  Navigator.pop(context);
                  _refreshTransactions();
                },
              ),
            ),
          ],
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('余额变动记录'),
        actions: [
          IconButton(
            icon: const Icon(Icons.filter_list),
            onPressed: _showFilterDialog,
            tooltip: '筛选',
          ),
          IconButton(
            icon: const Icon(Icons.date_range),
            onPressed: _selectDateRange,
            tooltip: '选择日期范围',
          ),
        ],
      ),
      body: RefreshIndicator(
        onRefresh: _refreshTransactions,
        child: _transactions.isEmpty && !_isLoading
            ? Center(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Icon(
                      Icons.receipt_long,
                      size: 64,
                      color: Colors.grey[400],
                    ),
                    const SizedBox(height: 16),
                    Text(
                      '暂无交易记录',
                      style: TextStyle(
                        fontSize: 16,
                        color: Colors.grey[600],
                      ),
                    ),
                  ],
                ),
              )
            : ListView.builder(
                controller: _scrollController,
                padding: const EdgeInsets.all(16),
                itemCount: _transactions.length + (_hasMore ? 1 : 0),
                itemBuilder: (context, index) {
                  if (index == _transactions.length) {
                    return const Center(
                      child: Padding(
                        padding: EdgeInsets.all(16.0),
                        child: CircularProgressIndicator(),
                      ),
                    );
                  }

                  final transaction = _transactions[index];
                  final isIncome = transaction.isIncome;
                  final dateFormat = DateFormat('yyyy-MM-dd HH:mm:ss');

                  return Card(
                    margin: const EdgeInsets.only(bottom: 12),
                    child: ListTile(
                      contentPadding: const EdgeInsets.all(16),
                      leading: CircleAvatar(
                        backgroundColor: isIncome
                            ? Colors.green.withOpacity(0.2)
                            : Colors.red.withOpacity(0.2),
                        child: Icon(
                          isIncome ? Icons.arrow_downward : Icons.arrow_upward,
                          color: isIncome ? Colors.green : Colors.red,
                        ),
                      ),
                      title: Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Expanded(
                            child: Text(
                              transaction.getTypeName(),
                              style: const TextStyle(
                                fontWeight: FontWeight.bold,
                                fontSize: 16,
                              ),
                            ),
                          ),
                          Text(
                            '${isIncome ? '+' : ''}${transaction.amount.toStringAsFixed(4)}',
                            style: TextStyle(
                              color: isIncome ? Colors.green : Colors.red,
                              fontWeight: FontWeight.bold,
                              fontSize: 16,
                            ),
                          ),
                        ],
                      ),
                      subtitle: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          const SizedBox(height: 8),
                          Text(
                            dateFormat.format(transaction.createdAt.toLocal()),
                            style: TextStyle(
                              color: Colors.grey[600],
                              fontSize: 12,
                            ),
                          ),
                          if (transaction.notes != null &&
                              transaction.notes!.isNotEmpty) ...[
                            const SizedBox(height: 4),
                            Text(
                              transaction.notes!,
                              style: TextStyle(
                                color: Colors.grey[600],
                                fontSize: 12,
                              ),
                            ),
                          ],
                          const SizedBox(height: 8),
                          Row(
                            children: [
                              // 类型 8（冻结转实扣）显示冻结金额变动，其他类型显示余额变动
                              Text(
                                transaction.type == '8'
                                    ? '冻结: ${transaction.frozenBefore.toStringAsFixed(4)} → ${transaction.frozenAfter.toStringAsFixed(4)}'
                                    : '余额: ${transaction.balanceBefore.toStringAsFixed(4)} → ${transaction.balanceAfter.toStringAsFixed(4)}',
                                style: TextStyle(
                                  color: Colors.grey[600],
                                  fontSize: 11,
                                ),
                              ),
                            ],
                          ),
                        ],
                      ),
                    ),
                  );
                },
              ),
      ),
    );
  }
}

