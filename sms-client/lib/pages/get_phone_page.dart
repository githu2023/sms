import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import '../core/api_client.dart';
import '../models/business_type.dart';
import '../l10n/app_localizations.dart';

class GetPhonePage extends StatefulWidget {
  const GetPhonePage({super.key});

  @override
  State<GetPhonePage> createState() => _GetPhonePageState();
}

class _GetPhonePageState extends State<GetPhonePage> {
  final ApiClient _apiClient = ApiClient();

  List<BusinessType> _businessTypes = [];
  BusinessType? _selectedBusinessType;
  String _selectedCardType = 'physical';
  String? _assignedPhone;
  bool _isLoading = false;
  bool _isLoadingTypes = false;
  String? _error;

  @override
  void initState() {
    super.initState();
    _loadBusinessTypes();
  }

  Future<void> _loadBusinessTypes() async {
    setState(() {
      _isLoadingTypes = true;
      _error = null;
    });

    try {
      final response = await _apiClient.getBusinessTypes();
      if (response.success && response.data != null) {
        setState(() {
          _businessTypes = response.data!;
          if (_businessTypes.isNotEmpty) {
            _selectedBusinessType = _businessTypes.first;
          }
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
        _isLoadingTypes = false;
      });
    }
  }

  Future<void> _assignPhone() async {
    if (_selectedBusinessType == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text(AppLocalizations.of(context)!.businessTypeRequired)),
      );
      return;
    }

    setState(() {
      _isLoading = true;
      _error = null;
      _assignedPhone = null;
    });

    try {
      final response = await _apiClient.assignPhone(
        businessType: _selectedBusinessType!.code,
        cardType: _selectedCardType,
      );

      if (response.success && response.data != null) {
        setState(() {
          _assignedPhone = response.data!;
        });
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(AppLocalizations.of(context)!.assignSuccess),
              backgroundColor: Colors.green,
            ),
          );
        }
      } else {
        setState(() {
          _error = response.message;
        });
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(response.message ?? 'Error'),
              backgroundColor: Colors.red,
            ),
          );
        }
      }
    } catch (e) {
      setState(() {
        _error = 'Network error: $e';
      });
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('Network error: $e'),
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

  void _copyPhone() {
    if (_assignedPhone != null) {
      Clipboard.setData(ClipboardData(text: _assignedPhone!));
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text(AppLocalizations.of(context)!.copied),
          duration: const Duration(seconds: 1),
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    return Scaffold(
      appBar: AppBar(
        title: Text(l10n.getPhoneTitle),
      ),
      body: _isLoadingTypes
          ? Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const CircularProgressIndicator(),
                  const SizedBox(height: 16),
                  Text(l10n.loadingBusinessTypes),
                ],
              ),
            )
          : _error != null && _businessTypes.isEmpty
              ? Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Icon(Icons.error_outline, size: 64, color: Colors.grey),
                      const SizedBox(height: 16),
                      Text(_error!, style: TextStyle(color: Colors.grey)),
                      const SizedBox(height: 16),
                      ElevatedButton.icon(
                        onPressed: _loadBusinessTypes,
                        icon: const Icon(Icons.refresh),
                        label: Text(l10n.retryGetCode),
                      ),
                    ],
                  ),
                )
              : SingleChildScrollView(
                  padding: const EdgeInsets.all(16),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    children: [
                      // Business Type Selection
                      Card(
                        child: Padding(
                          padding: const EdgeInsets.all(16),
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Text(
                                l10n.selectBusinessType,
                                style: Theme.of(context).textTheme.titleMedium,
                              ),
                              const SizedBox(height: 12),
                              if (_businessTypes.isEmpty)
                                Text(
                                  l10n.noBusinessTypes,
                                  style: TextStyle(color: Colors.grey),
                                )
                              else
                                DropdownButtonFormField<BusinessType>(
                                  value: _selectedBusinessType,
                                  decoration: InputDecoration(
                                    border: OutlineInputBorder(),
                                    contentPadding: const EdgeInsets.symmetric(
                                      horizontal: 12,
                                      vertical: 8,
                                    ),
                                  ),
                                  items: _businessTypes.map((type) {
                                    return DropdownMenuItem(
                                      value: type,
                                      child: Text(type.name),
                                    );
                                  }).toList(),
                                  onChanged: (value) {
                                    setState(() {
                                      _selectedBusinessType = value;
                                    });
                                  },
                                ),
                            ],
                          ),
                        ),
                      ),
                      const SizedBox(height: 16),

                      // Card Type Selection
                      Card(
                        child: Padding(
                          padding: const EdgeInsets.all(16),
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Text(
                                l10n.selectCardType,
                                style: Theme.of(context).textTheme.titleMedium,
                              ),
                              const SizedBox(height: 12),
                              Row(
                                children: [
                                  Expanded(
                                    child: RadioListTile<String>(
                                      title: Text(l10n.physicalCard),
                                      value: 'physical',
                                      groupValue: _selectedCardType,
                                      onChanged: (value) {
                                        setState(() {
                                          _selectedCardType = value!;
                                        });
                                      },
                                      contentPadding: EdgeInsets.zero,
                                    ),
                                  ),
                                  Expanded(
                                    child: RadioListTile<String>(
                                      title: Text(l10n.virtualCard),
                                      value: 'virtual',
                                      groupValue: _selectedCardType,
                                      onChanged: (value) {
                                        setState(() {
                                          _selectedCardType = value!;
                                        });
                                      },
                                      contentPadding: EdgeInsets.zero,
                                    ),
                                  ),
                                ],
                              ),
                            ],
                          ),
                        ),
                      ),
                      const SizedBox(height: 24),

                      // Assign Button
                      FilledButton.icon(
                        onPressed: _isLoading || _businessTypes.isEmpty ? null : _assignPhone,
                        icon: _isLoading
                            ? const SizedBox(
                                width: 20,
                                height: 20,
                                child: CircularProgressIndicator(
                                  strokeWidth: 2,
                                  color: Colors.white,
                                ),
                              )
                            : const Icon(Icons.phone_android),
                        label: Text(
                          _isLoading ? l10n.loading : l10n.assignNow,
                          style: const TextStyle(fontSize: 16),
                        ),
                        style: FilledButton.styleFrom(
                          padding: const EdgeInsets.symmetric(vertical: 16),
                        ),
                      ),
                      const SizedBox(height: 24),

                      // Result Display
                      if (_assignedPhone != null)
                        Card(
                          color: Theme.of(context).colorScheme.primaryContainer,
                          child: Padding(
                            padding: const EdgeInsets.all(16),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Row(
                                  children: [
                                    Icon(
                                      Icons.check_circle,
                                      color: Theme.of(context).colorScheme.primary,
                                    ),
                                    const SizedBox(width: 8),
                                    Text(
                                      l10n.assignmentResult,
                                      style: Theme.of(context).textTheme.titleMedium?.copyWith(
                                            color: Theme.of(context).colorScheme.primary,
                                            fontWeight: FontWeight.bold,
                                          ),
                                    ),
                                  ],
                                ),
                                const Divider(height: 24),
                                Text(
                                  l10n.phoneNumber,
                                  style: Theme.of(context).textTheme.bodySmall,
                                ),
                                const SizedBox(height: 8),
                                Row(
                                  children: [
                                    Expanded(
                                      child: Text(
                                        _assignedPhone!,
                                        style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                                              fontWeight: FontWeight.bold,
                                              letterSpacing: 1.5,
                                            ),
                                      ),
                                    ),
                                    IconButton.filled(
                                      onPressed: _copyPhone,
                                      icon: const Icon(Icons.copy),
                                      tooltip: l10n.copyPhone,
                                    ),
                                  ],
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
}
