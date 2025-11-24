import 'dart:async';

import 'package:flutter/foundation.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:intl/intl.dart' as intl;

import 'app_localizations_en.dart';
import 'app_localizations_zh.dart';

// ignore_for_file: type=lint

/// Callers can lookup localized strings with an instance of AppLocalizations
/// returned by `AppLocalizations.of(context)`.
///
/// Applications need to include `AppLocalizations.delegate()` in their app's
/// `localizationDelegates` list, and the locales they support in the app's
/// `supportedLocales` list. For example:
///
/// ```dart
/// import 'l10n/app_localizations.dart';
///
/// return MaterialApp(
///   localizationsDelegates: AppLocalizations.localizationsDelegates,
///   supportedLocales: AppLocalizations.supportedLocales,
///   home: MyApplicationHome(),
/// );
/// ```
///
/// ## Update pubspec.yaml
///
/// Please make sure to update your pubspec.yaml to include the following
/// packages:
///
/// ```yaml
/// dependencies:
///   # Internationalization support.
///   flutter_localizations:
///     sdk: flutter
///   intl: any # Use the pinned version from flutter_localizations
///
///   # Rest of dependencies
/// ```
///
/// ## iOS Applications
///
/// iOS applications define key application metadata, including supported
/// locales, in an Info.plist file that is built into the application bundle.
/// To configure the locales supported by your app, you’ll need to edit this
/// file.
///
/// First, open your project’s ios/Runner.xcworkspace Xcode workspace file.
/// Then, in the Project Navigator, open the Info.plist file under the Runner
/// project’s Runner folder.
///
/// Next, select the Information Property List item, select Add Item from the
/// Editor menu, then select Localizations from the pop-up menu.
///
/// Select and expand the newly-created Localizations item then, for each
/// locale your application supports, add a new item and select the locale
/// you wish to add from the pop-up menu in the Value field. This list should
/// be consistent with the languages listed in the AppLocalizations.supportedLocales
/// property.
abstract class AppLocalizations {
  AppLocalizations(String locale)
    : localeName = intl.Intl.canonicalizedLocale(locale.toString());

  final String localeName;

  static AppLocalizations? of(BuildContext context) {
    return Localizations.of<AppLocalizations>(context, AppLocalizations);
  }

  static const LocalizationsDelegate<AppLocalizations> delegate =
      _AppLocalizationsDelegate();

  /// A list of this localizations delegate along with the default localizations
  /// delegates.
  ///
  /// Returns a list of localizations delegates containing this delegate along with
  /// GlobalMaterialLocalizations.delegate, GlobalCupertinoLocalizations.delegate,
  /// and GlobalWidgetsLocalizations.delegate.
  ///
  /// Additional delegates can be added by appending to this list in
  /// MaterialApp. This list does not have to be used at all if a custom list
  /// of delegates is preferred or required.
  static const List<LocalizationsDelegate<dynamic>> localizationsDelegates =
      <LocalizationsDelegate<dynamic>>[
        delegate,
        GlobalMaterialLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
      ];

  /// A list of this localizations delegate's supported locales.
  static const List<Locale> supportedLocales = <Locale>[
    Locale('en'),
    Locale('zh'),
  ];

  /// No description provided for @appTitle.
  ///
  /// In zh, this message translates to:
  /// **'短信平台'**
  String get appTitle;

  /// No description provided for @login.
  ///
  /// In zh, this message translates to:
  /// **'登录'**
  String get login;

  /// No description provided for @register.
  ///
  /// In zh, this message translates to:
  /// **'注册'**
  String get register;

  /// No description provided for @username.
  ///
  /// In zh, this message translates to:
  /// **'用户名'**
  String get username;

  /// No description provided for @password.
  ///
  /// In zh, this message translates to:
  /// **'密码'**
  String get password;

  /// No description provided for @email.
  ///
  /// In zh, this message translates to:
  /// **'邮箱'**
  String get email;

  /// No description provided for @loginTitle.
  ///
  /// In zh, this message translates to:
  /// **'登录短信平台'**
  String get loginTitle;

  /// No description provided for @registerTitle.
  ///
  /// In zh, this message translates to:
  /// **'注册新账号'**
  String get registerTitle;

  /// No description provided for @noAccount.
  ///
  /// In zh, this message translates to:
  /// **'没有账号？注册'**
  String get noAccount;

  /// No description provided for @hasAccount.
  ///
  /// In zh, this message translates to:
  /// **'已有账号？登录'**
  String get hasAccount;

  /// No description provided for @home.
  ///
  /// In zh, this message translates to:
  /// **'主页'**
  String get home;

  /// No description provided for @profile.
  ///
  /// In zh, this message translates to:
  /// **'个人中心'**
  String get profile;

  /// No description provided for @whitelist.
  ///
  /// In zh, this message translates to:
  /// **'白名单'**
  String get whitelist;

  /// No description provided for @settings.
  ///
  /// In zh, this message translates to:
  /// **'设置'**
  String get settings;

  /// No description provided for @balance.
  ///
  /// In zh, this message translates to:
  /// **'当前余额'**
  String get balance;

  /// No description provided for @recharge.
  ///
  /// In zh, this message translates to:
  /// **'充值'**
  String get recharge;

  /// No description provided for @assignPhone.
  ///
  /// In zh, this message translates to:
  /// **'手机号分配'**
  String get assignPhone;

  /// No description provided for @getPhone.
  ///
  /// In zh, this message translates to:
  /// **'拉取手机号'**
  String get getPhone;

  /// No description provided for @getCode.
  ///
  /// In zh, this message translates to:
  /// **'拉取验证码'**
  String get getCode;

  /// No description provided for @history.
  ///
  /// In zh, this message translates to:
  /// **'分配历史'**
  String get history;

  /// No description provided for @recentAssignments.
  ///
  /// In zh, this message translates to:
  /// **'最近分配记录'**
  String get recentAssignments;

  /// No description provided for @refreshCode.
  ///
  /// In zh, this message translates to:
  /// **'刷新验证码'**
  String get refreshCode;

  /// No description provided for @businessType.
  ///
  /// In zh, this message translates to:
  /// **'业务类型'**
  String get businessType;

  /// No description provided for @cardType.
  ///
  /// In zh, this message translates to:
  /// **'卡类型'**
  String get cardType;

  /// No description provided for @phone.
  ///
  /// In zh, this message translates to:
  /// **'手机号'**
  String get phone;

  /// No description provided for @code.
  ///
  /// In zh, this message translates to:
  /// **'验证码'**
  String get code;

  /// No description provided for @cost.
  ///
  /// In zh, this message translates to:
  /// **'费用'**
  String get cost;

  /// No description provided for @status.
  ///
  /// In zh, this message translates to:
  /// **'状态'**
  String get status;

  /// No description provided for @time.
  ///
  /// In zh, this message translates to:
  /// **'时间'**
  String get time;

  /// No description provided for @completed.
  ///
  /// In zh, this message translates to:
  /// **'已完成'**
  String get completed;

  /// No description provided for @expired.
  ///
  /// In zh, this message translates to:
  /// **'已过期'**
  String get expired;

  /// No description provided for @apiKey.
  ///
  /// In zh, this message translates to:
  /// **'API密钥'**
  String get apiKey;

  /// No description provided for @lastLogin.
  ///
  /// In zh, this message translates to:
  /// **'上次登录'**
  String get lastLogin;

  /// No description provided for @regIp.
  ///
  /// In zh, this message translates to:
  /// **'注册IP'**
  String get regIp;

  /// No description provided for @changePassword.
  ///
  /// In zh, this message translates to:
  /// **'修改密码'**
  String get changePassword;

  /// No description provided for @logout.
  ///
  /// In zh, this message translates to:
  /// **'退出登录'**
  String get logout;

  /// No description provided for @language.
  ///
  /// In zh, this message translates to:
  /// **'界面语言'**
  String get language;

  /// No description provided for @theme.
  ///
  /// In zh, this message translates to:
  /// **'主题模式'**
  String get theme;

  /// No description provided for @light.
  ///
  /// In zh, this message translates to:
  /// **'浅色'**
  String get light;

  /// No description provided for @dark.
  ///
  /// In zh, this message translates to:
  /// **'深色'**
  String get dark;

  /// No description provided for @loading.
  ///
  /// In zh, this message translates to:
  /// **'加载中'**
  String get loading;

  /// No description provided for @pleaseWait.
  ///
  /// In zh, this message translates to:
  /// **'请稍候...'**
  String get pleaseWait;

  /// No description provided for @error.
  ///
  /// In zh, this message translates to:
  /// **'出错了'**
  String get error;

  /// No description provided for @noData.
  ///
  /// In zh, this message translates to:
  /// **'暂无数据'**
  String get noData;

  /// No description provided for @noDataMessage.
  ///
  /// In zh, this message translates to:
  /// **'当前没有可显示的内容。'**
  String get noDataMessage;

  /// No description provided for @back.
  ///
  /// In zh, this message translates to:
  /// **'返回'**
  String get back;

  /// No description provided for @add.
  ///
  /// In zh, this message translates to:
  /// **'新增'**
  String get add;

  /// No description provided for @delete.
  ///
  /// In zh, this message translates to:
  /// **'删除'**
  String get delete;

  /// No description provided for @ipAddress.
  ///
  /// In zh, this message translates to:
  /// **'IP地址'**
  String get ipAddress;

  /// No description provided for @notes.
  ///
  /// In zh, this message translates to:
  /// **'备注'**
  String get notes;

  /// No description provided for @usernameRequired.
  ///
  /// In zh, this message translates to:
  /// **'请输入用户名'**
  String get usernameRequired;

  /// No description provided for @passwordRequired.
  ///
  /// In zh, this message translates to:
  /// **'请输入密码'**
  String get passwordRequired;

  /// No description provided for @emailRequired.
  ///
  /// In zh, this message translates to:
  /// **'请输入邮箱'**
  String get emailRequired;

  /// No description provided for @noAccountRegister.
  ///
  /// In zh, this message translates to:
  /// **'没有账号？立即注册'**
  String get noAccountRegister;

  /// No description provided for @getPhoneTitle.
  ///
  /// In zh, this message translates to:
  /// **'手机号分配'**
  String get getPhoneTitle;

  /// No description provided for @getCodeTitle.
  ///
  /// In zh, this message translates to:
  /// **'获取验证码'**
  String get getCodeTitle;

  /// No description provided for @selectBusinessType.
  ///
  /// In zh, this message translates to:
  /// **'选择业务类型'**
  String get selectBusinessType;

  /// No description provided for @selectCardType.
  ///
  /// In zh, this message translates to:
  /// **'选择卡类型'**
  String get selectCardType;

  /// No description provided for @physicalCard.
  ///
  /// In zh, this message translates to:
  /// **'实体卡'**
  String get physicalCard;

  /// No description provided for @virtualCard.
  ///
  /// In zh, this message translates to:
  /// **'虚拟卡'**
  String get virtualCard;

  /// No description provided for @assignNow.
  ///
  /// In zh, this message translates to:
  /// **'立即分配'**
  String get assignNow;

  /// No description provided for @assignSuccess.
  ///
  /// In zh, this message translates to:
  /// **'分配成功'**
  String get assignSuccess;

  /// No description provided for @phoneNumber.
  ///
  /// In zh, this message translates to:
  /// **'手机号码'**
  String get phoneNumber;

  /// No description provided for @enterPhoneNumber.
  ///
  /// In zh, this message translates to:
  /// **'请输入手机号'**
  String get enterPhoneNumber;

  /// No description provided for @getCodeNow.
  ///
  /// In zh, this message translates to:
  /// **'获取验证码'**
  String get getCodeNow;

  /// No description provided for @codeReceived.
  ///
  /// In zh, this message translates to:
  /// **'验证码已获取'**
  String get codeReceived;

  /// No description provided for @copyPhone.
  ///
  /// In zh, this message translates to:
  /// **'复制手机号'**
  String get copyPhone;

  /// No description provided for @copyCode.
  ///
  /// In zh, this message translates to:
  /// **'复制验证码'**
  String get copyCode;

  /// No description provided for @copied.
  ///
  /// In zh, this message translates to:
  /// **'已复制'**
  String get copied;

  /// No description provided for @businessTypeRequired.
  ///
  /// In zh, this message translates to:
  /// **'请选择业务类型'**
  String get businessTypeRequired;

  /// No description provided for @cardTypeRequired.
  ///
  /// In zh, this message translates to:
  /// **'请选择卡类型'**
  String get cardTypeRequired;

  /// No description provided for @phoneRequired.
  ///
  /// In zh, this message translates to:
  /// **'请输入手机号'**
  String get phoneRequired;

  /// No description provided for @gettingCode.
  ///
  /// In zh, this message translates to:
  /// **'获取验证码中，请稍候...'**
  String get gettingCode;

  /// No description provided for @codeTimeout.
  ///
  /// In zh, this message translates to:
  /// **'验证码获取超时'**
  String get codeTimeout;

  /// No description provided for @noBusinessTypes.
  ///
  /// In zh, this message translates to:
  /// **'暂无可用业务类型'**
  String get noBusinessTypes;

  /// No description provided for @loadingBusinessTypes.
  ///
  /// In zh, this message translates to:
  /// **'加载业务类型中...'**
  String get loadingBusinessTypes;

  /// No description provided for @assignmentResult.
  ///
  /// In zh, this message translates to:
  /// **'分配结果'**
  String get assignmentResult;

  /// No description provided for @verificationCode.
  ///
  /// In zh, this message translates to:
  /// **'验证码'**
  String get verificationCode;

  /// No description provided for @waitingForCode.
  ///
  /// In zh, this message translates to:
  /// **'等待验证码中...'**
  String get waitingForCode;

  /// No description provided for @retryGetCode.
  ///
  /// In zh, this message translates to:
  /// **'重试获取'**
  String get retryGetCode;
}

class _AppLocalizationsDelegate
    extends LocalizationsDelegate<AppLocalizations> {
  const _AppLocalizationsDelegate();

  @override
  Future<AppLocalizations> load(Locale locale) {
    return SynchronousFuture<AppLocalizations>(lookupAppLocalizations(locale));
  }

  @override
  bool isSupported(Locale locale) =>
      <String>['en', 'zh'].contains(locale.languageCode);

  @override
  bool shouldReload(_AppLocalizationsDelegate old) => false;
}

AppLocalizations lookupAppLocalizations(Locale locale) {
  // Lookup logic when only language code is specified.
  switch (locale.languageCode) {
    case 'en':
      return AppLocalizationsEn();
    case 'zh':
      return AppLocalizationsZh();
  }

  throw FlutterError(
    'AppLocalizations.delegate failed to load unsupported locale "$locale". This is likely '
    'an issue with the localizations generation tool. Please file an issue '
    'on GitHub with a reproducible sample app and the gen-l10n configuration '
    'that was used.',
  );
}
