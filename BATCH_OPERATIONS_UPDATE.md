# 批量操作功能更新

**更新时间:** 2025年11月24日  
**版本:** v2.0 - 批量操作支持

## 🚀 新功能概览

本次更新为 `get_phone` 和 `get_code` 接口添加了批量操作支持，允许一次性获取多个手机号码和验证码，大大提高了效率。

## 📊 核心改进

### 1. 批量获取手机号 (get_phone)

#### 新特性
- **批量数量**: 支持 1-10 个手机号批量获取
- **成本统计**: 返回总成本和剩余余额
- **成功率统计**: 显示成功/失败数量
- **余额保护**: 余额不足时停止继续获取

#### 请求格式
```json
{
  "business_type": "verification",
  "card_type": "any", 
  "count": 5  // 新增：批量获取数量
}
```

#### 响应格式
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "phones": [
      {
        "phone_number": "+15551234567",
        "country_code": "US", 
        "cost": 0.10,
        "valid_until": "2025-11-24T15:30:00Z",
        "provider_id": "provider-1"
      }
    ],
    "total_cost": 0.50,
    "remaining_balance": 9.50,
    "success_count": 5,
    "failed_count": 0
  }
}
```

### 2. 批量获取验证码 (get_code)

#### 新特性
- **批量手机号**: 支持 1-10 个手机号同时获取验证码
- **并发处理**: 使用 goroutine 并发获取，提高效率
- **状态分类**: 成功/超时/失败状态统计
- **长轮询支持**: 每个号码独立等待验证码

#### 请求格式
```json
{
  "phone_numbers": [  // 改为数组格式
    "+15551234567",
    "+15551234568", 
    "+15551234569"
  ],
  "timeout": 60
}
```

#### 响应格式
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "codes": [
      {
        "phone_number": "+15551234567",
        "code": "123456",
        "message": "Your verification code is 123456",
        "received_at": "2025-11-24T15:30:00Z",
        "provider_id": "provider-1",
        "status": "success"
      }
    ],
    "success_count": 2,
    "timeout_count": 1,
    "failed_count": 0
  }
}
```

## 🔧 技术实现

### Go API 层改进

#### 1. DTO 结构更新
- **GetPhoneRequest**: 新增 `count` 字段
- **GetCodeRequest**: `phone_number` 改为 `phone_numbers` 数组
- **新增结构**: `PhoneInfo`, `CodeInfo` 用于批量返回

#### 2. Handler 层改进
- **批量循环处理**: 支持1-10个手机号/验证码批量获取
- **并发优化**: 验证码获取使用 goroutine 并发处理
- **错误处理**: 精细化的成功/失败统计

### Flutter 客户端改进

#### 1. API 方法更新
```dart
// 新的批量方法
Future<ApiResponse<Map<String, dynamic>>> assignPhone({
  required String businessType,
  required String cardType,
  int count = 1, // 支持批量
})

Future<ApiResponse<Map<String, dynamic>>> getVerificationCode({
  required List<String> phoneNumbers, // 支持批量
  int timeout = 60,
})
```

#### 2. 向后兼容
```dart
// 保持原有单个获取的便捷方法
Future<ApiResponse<String?>> assignPhoneSingle({...})
Future<ApiResponse<String?>> getVerificationCodeSingle({...})
```

## 📋 使用示例

### 批量获取手机号
```dart
// 获取5个手机号
final response = await apiClient.assignPhone(
  businessType: 'verification',
  cardType: 'any',
  count: 5,
);

if (response.success) {
  final phones = response.data!['phones'] as List;
  final totalCost = response.data!['total_cost'];
  final successCount = response.data!['success_count'];
  
  print('成功获取 $successCount 个手机号，总费用: $totalCost');
}
```

### 批量获取验证码
```dart
// 同时获取多个手机号的验证码
final response = await apiClient.getVerificationCode(
  phoneNumbers: ['+1555123456', '+1555123457', '+1555123458'],
  timeout: 120,
);

if (response.success) {
  final codes = response.data!['codes'] as List;
  final successCount = response.data!['success_count'];
  final timeoutCount = response.data!['timeout_count'];
  
  print('成功: $successCount, 超时: $timeoutCount');
}
```

### 单个获取（向后兼容）
```dart
// 原有的单个获取方式仍然有效
final phoneResponse = await apiClient.assignPhoneSingle(
  businessType: 'verification',
  cardType: 'any',
);

final codeResponse = await apiClient.getVerificationCodeSingle(
  phone: '+1555123456',
);
```

## ⚡ 性能优势

### 批量获取手机号
- **减少请求次数**: 10个号码从10次请求降为1次
- **减少网络开销**: 大幅降低网络IO
- **统一费用计算**: 一次性扣费，避免多次数据库操作

### 批量获取验证码  
- **并发处理**: 使用 goroutine 同时等待多个验证码
- **总体时间优化**: 不再是线性累加时间，而是并行等待
- **资源利用率**: 更好地利用服务器资源

## 🔄 兼容性保障

### API 兼容性
- **向后兼容**: 原有单个获取的调用方式完全保持不变
- **渐进升级**: 可以逐步迁移到批量方式，无需一次性改动
- **错误处理**: 保持原有的错误码和消息格式

### 数据库兼容性
- **无结构变更**: 数据库表结构无需修改
- **事务保护**: 批量操作仍然保持事务一致性
- **余额管理**: 批量扣费确保余额正确计算

## 📈 使用场景

### 1. 批量注册场景
```
用户场景：需要批量注册多个账号
优化前：逐个获取手机号，效率低下
优化后：一次获取多个手机号，大幅提升效率
```

### 2. 自动化测试场景
```
用户场景：自动化测试需要大量手机号和验证码
优化前：串行获取，测试时间长
优化后：并行获取，测试时间大幅缩短
```

### 3. 业务批处理场景
```
用户场景：定期批量处理业务数据
优化前：逐个处理，资源利用率低
优化后：批量处理，资源利用率高
```

## 🎯 性能指标

| 指标 | 单个获取 | 批量获取 (10个) | 改进 |
|------|---------|---------------|------|
| 请求次数 | 10次 | 1次 | **90% 减少** |
| 网络延迟 | 累积延迟 | 单次延迟 | **大幅减少** |
| 服务器负载 | 10倍处理 | 1倍处理 | **90% 减少** |
| 用户等待时间 | 线性增长 | 并发等待 | **显著缩短** |

## ✅ 测试覆盖

### 单元测试
- ✅ 批量获取手机号功能测试
- ✅ 批量获取验证码功能测试  
- ✅ 边界条件测试（1个、10个、11个）
- ✅ 错误处理测试（余额不足、超时等）

### 集成测试
- ✅ Go API 批量接口测试
- ✅ Flutter 客户端批量方法测试
- ✅ 向后兼容性测试
- ✅ 并发安全性测试

## 🔮 后续规划

### 短期优化
1. **缓存优化**: 添加手机号池预缓存机制
2. **限流控制**: 添加批量操作的频率限制
3. **监控指标**: 添加批量操作的性能监控

### 长期规划
1. **智能批量**: 根据历史使用模式自动推荐批量大小
2. **预测性获取**: 基于用户行为预先获取手机号
3. **全局优化**: 跨用户的资源池化和优化

---

**总结**: 本次批量操作更新显著提升了系统的性能和用户体验，在保持完全向后兼容的同时，为高频使用场景提供了强大的批量处理能力。