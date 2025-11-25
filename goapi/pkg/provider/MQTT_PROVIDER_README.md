# MQTT Provider 使用说明

## 概述

MQTT Provider 是一个新的运营商实现，用于对接 `http://szb.jczl70.com:6086/mqtt/msg/` 接口。

## API 文档

### 1. 获取号码
- **URL**: `http://szb.jczl70.com:6086/mqtt/msg/getNumber?id=%d&key=%s`
- **方法**: GET
- **参数**:
  - `id`: 运营商ID（从数据库 `merchant_id` 字段读取）
  - `key`: 运营商KEY（从数据库 `merchant_key` 字段读取）
- **返回**:
```json
{
  "number": "1888888888",
  "extId": "2025111215435966402734",
  "id": 1
}
```

### 2. 获取验证码
- **URL**: `http://szb.jczl70.com:6086/mqtt/msg/getCode?extId=%s`
- **方法**: GET
- **参数**:
  - `extId`: 外部ID（从获取号码接口返回）
- **返回**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "receiveStatus": 1,
    "message": "456039"
  }
}
```
- **说明**:
  - `code`: 0 表示接口操作成功
  - `receiveStatus`: 0=短信接收失败, 1=短信接收成功
  - `message`: 短信信息（status为1时读取）
  - `error`: 错误信息（status为0时读取）

### 3. 释放号码
- **URL**: `http://szb.jczl70.com:6086/mqtt/msg/release?extId=%s&status=%d`
- **方法**: GET
- **参数**:
  - `extId`: 外部ID
  - `status`: 状态码
    - 1: 注册成功
    - 2: 超时
    - 3: 已注册
    - 4: 其它问题
- **说明**: 如果注册失败，必须要执行此接口

## 数据库配置

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
    'mqtt',  -- 或包含 'mqtt' 的 code
    'MQTT运营商',
    'http://szb.jczl70.com:6086',  -- API网关地址（不含路径）
    '1',  -- 运营商ID（id参数，请找管理员拿取）
    'your_key_here',  -- 运营商KEY（key参数，请找管理员拿取）
    1
);
```

## 自动识别

ProviderManager 会根据以下条件自动识别为 MQTT Provider：

1. **API网关地址**包含 `jczl70.com` 或 `mqtt/msg`
2. **Provider Code**包含 `mqtt`

满足任一条件即可。

## 业务类型配置

在 `sms_providers_business_types` 表中配置支持的业务类型：

```sql
INSERT INTO sms_providers_business_types (
    provider_id,
    business_code,
    business_name,
    price,
    status
) VALUES (
    (SELECT id FROM sms_providers WHERE code = 'mqtt'),
    'qq',
    'QQ',
    1.0,
    1
);
```

## 使用说明

1. **获取号码**: 系统会自动调用 `getNumber` 接口，返回的 `extId` 会保存到 `sms_phone_assignments.ext_id` 字段
2. **获取验证码**: 定时任务会自动调用 `getCode` 接口，使用保存的 `extId` 获取验证码
3. **释放号码**: 当号码过期或完成时，系统会自动调用 `release` 接口，使用保存的 `extId` 释放号码

## 注意事项

1. **extId 必须保存**: 系统会在获取号码时自动保存 `extId` 到数据库，后续操作都依赖此字段
2. **释放状态**: 目前默认使用状态 4（其它问题），如需自定义状态，可以修改 `ReleasePhoneWithExtId` 方法
3. **超时设置**: 默认超时时间为 30 秒，可以在创建 Provider 时通过 `Timeout` 配置项调整
4. **轮询间隔**: 获取验证码时，每 2 秒轮询一次，直到超时或获取成功

## 错误处理

- 如果获取号码时返回错误，会记录日志并继续尝试下一个号码
- 如果获取验证码时超时，会返回 `ErrCodeTimeout` 错误
- 如果释放号码时缺少 `extId`，会返回 `INVALID_PHONE` 错误

## 日志

所有操作都会记录详细日志，日志前缀为 `[MQTTProvider]`，方便调试和排查问题。

