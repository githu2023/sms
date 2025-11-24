# Provider ExtraConfig 使用说明

## 概述

`ExtraConfig` 是一个 JSON 格式的字段，用于存储不同运营商的特殊配置。这样可以灵活地支持各种运营商的不同需求，而不需要为每个运营商添加专门的字段。

## 数据库字段

在 `sms_providers` 表中，`extra_config` 字段类型为 JSON，用于存储运营商的额外配置。

## 使用示例

### BigBus666 运营商配置

BigBus666 运营商需要在 `extra_config` 中存储 `projectName`：

```sql
INSERT INTO sms_providers (
    code, 
    name, 
    api_gateway, 
    merchant_id, 
    merchant_key, 
    extra_config,
    is_enabled
) VALUES (
    'bigbus666',
    'BigBus666运营商',
    'http://mhm1111z.bigbus666.top:2086/s/m',
    '11296564',
    'Z8UrC8H2cNgXe2Jw',
    '{"projectName": "hema"}',  -- JSON格式的额外配置
    1
);
```

### 其他运营商配置示例

不同的运营商可以在 `extra_config` 中存储不同的配置：

```sql
-- 示例1: 运营商A需要自定义超时时间
INSERT INTO sms_providers (..., extra_config, ...) VALUES (
    ...,
    '{"timeout": 60, "retryCount": 3}',
    ...
);

-- 示例2: 运营商B需要特殊的认证参数
INSERT INTO sms_providers (..., extra_config, ...) VALUES (
    ...,
    '{"apiVersion": "v2", "region": "us-east-1"}',
    ...
);

-- 示例3: 运营商C需要多个项目名称映射
INSERT INTO sms_providers (..., extra_config, ...) VALUES (
    ...,
    '{"projectName": "hema", "projectMapping": {"wx": "wechat", "qq": "tencent"}}',
    ...
);
```

## 代码中使用

### 在 ProviderManager 中读取配置

```go
// 从 ExtraConfig 中读取特殊配置
projectName := "hema" // 默认值
if p.ExtraConfig != nil {
    if pn, ok := (*p.ExtraConfig)["projectName"].(string); ok && pn != "" {
        projectName = pn
    }
}
```

### 在 Provider 实现中使用

```go
// BigBus666Provider 中使用
projectName := p.info.Metadata["project_name"]
if projectName == "" {
    projectName = businessType // 回退到使用业务类型
}
```

## 配置结构

`ExtraConfig` 是一个 `map[string]interface{}` 类型，可以存储任意 JSON 结构：

```json
{
  "projectName": "hema",
  "customParam1": "value1",
  "customParam2": 123,
  "customParam3": {
    "nested": "value"
  },
  "customParam4": ["array", "values"]
}
```

## 注意事项

1. **类型安全**: 从 `ExtraConfig` 读取值时需要进行类型断言
2. **默认值**: 建议为每个配置项设置合理的默认值
3. **向后兼容**: 如果 `ExtraConfig` 为空或不存在，应该使用默认值
4. **文档**: 每个运营商应该在自己的文档中说明需要哪些 `ExtraConfig` 字段

## 迁移说明

如果之前使用了 `api_config` 字段，可以将其迁移到 `extra_config`：

```sql
-- 迁移示例（根据实际情况调整）
UPDATE sms_providers 
SET extra_config = JSON_OBJECT('legacy_config', api_config)
WHERE api_config IS NOT NULL AND extra_config IS NULL;
```

