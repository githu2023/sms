# 短信平台客户端 Flutter 架构说明

## 1. 项目结构

```
sms_platform_flutter/
├── lib/
│   ├── main.dart                # 应用入口
│   ├── app.dart                 # App根组件，主题/路由/多语言配置
│   ├── l10n/                    # 国际化资源
│   ├── core/                    # 核心工具/全局配置
│   ├── models/                  # 数据模型
│   ├── pages/                   # 各页面
│   ├── widgets/                 # 通用组件
│   ├── providers/               # 状态管理
│   └── routes.dart              # 路由表
├── pubspec.yaml                 # 依赖声明
├── assets/                      # 图片、图标等资源
└── README.md
```

## 2. 主要模块说明

- **l10n/**：多语言资源，使用 `intl` 包，支持全局切换语言。
- **core/**：网络请求、主题、常量、全局状态等工具类。
- **models/**：所有后端API相关的数据结构。
- **pages/**：每个页面一个文件，保持单一职责。
- **widgets/**：可复用的UI组件，如底部导航栏、卡片、弹窗等。
- **providers/**：使用 Provider/Riverpod 管理全局和局部状态。
- **routes.dart**：统一管理页面路由，支持命名路由和参数传递。

## 3. 路由与导航
- 推荐使用 `go_router` 或 `Navigator 2.0`，支持页面跳转、底部导航栏。

## 4. 国际化与主题
- 国际化采用 `flutter_localizations` + `intl`，所有文案抽离到 l10n 文件夹。
- 主题支持浅色/深色切换，统一在 `core/theme.dart` 配置。

## 5. 状态管理
- 推荐 Provider 或 Riverpod，所有全局状态（如用户、语言、主题）集中管理。

## 6. 网络层
- 所有API请求统一在 `core/api.dart` 封装，支持错误处理、加载动画、空状态。

## 7. 代码规范
- 见 `CODE_STYLE.md`。

## 8. 其它
- 所有页面均有底部导航栏，风格统一。
- 资源文件（图片、图标）统一放在 assets 目录。

---
如需详细模块说明或开发约定，请查阅对应目录下的文档或联系项目负责人。
