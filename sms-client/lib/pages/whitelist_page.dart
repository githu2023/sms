import 'package:flutter/material.dart';
import '../l10n/app_localizations.dart';
import '../models/whitelist.dart';

class WhitelistPage extends StatefulWidget {
  const WhitelistPage({super.key});

  @override
  State<WhitelistPage> createState() => _WhitelistPageState();
}

class _WhitelistPageState extends State<WhitelistPage> {
  final List<Whitelist> _whitelists = [];
  bool _isLoading = false;

  @override
  void initState() {
    super.initState();
    _loadWhitelists();
  }

  Future<void> _loadWhitelists() async {
    setState(() {
      _isLoading = true;
    });

    // TODO: Load from API
    await Future.delayed(const Duration(seconds: 1));

    setState(() {
      _isLoading = false;
    });
  }

  Future<void> _addWhitelist() async {
    final result = await showDialog<Map<String, String>>(
      context: context,
      builder: (context) => const _AddWhitelistDialog(),
    );

    if (result != null) {
      // TODO: Add via API
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text('Added successfully'),
          backgroundColor: Colors.green,
        ),
      );
      _loadWhitelists();
    }
  }

  Future<void> _deleteWhitelist(Whitelist whitelist) async {
    final confirmed = await showDialog<bool>(
      context: context,
      builder: (context) => AlertDialog(
        title: Text(AppLocalizations.of(context)!.delete),
        content: Text('Delete ${whitelist.ipAddress}?'),
        actions: [
          TextButton(
            onPressed: () => Navigator.of(context).pop(false),
            child: const Text('Cancel'),
          ),
          TextButton(
            onPressed: () => Navigator.of(context).pop(true),
            style: TextButton.styleFrom(foregroundColor: Colors.red),
            child: Text(AppLocalizations.of(context)!.delete),
          ),
        ],
      ),
    );

    if (confirmed == true) {
      // TODO: Delete via API
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text('Deleted successfully'),
          backgroundColor: Colors.green,
        ),
      );
      _loadWhitelists();
    }
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    return Scaffold(
      appBar: AppBar(
        title: Text(l10n.whitelist),
        actions: [
          IconButton(
            icon: const Icon(Icons.add),
            onPressed: _addWhitelist,
          ),
        ],
      ),
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : _whitelists.isEmpty
              ? Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Icon(Icons.shield, size: 64, color: Colors.grey),
                      const SizedBox(height: 16),
                      Text(
                        l10n.noData,
                        style:
                            const TextStyle(color: Colors.grey, fontSize: 16),
                      ),
                    ],
                  ),
                )
              : ListView.builder(
                  itemCount: _whitelists.length,
                  itemBuilder: (context, index) {
                    final whitelist = _whitelists[index];
                    return Card(
                      margin: const EdgeInsets.symmetric(
                        horizontal: 16,
                        vertical: 8,
                      ),
                      child: ListTile(
                        leading: const Icon(Icons.computer),
                        title: Text(whitelist.ipAddress),
                        subtitle: whitelist.notes != null
                            ? Text(whitelist.notes!)
                            : null,
                        trailing: IconButton(
                          icon: const Icon(Icons.delete, color: Colors.red),
                          onPressed: () => _deleteWhitelist(whitelist),
                        ),
                      ),
                    );
                  },
                ),
    );
  }
}

class _AddWhitelistDialog extends StatefulWidget {
  const _AddWhitelistDialog();

  @override
  State<_AddWhitelistDialog> createState() => _AddWhitelistDialogState();
}

class _AddWhitelistDialogState extends State<_AddWhitelistDialog> {
  final _formKey = GlobalKey<FormState>();
  final _ipController = TextEditingController();
  final _notesController = TextEditingController();

  @override
  void dispose() {
    _ipController.dispose();
    _notesController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    return AlertDialog(
      title: Text(l10n.add),
      content: Form(
        key: _formKey,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextFormField(
              controller: _ipController,
              decoration: InputDecoration(
                labelText: l10n.ipAddress,
                hintText: '192.168.1.1',
              ),
              validator: (value) {
                if (value == null || value.trim().isEmpty) {
                  return 'Please enter IP address';
                }
                // Simple IP validation
                final ipPattern = RegExp(
                    r'^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$');
                if (!ipPattern.hasMatch(value)) {
                  return 'Invalid IP address';
                }
                return null;
              },
            ),
            const SizedBox(height: 16),
            TextFormField(
              controller: _notesController,
              decoration: InputDecoration(
                labelText: l10n.notes,
              ),
              maxLines: 2,
            ),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.of(context).pop(),
          child: const Text('Cancel'),
        ),
        ElevatedButton(
          onPressed: () {
            if (_formKey.currentState!.validate()) {
              Navigator.of(context).pop({
                'ip': _ipController.text.trim(),
                'notes': _notesController.text.trim(),
              });
            }
          },
          child: Text(l10n.add),
        ),
      ],
    );
  }
}
