# SMS Platform GoAPI - Development TODO

> 📅 **创建时间**: 2024年11月21日  
> 🎯 **目标**: 完整实现 SMS 平台的 Go API 服务

## 📋 开发计划总览

### 🔄 **项目状态**
- ✅ 基础框架已搭建
- ✅ 用户认证系统已完成  
- ✅ 业务类型查询已实现
- 🔄 核心 SMS 功能开发中
- ❌ IP白名单功能未开始
- ❌ 历史记录查询未开始

### 📊 **进度统计**
- **总任务**: 15 项
- **已完成**: 0 项  
- **进行中**: 0 项
- **待开始**: 15 项
- **完成度**: 0%

---

## 🚀 第一阶段：基础设施建设

### 1. 完善缺失的 Repository 实现
**优先级**: 🔴 最高  
**预估工期**: 1-2天  
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **WhitelistRepository 实现**
  - 文件: `internal/repository/whitelist_repo.go`
  - 实现方法: `Create`, `FindByCustomerID`, `FindByID`, `Delete`
  - 包含 IP 格式验证逻辑
  - 支持 CIDR 格式的 IP 段

- [ ] **TransactionRepository 实现** 
  - 文件: `internal/repository/transaction_repo.go`
  - 实现方法: `Create`, `FindByCustomerID`, `GetBalanceByCustomerID`, `UpdateBalance`
  - 支持事务操作，确保余额一致性
  - 分页查询交易历史

- [ ] **LogRepository 实现**
  - 文件: `internal/repository/log_repo.go`  
  - 实现方法: `CreateRequestLog`, `CreateThirdPartyLog`, `FindByCustomerID`
  - 日志分类：API请求日志、第三方调用日志
  - 支持按时间范围查询

#### 🔧 技术要求
- 所有 Repository 必须实现对应的接口
- 使用 GORM 作为 ORM
- 添加必要的数据库索引
- 包含完整的错误处理

#### ✅ 验收标准
- [ ] 所有 Repository 接口方法实现完成
- [ ] 单元测试覆盖率 > 80%
- [ ] 数据库迁移脚本已创建
- [ ] 代码通过 golangci-lint 检查

---

### 2. 实现 TransactionService
**优先级**: 🔴 最高  
**预估工期**: 2-3天  
**依赖**: 任务1 (TransactionRepository)  
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **余额管理核心逻辑**
  - 文件: `internal/service/transaction_service.go`
  - 实现充值功能 `TopUp(customerID, amount, notes)`
  - 实现扣费功能 `Deduct(customerID, amount, referenceID, notes)`
  - 实现余额查询 `GetBalance(customerID)`

- [ ] **交易安全机制**
  - 使用数据库事务确保操作原子性
  - 实现余额不足检查
  - 添加交易锁机制防止并发问题
  - 记录详细的交易日志

- [ ] **交易记录管理**
  - 实现交易历史查询 `GetTransactionHistory(customerID, page, limit)`
  - 支持按交易类型筛选
  - 支持按时间范围查询

#### 🔧 技术要求
- 使用数据库事务保证数据一致性
- 实现乐观锁或悲观锁机制
- 所有金额计算使用 decimal 类型
- 完整的错误定义和处理

#### ✅ 验收标准
- [ ] 余额扣费和充值功能正常
- [ ] 并发测试无数据不一致问题  
- [ ] 交易历史查询分页正常
- [ ] 单元测试覆盖关键业务逻辑

---

### 3. 实现 ThirdPartyService  
**优先级**: 🔴 最高  
**预估工期**: 3-4天  
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **第三方服务商抽象接口**
  - 文件: `internal/service/third_party_service.go`
  - 定义统一的第三方服务接口 `SMSProvider`
  - 接口方法: `GetPhone(businessType, cardType)`, `GetCode(phone, timeout)`

- [ ] **服务商配置管理**
  - 支持多个服务商配置
  - 实现服务商路由和负载均衡
  - 支持服务商降级和切换

- [ ] **Mock 服务商实现**  
  - 文件: `internal/service/providers/mock_provider.go`
  - 用于开发和测试的模拟服务商
  - 随机生成手机号和验证码
  - 可配置成功率和延迟

- [ ] **HTTP 客户端封装**
  - 统一的 HTTP 请求封装
  - 超时控制和重试机制  
  - 请求和响应日志记录
  - 错误码标准化

#### 🔧 技术要求
- 使用接口模式，方便扩展不同服务商
- HTTP 客户端使用连接池
- 完整的超时和重试策略
- 详细的调用日志记录

#### ✅ 验收标准
- [ ] Mock 服务商功能完整可用
- [ ] 服务商接口设计合理易扩展
- [ ] HTTP 调用有完善的错误处理
- [ ] 集成测试通过

---

## 🎯 第二阶段：核心功能实现

### 4. 实现 PhoneService
**优先级**: 🔴 最高  
**预估工期**: 2-3天  
**依赖**: 任务2, 任务3  
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **手机号获取业务逻辑**
  - 文件: `internal/service/phone_service.go`  
  - 实现 `GetPhone(customerID, businessType, cardType)` 方法
  - 集成第三方服务调用
  - 集成余额扣费逻辑

- [ ] **验证码获取业务逻辑**
  - 实现 `GetCode(customerID, phone, timeout)` 方法
  - 支持长轮询机制 (最长60秒)
  - 验证手机号归属权检查

- [ ] **分配记录管理**
  - 创建手机号分配记录
  - 更新验证码到记录
  - 处理过期和状态变更

#### 🔧 技术要求
- 业务逻辑与第三方服务解耦
- 完整的错误处理和回滚机制
- 长轮询使用合适的超时策略
- 记录详细的操作日志

#### ✅ 验收标准
- [ ] 手机号获取流程完整可用
- [ ] 验证码获取支持长轮询
- [ ] 余额扣费集成正常
- [ ] 异常情况处理完善

---

### 5. 实现手机号获取接口
**优先级**: 🔴 最高  
**预估工期**: 1-2天  
**依赖**: 任务4  
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **API 路由定义**
  - 新增路由: `POST /api/v1/get_phone`
  - 新增路由: `POST /client/v1/get_phone`

- [ ] **请求处理器实现**
  - 文件: `internal/api/handler/phone_handler.go`
  - 实现 `GetPhone` 方法
  - 请求参数验证: `business_type`, `card_type`
  - 响应格式标准化

- [ ] **DTO 定义**  
  - 文件: `internal/api/handler/phone_dto.go`
  - 定义请求结构: `GetPhoneRequest`
  - 定义响应结构: `GetPhoneResponse`

#### 🔧 技术要求  
- 严格的参数验证
- 统一的错误响应格式
- 支持两种认证方式 (JWT/API Token)
- 完整的 API 文档注释

#### ✅ 验收标准
- [ ] API 接口调用成功返回手机号
- [ ] 参数验证完整有效
- [ ] 错误码定义清晰
- [ ] API 文档更新

---

### 6. 实现验证码获取接口
**优先级**: 🔴 最高
**预估工期**: 2天  
**依赖**: 任务4, 任务5
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **API 路由定义**
  - 新增路由: `POST /api/v1/get_code`
  - 新增路由: `POST /client/v1/get_code`

- [ ] **长轮询实现**
  - 支持最长60秒的轮询等待
  - 实现优雅的超时处理  
  - 避免资源泄露

- [ ] **请求处理器实现**
  - 在 `phone_handler.go` 中添加 `GetCode` 方法
  - 验证手机号归属权
  - 返回验证码或超时信息

#### ✅ 验收标准
- [ ] 长轮询机制工作正常
- [ ] 超时处理优雅
- [ ] 手机号归属权验证有效
- [ ] API 响应时间合理

---

### 7. 实现余额查询接口
**优先级**: 🟡 中等
**预估工期**: 0.5天
**依赖**: 任务2
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **API 路由定义**
  - 修改现有路由: `GET /api/v1/balance`
  - 修改现有路由: `GET /client/v1/balance`

- [ ] **处理器实现**
  - 文件: `internal/api/handler/balance_handler.go`
  - 调用 TransactionService 获取余额
  - 返回标准格式响应

#### ✅ 验收标准
- [ ] 余额查询返回正确数值
- [ ] API 响应格式统一
- [ ] 权限验证正常

---

## 📊 第三阶段：管理功能实现

### 8. 实现 AssignmentService
**优先级**: 🟡 中等
**预估工期**: 1-2天
**依赖**: 任务1
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **历史记录管理**
  - 文件: `internal/service/assignment_service.go`
  - 实现分页查询用户历史记录
  - 支持按状态筛选
  - 支持按时间范围查询

- [ ] **状态更新管理**
  - 实现记录状态更新
  - 处理超时过期逻辑
  - 成本统计功能

---

### 9. 实现手机号分配历史接口
**优先级**: 🟡 中等  
**预估工期**: 1天
**依赖**: 任务8
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **API 实现**
  - 新增路由: `GET /client/v1/assignments`
  - 支持分页参数: `page`, `limit`
  - 支持筛选参数: `status`, `business_type`

---

### 10-12. 白名单和密码管理功能
**优先级**: 🟢 低等
**预估工期**: 3-4天
**状态**: ❌ 待开始

详细任务将在前面核心功能完成后补充...

---

## 🔧 第四阶段：系统完善

### 13-15. 路由、错误处理、测试文档
**优先级**: 🟢 低等  
**预估工期**: 2-3天
**状态**: ❌ 待开始

---

## 📖 开发规范

### 🔍 代码检查清单
- [ ] 所有代码使用 `goimports` 格式化
- [ ] 通过 `golangci-lint` 静态检查
- [ ] 错误码统一定义在常量中
- [ ] API 响应格式使用统一结构
- [ ] 添加结构化日志记录
- [ ] 包含必要的单元测试

### 📋 提交规范  
```
feat: 实现手机号获取接口

- 添加 POST /api/v1/get_phone 接口
- 集成第三方服务调用
- 添加余额扣费逻辑
- 完善错误处理机制

Closes #5
```

### 🧪 测试要求
- **单元测试**: 覆盖率 > 80%
- **集成测试**: API 接口功能测试  
- **性能测试**: 关键接口压力测试
- **Mock 测试**: 第三方服务调用测试

---

## 📞 联系方式

有问题请在项目中创建 Issue 或联系开发团队。

**最后更新**: 2024年11月21日