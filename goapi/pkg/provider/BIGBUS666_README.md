# BigBus666 Provider 使用说明

## 概述

BigBus666 Provider 是 SMS 平台的一个第三方运营商实现，支持通过 HTTP API 获取手机号和验证码。

## 配置说明

### 数据库配置

在 `sms_providers` 表中添加或更新记录：

```sql
INSERT INTO sms_providers (
    code, 
    name, 
    api_gateway, 
    merchant_id, 
    merchant_key, 
    is_enabled
) VALUES (
    'bigbus666',                    -- 运营商编码（必须包含 "bigbus"）
    'BigBus666运营商',              -- 运营商名称
    'http://mhm1111z.bigbus666.top:2086',  -- API网关地址
    'your_customer_out_number',     -- 客户外部数字（customerOutNumber）
    'your_encrypt_key',             -- AES加密密钥（16字节）
    1                                -- 是否启用
);
```

### 配置字段说明

- **code**: 运营商编码，必须包含 "bigbus" 才会使用 BigBus666Provider
- **api_gateway**: API网关地址，例如：`http://mhm1111z.bigbus666.top:2086`
- **merchant_id**: 客户外部数字（customerOutNumber），从运营商客服获取
- **merchant_key**: AES加密密钥，从运营商客服获取（必须是16字节）

### 业务类型配置

在 `sms_provider_business_types` 表中配置支持的业务类型：

```sql
INSERT INTO sms_provider_business_types (
    provider_id,
    business_code,
    business_name,
    price,
    status
) VALUES (
    (SELECT id FROM sms_providers WHERE code = 'bigbus666'),
    'wx',        -- 业务类型代码（对应 projectName）
    '微信',      -- 业务类型名称
    1.00,        -- 价格
    1            -- 启用状态
);
```

## API 接口说明

根据 [运营商文档](http://mhm1111z.bigbus666.top:2086/doc/doc_m_hema.html)，BigBus666 Provider 实现了以下接口：

### 1. 获取手机号 (GetPhone)

- **接口**: `/n/{customerOutNumber}`
- **请求参数**: `{"projectName": "业务类型"}`
- **返回**: `{extId, mobile}`

### 2. 获取验证码 (GetCode)

- **接口**: `/r/{customerOutNumber}`
- **请求参数**: `{"extId": "号码ID"}`
- **返回**: `{receiveStatus, message}`

### 3. 释放号码 (可选)

- **接口**: `/f/{customerOutNumber}`
- **请求参数**: `{extId, status, reason}`

## 加密方式

BigBus666 Provider 使用以下加密方式：

- **算法**: AES-128
- **模式**: ECB
- **填充**: PKCS7Padding
- **编码**: Base64

所有请求和响应都经过 AES 加密和 Base64 编码。

## 使用示例

Provider 会在系统启动时自动从数据库加载配置并注册。使用时无需额外代码，系统会自动根据业务类型映射选择合适的 Provider。

## 注意事项

1. **密钥长度**: AES 密钥必须是 16 字节，如果长度不足会自动填充，如果过长会自动截断
2. **超时设置**: 默认超时时间为 30 秒，可以在配置中调整
3. **验证码轮询**: GetCode 方法会每 2 秒轮询一次，直到获取到验证码或超时
4. **映射关系**: Provider 会维护 extId 和 phoneNumber 的映射关系，用于后续的验证码获取

## 错误处理

Provider 实现了完整的错误处理机制：

- `ErrProviderUnavailable`: 运营商不可用
- `NewProviderError("API_ERROR", message)`: API 返回错误
- `NewProviderError("INVALID_RESPONSE", message)`: 响应数据不完整
- `NewProviderError("INVALID_PHONE", message)`: 手机号无效
- `ErrCodeTimeout`: 验证码获取超时

## 日志记录

Provider 使用标准库的 `log` 包记录日志，所有关键操作都会记录日志，便于调试和监控。

