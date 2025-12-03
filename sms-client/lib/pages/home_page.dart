import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:provider/provider.dart';
import '../providers/auth_provider.dart';
import '../l10n/app_localizations.dart';
import '../widgets/language_picker.dart';
import '../widgets/theme_picker.dart';
import 'login_page.dart';
import 'change_password_page.dart';
import 'get_phone_page.dart';
import 'get_code_page.dart';
import 'transaction_history_page.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _selectedIndex = 0;

  @override
  void initState() {
    super.initState();
    // Load user profile on page load
    WidgetsBinding.instance.addPostFrameCallback((_) {
      context.read<AuthProvider>().loadUserProfile();
    });
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
    // 切换到主页时刷新余额
    if (index == 0) {
      context.read<AuthProvider>().loadUserProfile();
    }
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
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: '主页',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.receipt_long),
            label: '余额记录',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings),
            label: '设置',
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
        return '主页';
      case 1:
        return '余额记录';
      case 2:
        return '设置';
      default:
        return l10n.appTitle;
    }
  }

  Widget _buildBody(AppLocalizations l10n, user) {
    switch (_selectedIndex) {
      case 0:
        return _buildHomeTab(l10n, user);
      case 1:
        return const TransactionHistoryPage();
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
            // Balance card - 只显示当前余额
            Card(
              child: Padding(
                padding: const EdgeInsets.all(24),
                child: Column(
                  children: [
                    Text(
                      '当前余额',
                      style: Theme.of(context).textTheme.titleLarge,
                    ),
                    const SizedBox(height: 16),
                    Text(
                      user != null ? user.balance.toStringAsFixed(4) : '0.0000',
                      style: Theme.of(
                        context,
                      ).textTheme.headlineLarge?.copyWith(
                        color: Theme.of(context).primaryColor,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ],
                ),
              ),
            ),
            const SizedBox(height: 16),

            // Quick actions - 快捷操作
            Row(
              children: [
                Expanded(
                  child: _buildQuickActionCard(
                    icon: Icons.phone_android,
                    label: l10n.getPhone,
                    onTap: () async {
                      await Navigator.push(
                        context,
                        MaterialPageRoute(builder: (_) => const GetPhonePage()),
                      );
                      // 返回时刷新余额
                      if (mounted && _selectedIndex == 0) {
                        context.read<AuthProvider>().loadUserProfile();
                      }
                    },
                  ),
                ),
                const SizedBox(width: 16),
                Expanded(
                  child: _buildQuickActionCard(
                    icon: Icons.message,
                    label: l10n.getCode,
                    onTap: () async {
                      await Navigator.push(
                        context,
                        MaterialPageRoute(builder: (_) => const GetCodePage()),
                      );
                      // 返回时刷新余额
                      if (mounted && _selectedIndex == 0) {
                        context.read<AuthProvider>().loadUserProfile();
                      }
                    },
                  ),
                ),
              ],
            ),
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
          leading: const Icon(Icons.lock),
          title: Text(l10n.changePassword),
          onTap: () {
            Navigator.of(context).push(
              MaterialPageRoute(builder: (_) => const ChangePasswordPage()),
            );
          },
        ),
        ListTile(
          leading: const Icon(Icons.receipt_long),
          title: const Text('余额变动记录'),
          onTap: () {
            Navigator.of(context).push(
              MaterialPageRoute(
                builder: (_) => const TransactionHistoryPage(),
              ),
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
