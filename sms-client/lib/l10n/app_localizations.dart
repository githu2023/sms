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
  AppLocalizations(String locale) : localeName = intl.Intl.canonicalizedLocale(locale.toString());

  final String localeName;

  static AppLocalizations? of(BuildContext context) {
    return Localizations.of<AppLocalizations>(context, AppLocalizations);
  }

  static const LocalizationsDelegate<AppLocalizations> delegate = _AppLocalizationsDelegate();

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
  static const List<LocalizationsDelegate<dynamic>> localizationsDelegates = <LocalizationsDelegate<dynamic>>[
    delegate,
    GlobalMaterialLocalizations.delegate,
    GlobalCupertinoLocalizations.delegate,
    GlobalWidgetsLocalizations.delegate,
  ];

  /// A list of this localizations delegate's supported locales.
  static const List<Locale> supportedLocales = <Locale>[
    Locale('en'),
    Locale('zh')
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

  /// No description provided for @frozenAmount.
  ///
  /// In zh, this message translates to:
  /// **'冻结金额'**
  String get frozenAmount;

  /// No description provided for @availableBalance.
  ///
  /// In zh, this message translates to:
  /// **'可用余额'**
  String get availableBalance;

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

  /// No description provided for @noCodeReceived.
  ///
  /// In zh, this message translates to:
  /// **'未收到验证码'**
  String get noCodeReceived;

  /// No description provided for @retryGetCode.
  ///
  /// In zh, this message translates to:
  /// **'重试获取'**
  String get retryGetCode;

  /// No description provided for @requestCount.
  ///
  /// In zh, this message translates to:
  /// **'请求数量'**
  String get requestCount;

  /// No description provided for @requestCountHelper.
  ///
  /// In zh, this message translates to:
  /// **'每次最多10个手机号'**
  String get requestCountHelper;

  /// No description provided for @totalCostLabel.
  ///
  /// In zh, this message translates to:
  /// **'总费用'**
  String get totalCostLabel;

  /// No description provided for @remainingBalanceLabel.
  ///
  /// In zh, this message translates to:
  /// **'剩余余额'**
  String get remainingBalanceLabel;

  /// No description provided for @successCountLabel.
  ///
  /// In zh, this message translates to:
  /// **'成功数量'**
  String get successCountLabel;

  /// No description provided for @failedCountLabel.
  ///
  /// In zh, this message translates to:
  /// **'失败数量'**
  String get failedCountLabel;

  /// No description provided for @pendingCountLabel.
  ///
  /// In zh, this message translates to:
  /// **'等待数量'**
  String get pendingCountLabel;

  /// No description provided for @countryCode.
  ///
  /// In zh, this message translates to:
  /// **'国家/地区'**
  String get countryCode;

  /// No description provided for @providerLabel.
  ///
  /// In zh, this message translates to:
  /// **'服务商'**
  String get providerLabel;

  /// No description provided for @validUntilLabel.
  ///
  /// In zh, this message translates to:
  /// **'有效期至'**
  String get validUntilLabel;

  /// No description provided for @codesSummary.
  ///
  /// In zh, this message translates to:
  /// **'验证码统计'**
  String get codesSummary;

  /// No description provided for @enterPhonesHint.
  ///
  /// In zh, this message translates to:
  /// **'支持换行或逗号分隔输入多个手机号'**
  String get enterPhonesHint;

  /// No description provided for @enterPhonesHelper.
  ///
  /// In zh, this message translates to:
  /// **'最多同时查询10个手机号'**
  String get enterPhonesHelper;

  /// No description provided for @invalidPhoneInput.
  ///
  /// In zh, this message translates to:
  /// **'请输入1-10个有效手机号'**
  String get invalidPhoneInput;

  /// No description provided for @noCodesYet.
  ///
  /// In zh, this message translates to:
  /// **'暂未获取到验证码，若为等待状态请继续轮询'**
  String get noCodesYet;

  /// No description provided for @messageLabel.
  ///
  /// In zh, this message translates to:
  /// **'消息'**
  String get messageLabel;

  /// No description provided for @receivedAtLabel.
  ///
  /// In zh, this message translates to:
  /// **'接收时间'**
  String get receivedAtLabel;

  /// No description provided for @statusSuccess.
  ///
  /// In zh, this message translates to:
  /// **'成功'**
  String get statusSuccess;

  /// No description provided for @statusPending.
  ///
  /// In zh, this message translates to:
  /// **'等待'**
  String get statusPending;

  /// No description provided for @statusActive.
  ///
  /// In zh, this message translates to:
  /// **'激活中'**
  String get statusActive;

  /// No description provided for @statusFailed.
  ///
  /// In zh, this message translates to:
  /// **'失败'**
  String get statusFailed;

  /// No description provided for @getCodeInstructions.
  ///
  /// In zh, this message translates to:
  /// **'可同时查询最多10个手机号，若状态为等待请继续轮询此接口。'**
  String get getCodeInstructions;

  /// No description provided for @currentPasswordLabel.
  ///
  /// In zh, this message translates to:
  /// **'当前密码'**
  String get currentPasswordLabel;

  /// No description provided for @newPasswordLabel.
  ///
  /// In zh, this message translates to:
  /// **'新密码'**
  String get newPasswordLabel;

  /// No description provided for @confirmPasswordLabel.
  ///
  /// In zh, this message translates to:
  /// **'确认新密码'**
  String get confirmPasswordLabel;

  /// No description provided for @passwordTooShort.
  ///
  /// In zh, this message translates to:
  /// **'密码至少需要6位'**
  String get passwordTooShort;

  /// No description provided for @passwordUnchanged.
  ///
  /// In zh, this message translates to:
  /// **'新密码不能与旧密码相同'**
  String get passwordUnchanged;

  /// No description provided for @passwordMismatch.
  ///
  /// In zh, this message translates to:
  /// **'两次输入的密码不一致'**
  String get passwordMismatch;

  /// No description provided for @changePasswordSuccess.
  ///
  /// In zh, this message translates to:
  /// **'密码修改成功'**
  String get changePasswordSuccess;

  /// No description provided for @changePasswordFailed.
  ///
  /// In zh, this message translates to:
  /// **'密码修改失败'**
  String get changePasswordFailed;

  /// No description provided for @addWhitelistSuccess.
  ///
  /// In zh, this message translates to:
  /// **'白名单添加成功'**
  String get addWhitelistSuccess;

  /// No description provided for @addWhitelistFailed.
  ///
  /// In zh, this message translates to:
  /// **'白名单添加失败'**
  String get addWhitelistFailed;

  /// No description provided for @deleteWhitelistSuccess.
  ///
  /// In zh, this message translates to:
  /// **'白名单删除成功'**
  String get deleteWhitelistSuccess;

  /// No description provided for @deleteWhitelistFailed.
  ///
  /// In zh, this message translates to:
  /// **'白名单删除失败'**
  String get deleteWhitelistFailed;

  /// No description provided for @deleteWhitelistConfirm.
  ///
  /// In zh, this message translates to:
  /// **'确认删除以下IP？'**
  String get deleteWhitelistConfirm;

  /// No description provided for @notesOptional.
  ///
  /// In zh, this message translates to:
  /// **'备注（可选）'**
  String get notesOptional;

  /// No description provided for @ipAddressRequired.
  ///
  /// In zh, this message translates to:
  /// **'请输入IP地址或CIDR网段'**
  String get ipAddressRequired;

  /// No description provided for @invalidIpFormat.
  ///
  /// In zh, this message translates to:
  /// **'IP地址或CIDR格式不正确'**
  String get invalidIpFormat;

  /// No description provided for @viewAll.
  ///
  /// In zh, this message translates to:
  /// **'查看全部'**
  String get viewAll;

  /// No description provided for @recentRecords.
  ///
  /// In zh, this message translates to:
  /// **'最近记录'**
  String get recentRecords;

  /// No description provided for @pullDownToRefresh.
  ///
  /// In zh, this message translates to:
  /// **'下拉刷新，上滑加载更多'**
  String get pullDownToRefresh;

  /// No description provided for @noMoreData.
  ///
  /// In zh, this message translates to:
  /// **'没有更多了'**
  String get noMoreData;

  /// No description provided for @loadMore.
  ///
  /// In zh, this message translates to:
  /// **'加载更多'**
  String get loadMore;

  /// No description provided for @copy.
  ///
  /// In zh, this message translates to:
  /// **'复制'**
  String get copy;

  /// No description provided for @totalRecords.
  ///
  /// In zh, this message translates to:
  /// **'共 {count} 条记录'**
  String totalRecords(int count);

  /// No description provided for @refresh.
  ///
  /// In zh, this message translates to:
  /// **'刷新'**
  String get refresh;

  /// No description provided for @loadingFailed.
  ///
  /// In zh, this message translates to:
  /// **'加载失败'**
  String get loadingFailed;

  /// No description provided for @noAvailableBusinessTypes.
  ///
  /// In zh, this message translates to:
  /// **'暂无可用的业务类型，请联系管理员配置'**
  String get noAvailableBusinessTypes;

  /// No description provided for @getBusinessTypesFailed.
  ///
  /// In zh, this message translates to:
  /// **'获取业务类型失败'**
  String get getBusinessTypesFailed;

  /// No description provided for @getFailed.
  ///
  /// In zh, this message translates to:
  /// **'获取失败'**
  String get getFailed;

  /// No description provided for @registerFailed.
  ///
  /// In zh, this message translates to:
  /// **'注册失败'**
  String get registerFailed;

  /// No description provided for @registerSuccess.
  ///
  /// In zh, this message translates to:
  /// **'注册成功！请登录'**
  String get registerSuccess;
}

class _AppLocalizationsDelegate extends LocalizationsDelegate<AppLocalizations> {
  const _AppLocalizationsDelegate();

  @override
  Future<AppLocalizations> load(Locale locale) {
    return SynchronousFuture<AppLocalizations>(lookupAppLocalizations(locale));
  }

  @override
  bool isSupported(Locale locale) => <String>['en', 'zh'].contains(locale.languageCode);

  @override
  bool shouldReload(_AppLocalizationsDelegate old) => false;
}

AppLocalizations lookupAppLocalizations(Locale locale) {


  // Lookup logic when only language code is specified.
  switch (locale.languageCode) {
    case 'en': return AppLocalizationsEn();
    case 'zh': return AppLocalizationsZh();
  }

  throw FlutterError(
    'AppLocalizations.delegate failed to load unsupported locale "$locale". This is likely '
    'an issue with the localizations generation tool. Please file an issue '
    'on GitHub with a reproducible sample app and the gen-l10n configuration '
    'that was used.'
  );
}
