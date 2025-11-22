import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../providers/theme_provider.dart';
import '../l10n/app_localizations.dart';

class ThemePicker extends StatelessWidget {
  const ThemePicker({super.key});

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    final themeProvider = context.watch<ThemeProvider>();

    return AlertDialog(
      title: Text(l10n.theme),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          RadioListTile<ThemeMode>(
            title: Text(l10n.light),
            value: ThemeMode.light,
            groupValue: themeProvider.themeMode,
            onChanged: (mode) {
              if (mode != null) {
                themeProvider.setThemeMode(mode);
                Navigator.of(context).pop();
              }
            },
          ),
          RadioListTile<ThemeMode>(
            title: Text(l10n.dark),
            value: ThemeMode.dark,
            groupValue: themeProvider.themeMode,
            onChanged: (mode) {
              if (mode != null) {
                themeProvider.setThemeMode(mode);
                Navigator.of(context).pop();
              }
            },
          ),
        ],
      ),
    );
  }
}
