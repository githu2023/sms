import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../providers/auth_provider.dart';
import '../l10n/app_localizations.dart';
import '../widgets/language_picker.dart';
import '../widgets/theme_picker.dart';
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
  }

  Future<void> _handleLogout() async {
    final confirmed = await showDialog<bool>(
      context: context,
      builder: (context) => AlertDialog(
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
        Navigator.of(context).pushReplacement(
          MaterialPageRoute(builder: (_) => const LoginPage()),
        );
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
          IconButton(
            icon: const Icon(Icons.logout),
            onPressed: _handleLogout,
          ),
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
                      '¥${user?.balance.toStringAsFixed(2) ?? '0.00'}',
                      style: Theme.of(context).textTheme.headlineLarge?.copyWith(
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
            
            // Recent assignments
            Text(
              l10n.recentAssignments,
              style: Theme.of(context).textTheme.titleMedium,
            ),
            const SizedBox(height: 8),
            Card(
              child: Padding(
                padding: const EdgeInsets.all(16),
                child: Column(
                  children: [
                    const Icon(
                      Icons.inbox,
                      size: 48,
                      color: Colors.grey,
                    ),
                    const SizedBox(height: 8),
                    Text(
                      l10n.noData,
                      style: const TextStyle(color: Colors.grey),
                    ),
                  ],
                ),
              ),
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

  Widget _buildHistoryTab(AppLocalizations l10n) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const Icon(Icons.history, size: 64, color: Colors.grey),
          const SizedBox(height: 16),
          Text(
            l10n.noData,
            style: const TextStyle(color: Colors.grey, fontSize: 16),
          ),
        ],
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
            Navigator.of(context).push(
              MaterialPageRoute(builder: (_) => const WhitelistPage()),
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
        const Divider(),
        ListTile(
          leading: const Icon(Icons.logout, color: Colors.red),
          title: Text(
            l10n.logout,
            style: const TextStyle(color: Colors.red),
          ),
          onTap: _handleLogout,
        ),
      ],
    );
  }
}
