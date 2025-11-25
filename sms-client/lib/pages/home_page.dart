import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:provider/provider.dart';
import '../providers/auth_provider.dart';
import '../l10n/app_localizations.dart';
import '../widgets/language_picker.dart';
import '../widgets/theme_picker.dart';
import '../core/api_client.dart';
import '../models/assignment.dart';
import 'login_page.dart';
import 'whitelist_page.dart';
import 'change_password_page.dart';
import 'get_phone_page.dart';
import 'get_code_page.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final ApiClient _apiClient = ApiClient();
  int _selectedIndex = 0;
  List<Assignment> _assignments = [];
  bool _isLoadingAssignments = false;
  int _currentPage = 1;
  bool _hasMore = true;

  @override
  void initState() {
    super.initState();
    // Load user profile and assignments on page load
    WidgetsBinding.instance.addPostFrameCallback((_) {
      context.read<AuthProvider>().loadUserProfile();
      _loadAssignments(); // 立即加载历史记录
    });
  }

  Future<void> _loadAssignments({bool loadMore = false}) async {
    if (_isLoadingAssignments) return;

    setState(() {
      _isLoadingAssignments = true;
    });

    try {
      final page = loadMore ? _currentPage + 1 : 1;
      debugPrint('Loading assignments: page=$page, loadMore=$loadMore');
      final response = await _apiClient.getAssignments(page: page, limit: 20);

      debugPrint(
        'Assignments response: success=${response.success}, data=${response.data?.length}, message=${response.message}',
      );

      if (response.success && response.data != null) {
        debugPrint('Got ${response.data!.length} assignments');
        setState(() {
          if (loadMore) {
            _assignments.addAll(response.data!);
            _currentPage = page;
          } else {
            _assignments = response.data!;
            _currentPage = 1;
          }
          _hasMore = response.data!.length >= 20;
        });
      } else {
        debugPrint('Failed to load assignments: ${response.message}');
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(response.message ?? '加载失败'),
              backgroundColor: Colors.red,
            ),
          );
        }
      }
    } catch (e) {
      debugPrint('Exception loading assignments: $e');
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('加载失败: $e'), backgroundColor: Colors.red),
        );
      }
    } finally {
      setState(() {
        _isLoadingAssignments = false;
      });
    }
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  Future<void> _handleLogout() async {
    final confirmed = await showDialog<bool>(
      context: context,
      builder:
          (context) => AlertDialog(
            title: Text(AppLocalizations.of(context)!.logout),
            content: const Text('Are you sure you want to logout?'),
            actions: [
              TextButton(
                onPressed: () => Navigator.of(context).pop(false),
                child: const Text('Cancel'),
              ),
              TextButton(
                onPressed: () => Navigator.of(context).pop(true),
                child: const Text('Logout'),
              ),
            ],
          ),
    );

    if (confirmed == true && mounted) {
      await context.read<AuthProvider>().logout();
      if (mounted) {
        Navigator.of(
          context,
        ).pushReplacement(MaterialPageRoute(builder: (_) => const LoginPage()));
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    final authProvider = context.watch<AuthProvider>();
    final user = authProvider.user;

    return Scaffold(
      appBar: AppBar(
        title: Text(_getTitle(l10n)),
        actions: [
          IconButton(icon: const Icon(Icons.logout), onPressed: _handleLogout),
        ],
      ),
      body: _buildBody(l10n, user),
      bottomNavigationBar: BottomNavigationBar(
        items: [
          BottomNavigationBarItem(
            icon: const Icon(Icons.home),
            label: l10n.home,
          ),
          BottomNavigationBarItem(
            icon: const Icon(Icons.history),
            label: l10n.history,
          ),
          BottomNavigationBarItem(
            icon: const Icon(Icons.settings),
            label: l10n.settings,
          ),
        ],
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
      ),
    );
  }

  String _getTitle(AppLocalizations l10n) {
    switch (_selectedIndex) {
      case 0:
        return l10n.home;
      case 1:
        return l10n.history;
      case 2:
        return l10n.settings;
      default:
        return l10n.appTitle;
    }
  }

  Widget _buildBody(AppLocalizations l10n, user) {
    switch (_selectedIndex) {
      case 0:
        return _buildHomeTab(l10n, user);
      case 1:
        return _buildHistoryTab(l10n);
      case 2:
        return _buildSettingsTab(l10n);
      default:
        return const SizedBox();
    }
  }

  Widget _buildHomeTab(AppLocalizations l10n, user) {
    return RefreshIndicator(
      onRefresh: () async {
        await context.read<AuthProvider>().loadUserProfile();
      },
      child: SingleChildScrollView(
        physics: const AlwaysScrollableScrollPhysics(),
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            // Balance card
            Card(
              child: Padding(
                padding: const EdgeInsets.all(16),
                child: Column(
                  children: [
                    Text(
                      l10n.balance,
                      style: Theme.of(context).textTheme.titleMedium,
                    ),
                    const SizedBox(height: 8),
                    Text(
                      user != null ? user.balance.toStringAsFixed(4) : '0.0000',
                      style: Theme.of(
                        context,
                      ).textTheme.headlineLarge?.copyWith(
                        color: Theme.of(context).primaryColor,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    const SizedBox(height: 16),
                    ElevatedButton.icon(
                      onPressed: () {
                        // TODO: Navigate to recharge page
                      },
                      icon: const Icon(Icons.add),
                      label: Text(l10n.recharge),
                    ),
                  ],
                ),
              ),
            ),
            const SizedBox(height: 16),

            // Quick actions
            Row(
              children: [
                Expanded(
                  child: _buildQuickActionCard(
                    icon: Icons.phone_android,
                    label: l10n.getPhone,
                    onTap: () {
                      Navigator.push(
                        context,
                        MaterialPageRoute(builder: (_) => const GetPhonePage()),
                      );
                    },
                  ),
                ),
                const SizedBox(width: 16),
                Expanded(
                  child: _buildQuickActionCard(
                    icon: Icons.message,
                    label: l10n.getCode,
                    onTap: () {
                      Navigator.push(
                        context,
                        MaterialPageRoute(builder: (_) => const GetCodePage()),
                      );
                    },
                  ),
                ),
              ],
            ),
            const SizedBox(height: 16),

            // View history button
            ElevatedButton.icon(
              onPressed: () {
                setState(() {
                  _selectedIndex = 1; // 切换到历史记录标签页
                });
              },
              icon: const Icon(Icons.history),
              label: Text(l10n.history),
              style: ElevatedButton.styleFrom(
                padding: const EdgeInsets.all(16),
              ),
            ),
            const SizedBox(height: 24),

            // Recent assignments
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  l10n.recentAssignments,
                  style: Theme.of(context).textTheme.titleMedium?.copyWith(
                    fontWeight: FontWeight.bold,
                  ),
                ),
                if (_assignments.isNotEmpty)
                  TextButton(
                    onPressed: () {
                      setState(() {
                        _selectedIndex = 1;
                      });
                    },
                    child: const Text('查看全部'),
                  ),
              ],
            ),
            const SizedBox(height: 8),
            if (_assignments.isEmpty)
              Card(
                child: Padding(
                  padding: const EdgeInsets.all(16),
                  child: Column(
                    children: [
                      const Icon(Icons.inbox, size: 48, color: Colors.grey),
                      const SizedBox(height: 8),
                      Text(
                        l10n.noData,
                        style: const TextStyle(color: Colors.grey),
                      ),
                    ],
                  ),
                ),
              )
            else
              // 显示最近3条记录
              ...(_assignments
                  .take(3)
                  .map((assignment) => _buildAssignmentCard(assignment, l10n))),
          ],
        ),
      ),
    );
  }

  Widget _buildQuickActionCard({
    required IconData icon,
    required String label,
    required VoidCallback onTap,
  }) {
    return Card(
      child: InkWell(
        onTap: onTap,
        child: Padding(
          padding: const EdgeInsets.all(16),
          child: Column(
            children: [
              Icon(icon, size: 32, color: Theme.of(context).primaryColor),
              const SizedBox(height: 8),
              Text(label, textAlign: TextAlign.center),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildHistoryTab(AppLocalizations l10n) {
    if (_isLoadingAssignments && _assignments.isEmpty) {
      return const Center(child: CircularProgressIndicator());
    }

    if (_assignments.isEmpty) {
      return RefreshIndicator(
        onRefresh: () => _loadAssignments(),
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
                onPressed: _loadAssignments,
                icon: const Icon(Icons.refresh),
                label: const Text('刷新'),
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
                onPressed: _isLoadingAssignments ? null : _loadAssignments,
                tooltip: '刷新',
              ),
            ],
          ),
        ),
        // 列表
        Expanded(
          child: RefreshIndicator(
            onRefresh: () => _loadAssignments(),
            child: ListView.builder(
              physics: const AlwaysScrollableScrollPhysics(),
              itemCount: _assignments.length + (_hasMore ? 1 : 0),
              itemBuilder: (context, index) {
                if (index == _assignments.length) {
                  // 加载更多 - 只在滚动到这里时触发一次
                  if (!_isLoadingAssignments && _hasMore) {
                    // 使用 Future.microtask 避免在 build 期间调用 setState
                    Future.microtask(() => _loadAssignments(loadMore: true));
                  }
                  return const Center(
                    child: Padding(
                      padding: EdgeInsets.all(16),
                      child: CircularProgressIndicator(),
                    ),
                  );
                }

                final assignment = _assignments[index];
                return _buildAssignmentCard(assignment, l10n);
              },
            ),
          ),
        ),
      ],
    );
  }

  Widget _buildAssignmentCard(Assignment assignment, AppLocalizations l10n) {
    // 判断是否需要显示"获取验证码"按钮
    // 条件：没有验证码 且 状态是 pending 或 active（等待中或激活中）
    final bool canFetchCode =
        assignment.code == null &&
        (assignment.status == 'pending' || assignment.status == 'active');

    return Card(
      margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
      child: Padding(
        padding: const EdgeInsets.all(12),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Icon(
                  Icons.phone_android,
                  color: Theme.of(context).primaryColor,
                ),
                const SizedBox(width: 12),
                Expanded(
                  child: Text(
                    assignment.phone,
                    style: const TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                    ),
                  ),
                ),
                _buildStatusChip(assignment.status, l10n),
              ],
            ),
            const SizedBox(height: 12),
            Text('${l10n.businessType}: ${assignment.businessType}'),
            const SizedBox(height: 4),
            Text(
              '${l10n.code}: ${_getCodeDisplay(assignment)}',
              style: TextStyle(
                color: _getCodeColor(assignment),
                fontWeight:
                    assignment.code == null
                        ? FontWeight.bold
                        : FontWeight.normal,
                fontSize: assignment.code == null ? 16 : 14,
              ),
            ),
            if (canFetchCode) ...[
              const SizedBox(height: 8),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton.icon(
                  onPressed: () => _fetchCodeForPhone(assignment.phone),
                  icon: const Icon(Icons.sms, size: 18),
                  label: const Text('获取验证码'),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.orange,
                    foregroundColor: Colors.white,
                  ),
                ),
              ),
            ],
            const SizedBox(height: 4),
            Text('${l10n.cost}: ¥${assignment.cost.toStringAsFixed(4)}'),
            const SizedBox(height: 4),
            Text(
              '${l10n.time}: ${DateFormat('yyyy-MM-dd HH:mm').format(assignment.createdAt)}',
              style: const TextStyle(fontSize: 12, color: Colors.grey),
            ),
          ],
        ),
      ),
    );
  }

  Future<void> _fetchCodeForPhone(String phoneNumber) async {
    try {
      final response = await _apiClient.getVerificationCodes(
        phoneNumbers: [phoneNumber],
      );

      if (response.success && response.data != null) {
        final result = response.data!;

        if (result.codes.isNotEmpty) {
          final codeEntry = result.codes.first;

          if (codeEntry.status == 'success') {
            // 成功获取验证码
            if (mounted) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text('验证码: ${codeEntry.code}'),
                  backgroundColor: Colors.green,
                  duration: const Duration(seconds: 5),
                ),
              );
              // 刷新列表
              _loadAssignments();
            }
          } else if (codeEntry.status == 'pending') {
            // 还在等待中
            if (mounted) {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text('验证码还未收到，请稍后再试'),
                  backgroundColor: Colors.orange,
                ),
              );
            }
          } else {
            // 失败
            if (mounted) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text(codeEntry.message),
                  backgroundColor: Colors.red,
                ),
              );
            }
          }
        }
      } else {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(response.message ?? '请求失败'),
              backgroundColor: Colors.red,
            ),
          );
        }
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('网络错误: $e'), backgroundColor: Colors.red),
        );
      }
    }
  }

  String _getCodeDisplay(Assignment assignment) {
    if (assignment.code != null) {
      return assignment.code!;
    }

    // 根据状态返回不同的提示
    switch (assignment.status) {
      case 'pending':
      case 'active':
        return '等待验证码中...';
      case 'completed':
        return '未收到验证码';
      case 'expired':
        return '已过期';
      default:
        return '--';
    }
  }

  Color _getCodeColor(Assignment assignment) {
    if (assignment.code != null) {
      return Colors.black;
    }

    switch (assignment.status) {
      case 'pending':
      case 'active':
        return Colors.orange; // 等待中 - 橙色
      case 'completed':
        return Colors.grey; // 已完成但没验证码 - 灰色
      case 'expired':
        return Colors.red; // 已过期 - 红色
      default:
        return Colors.grey;
    }
  }

  Widget _buildStatusChip(String status, AppLocalizations l10n) {
    Color color;
    String label;
    switch (status) {
      case 'completed':
        color = Colors.green;
        label = l10n.completed;
        break;
      case 'expired':
        color = Colors.red;
        label = l10n.expired;
        break;
      default:
        color = Colors.orange;
        label = status;
    }

    return Chip(
      label: Text(label, style: const TextStyle(fontSize: 12)),
      backgroundColor: color.withValues(alpha: 0.2),
      labelStyle: TextStyle(color: color),
    );
  }

  Widget _buildSettingsTab(AppLocalizations l10n) {
    final authProvider = context.watch<AuthProvider>();
    final user = authProvider.user;

    return ListView(
      children: [
        if (user != null) ...[
          ListTile(
            leading: const Icon(Icons.person),
            title: Text(l10n.username),
            subtitle: Text(user.username),
          ),
          ListTile(
            leading: const Icon(Icons.email),
            title: Text(l10n.email),
            subtitle: Text(user.email),
          ),
          ListTile(
            leading: const Icon(Icons.key),
            title: Text(l10n.apiKey),
            subtitle: Text(user.apiKey ?? 'N/A'),
          ),
          ListTile(
            leading: const Icon(Icons.public),
            title: Text(l10n.regIp),
            subtitle: Text(user.registrationIp ?? '--'),
          ),
          ListTile(
            leading: const Icon(Icons.schedule),
            title: Text(l10n.lastLogin),
            subtitle: Text(_formatDateTime(user.lastLoginAt)),
          ),
          const Divider(),
        ],
        ListTile(
          leading: const Icon(Icons.language),
          title: Text(l10n.language),
          onTap: () {
            showDialog(
              context: context,
              builder: (context) => const LanguagePicker(),
            );
          },
        ),
        ListTile(
          leading: const Icon(Icons.palette),
          title: Text(l10n.theme),
          onTap: () {
            showDialog(
              context: context,
              builder: (context) => const ThemePicker(),
            );
          },
        ),
        ListTile(
          leading: const Icon(Icons.shield),
          title: Text(l10n.whitelist),
          onTap: () {
            Navigator.of(
              context,
            ).push(MaterialPageRoute(builder: (_) => const WhitelistPage()));
          },
        ),
        ListTile(
          leading: const Icon(Icons.lock),
          title: Text(l10n.changePassword),
          onTap: () {
            Navigator.of(context).push(
              MaterialPageRoute(builder: (_) => const ChangePasswordPage()),
            );
          },
        ),
        const Divider(),
        ListTile(
          leading: const Icon(Icons.logout, color: Colors.red),
          title: Text(l10n.logout, style: const TextStyle(color: Colors.red)),
          onTap: _handleLogout,
        ),
      ],
    );
  }

  String _formatDateTime(DateTime? dateTime) {
    if (dateTime == null) return '--';
    return DateFormat('yyyy-MM-dd HH:mm:ss').format(dateTime.toLocal());
  }
}
