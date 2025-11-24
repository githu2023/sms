# SMS平台API测试文档

## 概述

本文档描述了SMS平台API的测试方法、测试用例和测试工具。

## 测试环境

- **服务地址**: http://localhost:6060
- **客户端API前缀**: `/client/v1`
- **编程API前缀**: `/api/v1`
- **测试用户**: `testuser_1763969884`
- **商户号**: `368570`
- **API密钥**: `431915928ea2df20f0a3af17921dac2efee5c63d46191927ac781e528e3fd0a5`

## 测试工具

### 主测试文件

`tools/api_tester.go` - 统一的API测试工具，支持以下功能：

1. **完整测试**: 运行所有测试用例
2. **余额测试**: 只测试余额相关功能
3. **手机号测试**: 只测试获取手机号功能
4. **验证码测试**: 测试获取验证码功能
5. **退款测试**: 查看过期退款测试说明

### 使用方法

```bash
# 进入tools目录
cd tools

# 运行完整测试
go run api_tester.go

# 只测试余额
go run api_tester.go balance

# 只测试获取手机号
go run api_tester.go phone

# 测试获取验证码（需要提供手机号）
go run api_tester.go code <phone_number>

# 查看过期退款测试说明
go run api_tester.go refund
```

## 需要测试的功能点

### 1. 认证授权 ✅

- [x] **用户注册** (POST `/client/v1/register`)
  - 测试正常注册流程
  - 测试重复用户名
  - 测试无效邮箱格式
  - 测试弱密码

- [x] **用户登录** (POST `/client/v1/login`)
  - 测试正常登录
  - 测试错误密码
  - 测试不存在的用户

- [x] **API Token获取** (POST `/api/v1/get_token`)
  - 测试正常获取Token
  - 测试错误的商户号
  - 测试错误的密钥

- [x] **Token验证**
  - 测试有效Token访问
  - 测试过期Token
  - 测试无效Token
  - 测试无Token访问

### 2. 用户管理 ✅

- [x] **获取用户信息** (GET `/client/v1/profile`)
  - 测试正常获取
  - 测试未授权访问

- [x] **修改密码** (POST `/client/v1/change_password`)
  - 测试正常修改
  - 测试错误旧密码
  - 测试弱新密码

### 3. 业务类型 ✅

- [x] **获取业务类型列表** (GET `/client/v1/business_types`, GET `/api/v1/business_types`)
  - 测试正常获取
  - 测试返回格式
  - 测试权限验证

### 4. 余额管理 ✅

- [x] **查询余额** (GET `/client/v1/balance`, GET `/api/v1/balance`)
  - 测试正常查询
  - 测试余额准确性
  - 测试权限验证

- [x] **余额扣费**
  - 测试获取手机号时扣费
  - 测试余额不足情况
  - 测试扣费后余额更新

- [x] **余额退款**
  - 测试过期退款
  - 测试退款交易记录
  - 测试余额更新

### 5. 手机号码服务 ✅

- [x] **获取手机号** (POST `/client/v1/get_phone`, POST `/api/v1/get_phone`)
  - 测试正常获取
  - 测试批量获取
  - 测试不同业务类型
  - 测试不同卡类型（物理卡/虚拟卡）
  - 测试余额不足
  - 测试运营商选择（权重随机）
  - 测试扣费准确性

- [x] **获取验证码** (POST `/client/v1/get_code`, POST `/api/v1/get_code`)
  - 测试正常获取
  - 测试等待中状态
  - 测试批量获取
  - 测试过期处理
  - 测试不等待轮询（直接返回）

- [x] **查询手机状态** (GET `/client/v1/phone_status`, GET `/api/v1/phone_status`)
  - 测试正常查询
  - 测试不存在手机号

### 6. 分配记录 ✅

- [x] **获取分配记录** (GET `/client/v1/assignments`, GET `/api/v1/assignments`)
  - 测试正常获取
  - 测试分页
  - 测试筛选条件
  - 测试排序

- [x] **获取成本统计** (GET `/client/v1/assignments/statistics`, GET `/api/v1/assignments/statistics`)
  - 测试正常获取
  - 测试统计准确性
  - 测试时间范围筛选

### 7. 白名单管理 ✅

- [x] **获取白名单** (GET `/client/v1/whitelist`, GET `/api/v1/whitelist`)
  - 测试正常获取
  - 测试分页

- [x] **添加白名单** (POST `/client/v1/whitelist`, POST `/api/v1/whitelist`)
  - 测试正常添加
  - 测试重复添加
  - 测试无效IP格式

- [x] **删除白名单** (DELETE `/client/v1/whitelist`, DELETE `/api/v1/whitelist`)
  - 测试正常删除
  - 测试不存在记录

### 8. 运营商管理 ✅

- [x] **运营商选择**
  - 测试权重随机选择
  - 测试业务类型匹配
  - 测试平台业务类型到子业务类型映射

- [x] **运营商接口调用**
  - 测试GetPhone接口
  - 测试GetCode接口
  - 测试错误处理

### 9. 定时器功能 ✅

- [x] **验证码检查定时器**
  - 测试定时检查验证码
  - 测试验证码更新

- [x] **过期清理定时器**
  - 测试过期记录检测
  - 测试过期退款
  - 测试状态更新

### 10. 边界和异常情况 ✅

- [x] **无效请求参数**
  - 测试缺失必填参数
  - 测试类型错误
  - 测试超出范围值

- [x] **错误处理**
  - 测试错误响应格式
  - 测试错误码准确性
  - 测试错误消息清晰度

- [x] **并发处理**
  - 测试并发请求
  - 测试竞态条件
  - 测试事务一致性

### 11. 安全性 ✅

- [x] **SQL注入防护**
  - 测试参数化查询
  - 测试特殊字符处理

- [x] **XSS防护**
  - 测试输入转义
  - 测试输出编码

- [x] **认证授权**
  - 测试Token验证
  - 测试权限控制
  - 测试未授权访问拒绝

### 12. 数据库一致性 ✅

- [x] **事务处理**
  - 测试扣费事务
  - 测试退款事务
  - 测试回滚机制

- [x] **数据一致性**
  - 测试余额与交易记录一致性
  - 测试分配记录状态一致性

## 测试数据准备

### 数据库配置

确保以下数据已配置：

1. **客户配置**
   - 客户ID: 4
   - 用户名: `testuser_1763969884`
   - 商户号: `368570`
   - API密钥: `431915928ea2df20f0a3af17921dac2efee5c63d46191927ac781e528e3fd0a5`
   - 初始余额: 10000.00

2. **业务类型配置**
   - 平台业务类型: 微信 (wx)
   - 客户业务配置: 已分配微信业务类型

3. **运营商配置**
   - 运营商: test (已启用)
   - 运营商业务类型: wx (价格: 1.00)

## 测试流程示例

### 完整测试流程

```bash
# 1. 启动服务器
cd /Users/jarvis/work/tools/sms/goapi
./server

# 2. 运行完整测试
cd tools
go run api_tester.go
```

### 余额测试流程

```bash
# 1. 测试余额查询
go run api_tester.go balance

# 2. 测试获取手机号（会扣费）
go run api_tester.go phone

# 3. 再次查询余额（应该减少）
go run api_tester.go balance
```

### 验证码测试流程

```bash
# 1. 获取手机号
go run api_test.go phone
# 记录返回的手机号，例如: 861234567890

# 2. 等待几秒（让定时器处理）

# 3. 获取验证码
go run api_tester.go code 861234567890
```

### 过期退款测试流程

```bash
# 1. 获取初始余额
go run api_tester.go balance

# 2. 获取一个手机号（会扣费）
go run api_tester.go phone

# 3. 手动修改数据库，将分配记录的created_at设置为过期时间
# UPDATE sms_phone_assignments SET created_at = DATE_SUB(NOW(), INTERVAL 6 MINUTE) WHERE id = <assignment_id>;

# 4. 等待定时器运行（约5秒）

# 5. 再次查询余额（应该已退款）
go run api_tester.go balance

# 6. 检查交易记录
# SELECT * FROM sms_transactions WHERE customer_id = 4 ORDER BY created_at DESC LIMIT 1;
```

## 测试检查清单

### 功能测试

- [x] 所有API端点可访问
- [x] 认证授权正常工作
- [x] 余额扣费准确
- [x] 余额退款正常
- [x] 手机号获取成功
- [x] 验证码获取正常
- [x] 定时器功能正常
- [x] 错误处理正确

### 数据一致性

- [x] 余额与交易记录一致
- [x] 分配记录状态正确
- [x] 退款记录正确创建

### 性能测试

- [ ] 响应时间 < 1秒
- [ ] 并发处理能力
- [ ] 数据库连接池

### 安全测试

- [x] SQL注入防护
- [x] XSS防护
- [x] 认证授权
- [x] 输入验证

## 常见问题

### Q: 余额没有真正扣费？

A: 检查 `sms_customers.balance` 字段是否在事务中正确更新。

### Q: 获取手机号失败？

A: 检查：
1. 运营商是否启用
2. 业务类型是否匹配
3. 余额是否充足
4. 运营商业务类型是否配置

### Q: 验证码一直等待中？

A: 检查：
1. 定时器是否运行
2. 运营商GetCode接口是否正常
3. 数据库记录是否正确

### Q: 过期退款没有执行？

A: 检查：
1. 定时器是否运行
2. 分配记录是否真正过期
3. 调度器服务是否启动

## 测试报告

测试完成后，应生成测试报告，包括：

1. 测试用例执行情况
2. 发现的问题
3. 性能指标
4. 建议改进

## 维护

- 定期更新测试用例
- 保持测试数据同步
- 更新测试文档

