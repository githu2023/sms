# SMS Platform Client

一个基于 Flutter 开发的跨平台短信平台客户端应用，支持 iOS、Android 和 Web。

## 功能特性

### ✅ 已完成功能

#### 1. 用户认证
- ✅ 用户注册
- ✅ 用户登录
- ✅ 自动登录状态保持
- ✅ 安全退出登录

#### 2. 多语言支持
- ✅ 中文界面
- ✅ 英文界面
- ✅ 实时语言切换
- ✅ 语言偏好持久化

#### 3. 主题定制
- ✅ 浅色模式
- ✅ 深色模式
- ✅ Google Material Design 配色
- ✅ 主题偏好持久化

#### 4. 主页功能
- ✅ 余额显示
- ✅ 充值入口
- ✅ 快捷操作（拉取手机号、拉取验证码）
- ✅ 最近记录展示
- ✅ 下拉刷新

#### 5. 个人中心
- ✅ 用户信息展示
- ✅ API密钥管理
- ✅ 修改密码
- ✅ 语言设置
- ✅ 主题设置
- ✅ 白名单管理

#### 6. 白名单管理
- ✅ IP白名单列表
- ✅ 添加IP
- ✅ 删除IP
- ✅ 备注信息

### 🚧 待开发功能

- ⏳ 业务类型选择
- ⏳ 手机号分配
- ⏳ 验证码获取
- ⏳ 历史记录查询
- ⏳ 交易记录
- ⏳ 充值功能

## 技术栈

- **框架**: Flutter 3.8+
- **语言**: Dart
- **状态管理**: Provider
- **网络请求**: http
- **本地存储**: SharedPreferences
- **国际化**: flutter_localizations + intl
- **UI设计**: Material Design 3

## 项目结构

```
lib/
├── core/                 # 核心功能
│   ├── api_client.dart  # API客户端封装
│   └── theme.dart       # 主题配置
├── models/              # 数据模型
│   ├── api_response.dart
│   ├── user.dart
│   └── whitelist.dart
├── pages/               # 页面
│   ├── login_page.dart
│   ├── register_page.dart
│   ├── home_page.dart
│   ├── whitelist_page.dart
│   └── change_password_page.dart
├── providers/           # 状态管理
│   ├── auth_provider.dart
│   ├── locale_provider.dart
│   └── theme_provider.dart
├── widgets/             # 通用组件
│   ├── language_picker.dart
│   └── theme_picker.dart
├── l10n/               # 国际化资源
│   ├── intl_zh.arb
│   ├── intl_en.arb
│   └── app_localizations.dart
├── app.dart            # 应用根组件
└── main.dart           # 入口文件
```

## 开始使用

### 环境要求

- Flutter SDK: 3.8.0 或更高版本
- Dart SDK: 3.0.0 或更高版本
- iOS: Xcode 12.0+
- Android: Android Studio / Gradle

### 安装依赖

```bash
cd sms-client
flutter pub get
```

### 运行应用

```bash
# iOS
flutter run -d ios

# Android  
flutter run -d android

# Web
flutter run -d chrome

# 所有平台
flutter run
```

### 构建发布版本

```bash
# iOS
flutter build ios --release

# Android APK
flutter build apk --release

# Android App Bundle
flutter build appbundle --release

# Web
flutter build web --release
```

## API 配置

默认 API 地址配置在 `lib/core/api_client.dart`:

```dart
static const String baseUrl = 'http://localhost:8080/api';
```

发布前请修改为实际的生产环境地址。

## 国际化

### 添加新语言

1. 在 `lib/l10n/` 目录创建新的 ARB 文件，如 `intl_ja.arb`
2. 复制现有翻译文件内容并翻译
3. 在 `lib/app.dart` 中添加新语言支持:

```dart
supportedLocales: const [
  Locale('zh'),
  Locale('en'),
  Locale('ja'), // 新增
],
```

4. 运行 `flutter gen-l10n` 生成翻译代码

### 添加新翻译键

1. 在 `intl_zh.arb` 和 `intl_en.arb` 中添加新键值
2. 运行 `flutter gen-l10n` 重新生成
3. 在代码中使用: `AppLocalizations.of(context)!.yourKey`

## 主题定制

主题配置位于 `lib/core/theme.dart`，可自定义:

- 主色调
- 卡片样式
- 按钮样式
- 输入框样式
- 文本样式

## 开发规范

### 代码风格

- 遵循 Dart 官方代码规范
- 使用 `flutter analyze` 检查代码质量
- 变量命名采用驼峰命名法
- 类名首字母大写

### 提交规范

- feat: 新功能
- fix: 修复bug
- docs: 文档更新
- style: 代码格式调整
- refactor: 重构
- test: 测试相关
- chore: 构建/工具链相关

## 常见问题

### 1. 本地化代码未生成

运行以下命令生成:
```bash
flutter gen-l10n
```

### 2. 依赖冲突

清理缓存并重新安装:
```bash
flutter clean
flutter pub get
```

### 3. iOS构建失败

更新CocoaPods:
```bash
cd ios
pod install
cd ..
```

## 许可证

本项目采用 MIT 许可证。

## 联系方式

如有问题或建议，请提交 Issue。
