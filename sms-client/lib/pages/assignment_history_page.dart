import 'package:flutter/material.dart';
import '../core/api_client.dart';
import '../models/assignment.dart';
import '../l10n/app_localizations.dart';
import '../widgets/assignment_card.dart';

class AssignmentHistoryPage extends StatefulWidget {
  const AssignmentHistoryPage({super.key});

  @override
  State<AssignmentHistoryPage> createState() => _AssignmentHistoryPageState();
}

class _AssignmentHistoryPageState extends State<AssignmentHistoryPage> {
  final ApiClient _apiClient = ApiClient();
  final ScrollController _scrollController = ScrollController();
  
  List<Assignment> _assignments = [];
  bool _isLoading = false;
  bool _hasMore = true;
  int _currentPage = 1;
  final int _limit = 20;

  @override
  void initState() {
    super.initState();
    _scrollController.addListener(_onScroll);
    _loadAssignments();
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
      _loadMoreAssignments();
    }
  }

  Future<void> _loadAssignments({bool loadMore = false}) async {
    if (_isLoading) return;

    setState(() {
      _isLoading = true;
    });

    try {
      final page = loadMore ? _currentPage + 1 : 1;
      final response = await _apiClient.getAssignments(page: page, limit: _limit);

      if (response.success && response.data != null) {
        setState(() {
          if (loadMore) {
            _assignments.addAll(response.data!);
            _currentPage = page;
          } else {
            _assignments = response.data!;
            _currentPage = 1;
          }
          _hasMore = response.data!.length >= _limit;
        });
      } else {
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
      debugPrint('Error loading assignments: $e');
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

  Future<void> _loadMoreAssignments() async {
    await _loadAssignments(loadMore: true);
  }

  Future<void> _refreshAssignments() async {
    setState(() {
      _currentPage = 1;
      _hasMore = true;
    });
    await _loadAssignments(loadMore: false);
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    if (_isLoading && _assignments.isEmpty) {
      return const Center(child: CircularProgressIndicator());
    }

    if (_assignments.isEmpty) {
      return RefreshIndicator(
        onRefresh: _refreshAssignments,
        child: ListView(
          physics: const AlwaysScrollableScrollPhysics(),
          children: [
            const SizedBox(height: 120),
            const Icon(Icons.history, size: 64, color: Colors.grey),
            const SizedBox(height: 16),
            Center(
              child: Text(
                l10n.noData,
                style: const TextStyle(color: Colors.grey, fontSize: 16),
              ),
            ),
            const SizedBox(height: 16),
            Center(
              child: ElevatedButton.icon(
                onPressed: _refreshAssignments,
                icon: const Icon(Icons.refresh),
                label: Text(l10n.refresh),
              ),
            ),
          ],
        ),
      );
    }

    return Column(
      children: [
        // 顶部工具栏：刷新按钮
        Container(
          padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
          decoration: BoxDecoration(
            color: Colors.grey[100],
            border: Border(bottom: BorderSide(color: Colors.grey[300]!)),
          ),
          child: Row(
            children: [
              Text(
                '共 ${_assignments.length} 条记录',
                style: TextStyle(color: Colors.grey[600], fontSize: 14),
              ),
              const Spacer(),
              IconButton(
                icon: const Icon(Icons.refresh),
                onPressed: _isLoading ? null : _refreshAssignments,
                tooltip: l10n.refresh,
              ),
            ],
          ),
        ),
        // 列表
        Expanded(
          child: RefreshIndicator(
            onRefresh: _refreshAssignments,
            child: ListView.builder(
              controller: _scrollController,
              physics: const AlwaysScrollableScrollPhysics(),
              padding: const EdgeInsets.all(16),
              itemCount: _assignments.length + (_hasMore ? 1 : 0),
              itemBuilder: (context, index) {
                if (index == _assignments.length) {
                  return const Center(
                    child: Padding(
                      padding: EdgeInsets.all(16),
                      child: CircularProgressIndicator(),
                    ),
                  );
                }

                final assignment = _assignments[index];
                return AssignmentCard(
                  assignment: assignment,
                  onRefresh: _refreshAssignments,
                );
              },
            ),
          ),
        ),
      ],
    );
  }
}

