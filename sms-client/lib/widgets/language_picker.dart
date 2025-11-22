import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../providers/locale_provider.dart';
import '../l10n/app_localizations.dart';

class LanguagePicker extends StatelessWidget {
  const LanguagePicker({super.key});

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    final localeProvider = context.watch<LocaleProvider>();

    return AlertDialog(
      title: Text(l10n.language),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          RadioListTile<Locale>(
            title: const Text('中文'),
            value: const Locale('zh'),
            groupValue: localeProvider.locale,
            onChanged: (locale) {
              if (locale != null) {
                localeProvider.setLocale(locale);
                Navigator.of(context).pop();
              }
            },
          ),
          RadioListTile<Locale>(
            title: const Text('English'),
            value: const Locale('en'),
            groupValue: localeProvider.locale,
            onChanged: (locale) {
              if (locale != null) {
                localeProvider.setLocale(locale);
                Navigator.of(context).pop();
              }
            },
          ),
        ],
      ),
    );
  }
}
