# Go API - Development Guidelines

本文档包含了 `goapi` 子项目在开发过程中需要遵循的规范，包括日志、配置和代码风格，旨在确保代码质量和团队协作效率。所有开发者（包括AI助手）都应遵守此规范。

帐号 testuser_1763969884 
商户号 368570
api key 431915928ea2df20f0a3af17921dac2efee5c63d46191927ac781e528e3fd0a5
---

## 1. 日志 (Logging)

统一的日志规范对于问题排查和系统监控至关重要。

*   **推荐库**: `Zap` (by Uber)。
*   **日志格式**: **必须使用结构化日志 (Structured Logging)**，输出为 **JSON** 格式，方便后续被日志系统收集和检索。
    *   **示例**:
        ```json
        {"level":"error", "ts":1668853335.779, "caller":"service/user_service.go:42", "msg":"Failed to create user", "error":"email already exists", "request_id":"b9v2h5g2a9d8f7h6g"}
        ```
*   **日志级别**:
    *   `Debug`: 开发调试时使用，生产环境通常关闭。
    *   `Info`: 记录关键的业务流程信息，如“用户登录成功”。
    *   `Warn`: 出现可恢复的、不影响主流程的异常情况。
    *   `Error`: 发生错误，影响了本次操作，需要关注。
    *   `Fatal`: 导致程序崩溃的严重错误。
*   **日志上下文 (Context)**:
    *   应在 `Gin` 的中间件中为每个请求生成唯一的 `request_id`。
    *   通过 `context.Context` 将 `request_id` 和用户 `user_id` 等信息传递到程序的每一层，并在打印日志时始终带上这些字段，方便追踪一个请求的完整调用链。

---

## 2. 配置 (Configuration)

配置管理的目标是让应用在不同环境（开发、测试、生产）中能够灵活部署。

*   **推荐库**: `Viper`。
*   **配置文件**:
    *   在 `goapi` 项目根目录创建 `config/` 目录。
    *   在 `config/` 目录中创建 `config.yaml` 文件作为主配置文件。
    *   **示例 `config.yaml`**:
        ```yaml
        server:
          port: 8080
          mode: "debug" # gin mode: debug, test, release

        database:
          host: "localhost"
          port: 3306
          user: "root"
          password: "your_password"
          dbname: "sms_platform"

        redis:
          host: "localhost"
          port: 6379
          password: ""
        ```
*   **环境变量覆盖**:
    *   `Viper` 必须配置为支持用**环境变量**来覆盖配置文件中的值，这是生产环境部署的最佳实践（例如，通过 `DATABASE_PASSWORD` 环境变量注入数据库密码）。

---

## 3. 代码规范 (Code Specification)

所有代码都应遵循Go社区的最佳实践。

*   **格式化**:
    *   所有代码提交前 **必须** 使用 `goimports` 进行格式化。
*   **响应要求**:
    *  code 必须是有定义，不能直接写数字，方便作多语言 
    *  返回接口需要统一定义，不能直接写JSON 
*   **静态检查 (Linting)**:
    *   使用 `golangci-lint` 作为代码检查工具，并应在CI/CD流程中集成。
*   **命名规范**:
    *   **包名 (package)**: 小写、简短、有意义，不使用下划线或驼峰。例如 `service`, `repository`。
    *   **变量名**: 驼峰命名法 `camelCase`。
    *   **公开内容 (Exported)**: 所有需要对外暴露的函数、结构体、变量、接口，首字母**必须**大写（`PascalCase`）。
    *   **接口 (interface)**: 单个方法的接口名建议以 "er" 结尾，例如 `Reader`, `Writer`。
*   **错误处理**:
    *   **必须**显式处理每一个 `error` 返回值，不使用 `_` 忽略。
    *   向上传递错误时，必须使用 `fmt.Errorf("...: %w", err)` 的方式包装（wrap）原始错误，以保留完整的错误调用栈。
*   **注释**:
    *   注释应该解释代码的“**为什么** (why)”，而不是“**做什么** (what)”。而且是要中文注释
    *   所有公开的（首字母大写的）函数、类型、变量都**必须**有文档注释。
*   **项目目录结构**:
    ```
    .
    ├── cmd/
    │   └── server/
    │       └── main.go
    ├── config/
    │   └── config.yaml
    ├── internal/
    │   ├── api/
    │   ├── service/
    │   ├── repository/
    │   └── domain/
    ├── pkg/
    ├── go.mod
    └── DEVELOPMENT.md  (本文件)
    ```