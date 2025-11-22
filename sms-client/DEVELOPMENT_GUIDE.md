# 快速开发指南

## 项目概览

SMS Platform Client 是一个功能完整的 Flutter 跨平台应用，已实现核心功能框架。

## 当前状态

### ✅ 已完成 (可直接使用)

1. **完整的项目架构**
   - Provider 状态管理
   - 模块化代码组织
   - 多语言支持框架
   - 主题系统

2. **用户认证流程**
   - 登录页面 + 表单验证
   - 注册页面 + 邮箱验证
   - Token 管理
   - 自动登录

3. **主页框架**
   - 三标签导航 (主页/历史/设置)
   - 余额显示卡片
   - 快捷操作按钮
   - 下拉刷新

4. **设置功能**
   - 语言切换 (中文/English)
   - 主题切换 (浅色/深色)
   - 白名单管理
   - 修改密码

5. **网络层**
   - 统一 API 客户端
   - 请求/响应封装
   - 错误处理
   - Token 自动注入

## 下一步开发建议

### 优先级 1: 核心业务功能

#### 1. 手机号分配功能
创建 `lib/pages/get_phone_page.dart`:
```dart
- 业务类型选择器
- 卡类型选择器  
- 分配手机号按钮
- 显示分配结果
```

#### 2. 验证码获取功能
创建 `lib/pages/get_code_page.dart`:
```dart
- 输入手机号
- 获取验证码按钮
- 显示验证码
- 刷新验证码
```

#### 3. 历史记录
完善 `home_page.dart` 的历史标签:
```dart
- 分配记录列表
- 筛选功能
- 分页加载
- 详情查看
```

### 优先级 2: API 集成

#### 更新 `lib/core/api_client.dart`

添加业务 API 方法:

```dart
// 获取业务类型列表
Future<ApiResponse<List<BusinessType>>> getBusinessTypes();

// 分配手机号
Future<ApiResponse<Assignment>> assignPhone({
  required String businessType,
  required String cardType,
});

// 获取验证码
Future<ApiResponse<String>> getVerificationCode({
  required String phone,
});

// 获取历史记录
Future<ApiResponse<List<Assignment>>> getAssignments({
  int page = 1,
  int limit = 20,
});

// 白名单操作
Future<ApiResponse<List<Whitelist>>> getWhitelists();
Future<ApiResponse<void>> addWhitelist(String ip, String? notes);
Future<ApiResponse<void>> deleteWhitelist(String id);
```

#### 创建对应数据模型

`lib/models/`:
- `business_type.dart` - 业务类型
- `assignment.dart` - 分配记录
- `transaction.dart` - 交易记录

### 优先级 3: UI 优化

1. **添加加载状态**
   - 全局加载指示器
   - 按钮加载动画
   - 骨架屏

2. **错误处理**
   - 网络错误提示
   - 重试机制
   - 友好的错误页面

3. **空状态**
   - 暂无数据提示
   - 引导用户操作

## 文件组织建议

```
lib/
├── core/
│   ├── api_client.dart      ✅ 已完成
│   ├── theme.dart           ✅ 已完成
│   ├── constants.dart       ⏳ 建议添加 (常量定义)
│   └── utils.dart           ⏳ 建议添加 (工具函数)
│
├── models/
│   ├── user.dart            ✅ 已完成
│   ├── api_response.dart    ✅ 已完成
│   ├── whitelist.dart       ✅ 已完成
│   ├── business_type.dart   ⏳ 待创建
│   ├── assignment.dart      ⏳ 待创建
│   └── transaction.dart     ⏳ 待创建
│
├── pages/
│   ├── login_page.dart      ✅ 已完成
│   ├── register_page.dart   ✅ 已完成
│   ├── home_page.dart       ✅ 已完成 (框架)
│   ├── whitelist_page.dart  ✅ 已完成
│   ├── change_password_page.dart ✅ 已完成
│   ├── get_phone_page.dart  ⏳ 待创建
│   ├── get_code_page.dart   ⏳ 待创建
│   └── assignment_detail_page.dart ⏳ 待创建
│
├── widgets/
│   ├── language_picker.dart ✅ 已完成
│   ├── theme_picker.dart    ✅ 已完成
│   ├── loading_widget.dart  ⏳ 建议添加
│   ├── empty_widget.dart    ⏳ 建议添加
│   └── error_widget.dart    ⏳ 建议添加
│
└── providers/
    ├── auth_provider.dart   ✅ 已完成
    ├── locale_provider.dart ✅ 已完成
    ├── theme_provider.dart  ✅ 已完成
    └── assignment_provider.dart ⏳ 建议添加
```

## 快速开发模板

### 创建新页面模板

```dart
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../l10n/app_localizations.dart';

class NewPage extends StatefulWidget {
  const NewPage({super.key});

  @override
  State<NewPage> createState() => _NewPageState();
}

class _NewPageState extends State<NewPage> {
  bool _isLoading = false;

  @override
  void initState() {
    super.initState();
    _loadData();
  }

  Future<void> _loadData() async {
    setState(() => _isLoading = true);
    // TODO: Load data
    setState(() => _isLoading = false);
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    
    return Scaffold(
      appBar: AppBar(
        title: Text('Page Title'),
      ),
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : _buildContent(),
    );
  }

  Widget _buildContent() {
    return SingleChildScrollView(
      padding: const EdgeInsets.all(16),
      child: Column(
        children: [
          // Your content here
        ],
      ),
    );
  }
}
```

### 添加 API 方法模板

```dart
// 在 api_client.dart 中添加
Future<ApiResponse<YourModel>> yourMethod() async {
  return await get(
    '/your/endpoint',
    fromJson: (data) => YourModel.fromJson(data),
  );
}
```

### 创建 Provider 模板

```dart
import 'package:flutter/foundation.dart';
import '../core/api_client.dart';
import '../models/your_model.dart';

class YourProvider extends ChangeNotifier {
  final ApiClient _apiClient = ApiClient();
  
  List<YourModel> _items = [];
  bool _isLoading = false;
  String? _error;

  List<YourModel> get items => _items;
  bool get isLoading => _isLoading;
  String? get error => _error;

  Future<void> loadItems() async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      final response = await _apiClient.yourMethod();
      if (response.success && response.data != null) {
        _items = response.data!;
      } else {
        _error = response.message;
      }
    } catch (e) {
      _error = 'Network error: $e';
    }

    _isLoading = false;
    notifyListeners();
  }
}
```

## 测试建议

### 1. 启动后端 API

确保后端服务运行在 `http://localhost:8080`

### 2. 测试流程

```bash
# 1. 启动应用
flutter run -d chrome

# 2. 测试功能
- 注册新用户
- 登录
- 切换语言
- 切换主题
- 白名单管理
- 修改密码
- 退出登录
```

### 3. 调试技巧

```dart
// 添加调试日志
debugPrint('API Response: $response');

// 使用 DevTools
flutter pub global activate devtools
flutter pub global run devtools
```

## 常用命令

```bash
# 代码检查
flutter analyze

# 代码格式化
dart format lib/

# 生成本地化
flutter gen-l10n

# 清理构建
flutter clean

# 获取依赖
flutter pub get

# 运行应用
flutter run

# 构建 APK
flutter build apk --release

# 构建 Web
flutter build web --release
```

## 注意事项

1. **API 地址配置**
   - 开发: `http://localhost:8080/api`
   - 生产: 修改 `lib/core/api_client.dart` 中的 `baseUrl`

2. **Token 管理**
   - Token 自动存储在 SharedPreferences
   - 自动添加到请求头
   - 退出登录时清除

3. **错误处理**
   - 所有 API 调用已封装错误处理
   - 使用 `ApiResponse` 统一响应格式
   - 通过 Provider 管理错误状态

4. **性能优化**
   - 使用 `const` 构造函数
   - 避免不必要的 `setState`
   - 合理使用 `Provider` 的 `select`

## 获取帮助

- Flutter 文档: https://flutter.dev/docs
- Provider 文档: https://pub.dev/packages/provider
- Material Design: https://m3.material.io
