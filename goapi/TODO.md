# SMS Platform GoAPI - Development TODO

> 📅 **创建时间**: 2024年11月21日  
> 🎯 **目标**: 完整实现 SMS 平台的 Go API 服务

## 📋 开发计划总览

### 🔄 **项目状态**
- ✅ 基础框架已搭建
- ✅ 用户认证系统已完成  
- ✅ 业务类型查询已实现
- ✅ Domain Models 已完善
- ✅ 测试套件修复完成 (所有指针类型错误已修复)
- ✅ Mock Provider 实现完成 (包括演示程序)
- ✅ **第一阶段：Mock服务商实现已完成**
- ✅ **架构重构：接口从service包移至pkg/provider包** (2024年11月24日)
- ❌ 核心 SMS 功能未开始
- ❌ IP白名单功能未开始
- ❌ 历史记录查询未开始

### 📊 **进度统计**
- **总任务**: 16 项
- **已完成**: 7 项 (基础框架 + 认证 + Domain + Tests + Mock Provider + Provider管理 + 架构重构)
- **进行中**: 0 项
- **完成率**: 43.8%
- **待开始**: 9 项

---

## 🚀 第一阶段：Mock服务商实现 ✅ **已完成**

### 1. 实现 Mock 服务商
**优先级**: 🔴 最高  
**预估工期**: 2-3天  
**状态**: ✅ **已完成** (2024年11月24日)

#### 📝 详细任务
- ✅ **定义服务商统一接口**
  - 文件: `internal/service/providers/provider.go`
  - 定义 `SMSProvider` 接口
  - 方法: `GetPhone(businessType, cardType)`, `GetCode(phone, timeout)`
  - 定义统一的响应结构和错误码

- ✅ **Mock 服务商实现**
  - 文件: `internal/service/providers/mock_provider.go`
  - 模拟真实服务商的行为
  - 随机生成手机号码 (支持多个国家)
  - 随机生成验证码 (4-6位数字)
  - 可配置成功率和延迟时间
  - 支持不同业务类型的模拟

- ✅ **服务商配置管理**
  - 文件: `internal/service/third_party_service.go`
  - 配置不同服务商的参数
  - 支持权重配置用于路由
  - 支持开关配置用于降级

- ✅ **Mock 数据生成器**
  - 生成符合各国格式的手机号
  - 模拟不同网络延迟
  - 模拟各种错误场景 (超时、无号码等)

#### 🔧 技术要求
- ✅ 使用接口模式，方便后续扩展真实服务商
- ✅ Mock 数据要足够真实，便于测试
- ✅ 支持配置文件动态调整参数
- ✅ 完整的错误模拟机制

#### ✅ 验收标准
- ✅ Mock 服务商接口实现完成
- ✅ 能够稳定返回手机号和验证码
- ✅ 支持多种业务类型模拟
- ✅ 错误场景覆盖完整
- ✅ 单元测试覆盖率 > 90%

---

### 2. 实现服务商管理服务
**优先级**: 🔴 最高
**预估工期**: 1-2天
**依赖**: 任务1
**状态**: ✅ **已完成** (2024年11月24日)

#### 📝 详细任务
- ✅ **服务商路由器实现**
  - 文件: `internal/service/third_party_service.go`
  - 根据业务类型选择服务商
  - 实现基于权重的负载均衡
  - 支持服务商降级和故障转移

- ✅ **服务商工厂模式**
  - 通过Provider接口统一管理
  - 根据配置创建不同的服务商实例
  - 支持运行时动态切换服务商

- ✅ **服务商健康检查**
  - 定期检查服务商可用性
  - 自动降级不可用的服务商
  - 记录服务商性能指标

#### ✅ 验收标准
- ✅ 权重路由算法正确实现
- ✅ 故障转移机制工作正常
- ✅ 健康检查稳定运行
- ✅ 支持配置热更新

### 🎯 **第一阶段成果总结**
- ✅ **Demo程序**: `go run cmd/mock_demo/main.go` 展示完整功能
- ✅ **测试覆盖**: 所有测试通过，107个测试用例全部成功
- ✅ **功能验证**: 支持qq、wechat、test业务类型，物理卡和虚拟卡
- ✅ **错误模拟**: 可配置失败率，延迟模拟，健康检查
- ✅ **负载均衡**: 基于优先级的Provider选择和故障转移

### 🏗️ **架构重构完成** (2024年11月24日)
- ✅ **接口重构**: 将SMSProvider接口从`internal/service`移至`pkg/provider`
- ✅ **清洁架构**: 遵循Go最佳实践，接口定义独立于业务逻辑
- ✅ **包结构优化**:
  ```
  pkg/provider/
  ├── interface.go    # SMSProvider接口、PhoneResponse、CodeResponse等类型定义
  └── README.md       # 接口使用说明
  
  internal/service/
  ├── third_party_service.go    # 业务逻辑实现
  └── providers/
      ├── mock_provider.go      # Mock实现
      └── http_provider.go      # HTTP实现
  ```
- ✅ **类型引用更新**: 所有文件已更新使用`provider.*`类型
- ✅ **测试验证**: 编译通过，所有测试正常，Demo程序功能完整

### 📈 **架构改进收益**
- 🎯 **符合Go规范**: 接口定义独立，便于其他包导入
- 🔧 **易于扩展**: 新增Provider实现时只需实现pkg/provider接口
- 🧪 **便于测试**: 接口与实现分离，Mock测试更简洁
- 📦 **减少依赖**: 业务逻辑与接口定义解耦

---

## 🎯 第二阶段：基础设施建设

### 3. 完善缺失的 Repository 实现
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

### 4. 实现 TransactionService
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
**依赖**: 任务3 (TransactionRepository)  
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

### 5. 实现 PhoneService
**优先级**: 🔴 最高  
**预估工期**: 2-3天  
**依赖**: 任务1, 任务2, 任务4  
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **手机号获取业务逻辑**
  - 文件: `internal/service/phone_service.go`  
  - 实现 `GetPhone(customerID, businessType, cardType)` 方法
  - 集成Mock服务商调用
  - 集成余额扣费逻辑
  - 实现权重路由算法

- [ ] **验证码获取业务逻辑**
  - 实现 `GetCode(customerID, phone, timeout)` 方法
  - 支持长轮询机制 (最长60秒)
  - 验证手机号归属权检查

- [ ] **分配记录管理**
  - 创建手机号分配记录
  - 更新验证码到记录
  - 处理过期和状态变更

#### 🔧 技术要求
- 业务逻辑与服务商实现解耦
- 完整的错误处理和回滚机制
- 长轮询使用合适的超时策略
- 记录详细的操作日志

#### ✅ 验收标准
- [ ] 手机号获取流程完整可用
- [ ] Mock服务商集成正常
- [ ] 权重路由算法正确执行
- [ ] 验证码获取支持长轮询
- [ ] 余额扣费集成正常
- [ ] 异常情况处理完善

---

## 📊 第三阶段：核心功能实现

### 6. 实现手机号获取接口
**优先级**: 🔴 最高  
**预估工期**: 1-2天  
**依赖**: 任务5  
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

### 7. 实现验证码获取接口
**优先级**: 🔴 最高
**预估工期**: 2天  
**依赖**: 任务5, 任务6
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

### 8. 实现余额查询接口
**优先级**: 🟡 中等
**预估工期**: 0.5天
**依赖**: 任务4
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

## 🔧 第四阶段：管理功能实现

### 9. 实现 AssignmentService
**优先级**: 🟡 中等
**预估工期**: 1-2天
**依赖**: 任务3
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

### 10. 实现手机号分配历史接口
**优先级**: 🟡 中等  
**预估工期**: 1天
**依赖**: 任务9
**状态**: ❌ 待开始

#### 📝 详细任务
- [ ] **API 实现**
  - 新增路由: `GET /client/v1/assignments`
  - 支持分页参数: `page`, `limit`
  - 支持筛选参数: `status`, `business_type`

---

### 11-13. 白名单和密码管理功能
**优先级**: 🟢 低等
**预估工期**: 3-4天
**状态**: ❌ 待开始

详细任务将在前面核心功能完成后补充...

---

## 🔧 第五阶段：系统完善

### 14-16. 路由、错误处理、测试文档
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
feat: 实现Mock服务商

- 添加 SMSProvider 接口定义
- 实现 MockProvider 用于测试
- 添加权重路由算法
- 完善错误处理机制

Closes #1
```

### 🧪 测试要求
- **单元测试**: 覆盖率 > 80%
- **集成测试**: API 接口功能测试  
- **Mock测试**: Mock服务商功能测试
- **性能测试**: 关键接口压力测试

---

## 📝 开发注意事项

### Mock 服务商设计原则
1. **真实性**: Mock 数据要尽可能接近真实场景
2. **可配置**: 支持通过配置文件调整行为
3. **可扩展**: 接口设计要便于后续接入真实服务商
4. **可测试**: 提供足够的错误场景用于测试

### 权重路由算法
- 支持基于权重的随机选择
- 实现故障转移机制
- 记录服务商性能指标
- 支持动态调整权重

---

## ✅ **Phase 1 完成总结: Mock Provider Implementation**

**完成时间**: 2024年11月24日  
**状态**: ✅ 已完成

### 🎯 已实现功能
1. **Mock服务商核心功能**
   - ✅ 完整的MockProvider实现 (`internal/service/providers/mock_provider.go`)
   - ✅ Third Party Service管理 (`internal/service/third_party_service.go`)
   - ✅ Provider接口定义和实现
   - ✅ 手机号分配和验证码生成模拟

2. **功能特性**
   - ✅ 支持多种业务类型：qq, wechat, test
   - ✅ 支持物理卡和虚拟卡类型
   - ✅ 可配置成功率（默认95%）
   - ✅ 延迟模拟（100-500ms）
   - ✅ 健康检查机制
   - ✅ 错误处理和故障转移

3. **测试和演示**
   - ✅ 完整的单元测试覆盖
   - ✅ 演示程序 (`cmd/mock_demo/main.go`)
   - ✅ 错误场景测试

### 🏆 成果展示
```bash
# 运行演示程序
go run cmd/mock_demo/main.go

# 运行测试
go test ./internal/service/providers -v
```

### 📈 下一步计划
现在Mock Provider已完全可用，可以作为测试基础来实现其他核心功能：
- SMS发送和接收功能
- 用户余额管理
- 历史记录查询
- IP白名单功能

---

## 📞 联系方式

有问题请在项目中创建 Issue 或联系开发团队。

**最后更新**: 2024年11月24日