import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:provider/provider.dart';
import 'l10n/app_localizations.dart';

import 'core/theme.dart';
import 'providers/locale_provider.dart';
import 'providers/theme_provider.dart';
import 'providers/auth_provider.dart';
import 'pages/splash_page.dart';

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => LocaleProvider()),
        ChangeNotifierProvider(create: (_) => ThemeProvider()),
        ChangeNotifierProvider(create: (_) => AuthProvider()),
      ],
      child: Consumer2<LocaleProvider, ThemeProvider>(
        builder: (context, localeProvider, themeProvider, child) {
          return MaterialApp(
            title: 'SMS Platform',
            debugShowCheckedModeBanner: false,

            // 主题配置
            theme: AppTheme.lightTheme,
            darkTheme: AppTheme.darkTheme,
            themeMode: themeProvider.themeMode,

            // 国际化配置
            locale: localeProvider.locale,
            localizationsDelegates: const [
              AppLocalizations.delegate,
              GlobalMaterialLocalizations.delegate,
              GlobalWidgetsLocalizations.delegate,
              GlobalCupertinoLocalizations.delegate,
            ],
            supportedLocales: const [Locale('zh'), Locale('en')],

            // 限制最大宽度，保持移动端样式（H5样式）
            builder: (context, child) {
              return LayoutBuilder(
                builder: (context, constraints) {
                  // 移动端最大宽度（类似手机屏幕宽度）
                  const double maxMobileWidth = 600.0;
                  
                  // 如果屏幕宽度超过移动端宽度，居中显示并限制宽度
                  if (constraints.maxWidth > maxMobileWidth) {
                    return Container(
                      color: Theme.of(context).scaffoldBackgroundColor,
                      child: Center(
                        child: Container(
                          width: maxMobileWidth,
                          constraints: BoxConstraints(
                            maxWidth: maxMobileWidth,
                            minHeight: constraints.maxHeight,
                          ),
                          decoration: BoxDecoration(
                            color: Theme.of(context).scaffoldBackgroundColor,
                            boxShadow: [
                              BoxShadow(
                                color: Colors.black.withOpacity(0.1),
                                blurRadius: 10,
                                spreadRadius: 2,
                              ),
                            ],
                          ),
                          child: child,
                        ),
                      ),
                    );
                  }
                  
                  // 移动端直接显示
                  return child!;
                },
              );
            },

            // 路由配置 - 启动时检查登录状态
            home: const SplashPage(),
          );
        },
      ),
    );
  }
}
