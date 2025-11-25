import 'package:flutter/material.dart';
import '../core/api_client.dart';
import '../l10n/app_localizations.dart';
import '../models/whitelist.dart';

class WhitelistPage extends StatefulWidget {
  const WhitelistPage({super.key});

  @override
  State<WhitelistPage> createState() => _WhitelistPageState();
}

class _WhitelistPageState extends State<WhitelistPage> {
  final ApiClient _apiClient = ApiClient();
  List<Whitelist> _whitelists = [];
  bool _isLoading = false;
  String? _error;

  @override
  void initState() {
    super.initState();
    _loadWhitelists();
  }

  Future<void> _loadWhitelists() async {
    setState(() {
      _isLoading = true;
      _error = null;
    });

    try {
      final response = await _apiClient.getWhitelists();
      if (response.success && response.data != null) {
        setState(() {
          _whitelists = response.data!;
        });
      } else {
        setState(() {
          _error = response.message;
        });
      }
    } catch (e) {
      setState(() {
        _error = 'Network error: $e';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  Future<void> _addWhitelist() async {
    final result = await showDialog<Map<String, String>>(
      context: context,
      builder: (context) => const _AddWhitelistDialog(),
    );

    if (result != null) {
      final ip = result['ip']!.trim();
      final notes = result['notes']?.trim();

      setState(() {
        _isLoading = true;
      });

      try {
        final response = await _apiClient.addWhitelist(
          ipAddress: ip,
          notes: notes?.isEmpty ?? true ? null : notes,
        );

        if (response.success && response.data != null) {
          setState(() {
            _whitelists = [..._whitelists, response.data!];
          });
          if (mounted) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(
                content: Text(
                  AppLocalizations.of(context)!.addWhitelistSuccess,
                ),
                backgroundColor: Colors.green,
              ),
            );
          }
        } else {
          if (mounted) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(
                content: Text(
                  response.message ??
                      AppLocalizations.of(context)!.addWhitelistFailed,
                ),
                backgroundColor: Colors.red,
              ),
            );
          }
        }
      } catch (e) {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(
                '${AppLocalizations.of(context)!.addWhitelistFailed}: $e',
              ),
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
  }

  Future<void> _deleteWhitelist(Whitelist whitelist) async {
    final l10n = AppLocalizations.of(context)!;
    final confirmed = await showDialog<bool>(
      context: context,
      builder:
          (context) => AlertDialog(
            title: Text(l10n.delete),
            content: Text(
              '${l10n.deleteWhitelistConfirm}\n${whitelist.ipAddress}',
            ),
            actions: [
              TextButton(
                onPressed: () => Navigator.of(context).pop(false),
                child: const Text('Cancel'),
              ),
              TextButton(
                onPressed: () => Navigator.of(context).pop(true),
                style: TextButton.styleFrom(foregroundColor: Colors.red),
                child: Text(l10n.delete),
              ),
            ],
          ),
    );

    if (confirmed == true) {
      setState(() {
        _isLoading = true;
      });

      try {
        final response = await _apiClient.deleteWhitelist(
          ipAddress: whitelist.ipAddress,
        );
        if (response.success) {
          setState(() {
            _whitelists =
                _whitelists.where((item) => item.id != whitelist.id).toList();
          });
          if (mounted) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(
                content: Text(l10n.deleteWhitelistSuccess),
                backgroundColor: Colors.green,
              ),
            );
          }
        } else {
          if (mounted) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(
                content: Text(response.message ?? l10n.deleteWhitelistFailed),
                backgroundColor: Colors.red,
              ),
            );
          }
        }
      } catch (e) {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text('${l10n.deleteWhitelistFailed}: $e'),
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
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    return Scaffold(
      appBar: AppBar(
        title: Text(l10n.whitelist),
        actions: [
          IconButton(icon: const Icon(Icons.add), onPressed: _addWhitelist),
        ],
      ),
      body:
          _isLoading
              ? const Center(child: CircularProgressIndicator())
              : _error != null
              ? _buildErrorState(l10n)
              : _whitelists.isEmpty
              ? _buildEmptyState(l10n)
              : RefreshIndicator(
                onRefresh: _loadWhitelists,
                child: ListView.builder(
                  physics: const AlwaysScrollableScrollPhysics(),
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
                        subtitle:
                            whitelist.notes?.isNotEmpty == true
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
              ),
    );
  }

  Widget _buildEmptyState(AppLocalizations l10n) {
    return RefreshIndicator(
      onRefresh: _loadWhitelists,
      child: ListView(
        physics: const AlwaysScrollableScrollPhysics(),
        children: [
          const SizedBox(height: 120),
          Icon(Icons.shield, size: 64, color: Colors.grey.shade400),
          const SizedBox(height: 16),
          Center(
            child: Text(
              l10n.noData,
              style: const TextStyle(color: Colors.grey, fontSize: 16),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildErrorState(AppLocalizations l10n) {
    return RefreshIndicator(
      onRefresh: _loadWhitelists,
      child: ListView(
        physics: const AlwaysScrollableScrollPhysics(),
        padding: const EdgeInsets.all(24),
        children: [
          const SizedBox(height: 80),
          Icon(Icons.error_outline, size: 64, color: Colors.red.shade300),
          const SizedBox(height: 16),
          Text(
            _error ?? l10n.error,
            textAlign: TextAlign.center,
            style: const TextStyle(color: Colors.red),
          ),
          const SizedBox(height: 16),
          FilledButton.icon(
            onPressed: _loadWhitelists,
            icon: const Icon(Icons.refresh),
            label: Text(l10n.retryGetCode),
          ),
        ],
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
                hintText: '1.2.3.4 / 1.2.3.0/24',
              ),
              validator: (value) {
                if (value == null || value.trim().isEmpty) {
                  return l10n.ipAddressRequired;
                }
                final pattern = RegExp(
                  r'^(\d{1,3}\.){3}\d{1,3}(\/([0-2]?[0-9]|3[0-2]))?$',
                );
                if (!pattern.hasMatch(value.trim())) {
                  return l10n.invalidIpFormat;
                }
                return null;
              },
            ),
            const SizedBox(height: 16),
            TextFormField(
              controller: _notesController,
              decoration: InputDecoration(labelText: l10n.notesOptional),
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
