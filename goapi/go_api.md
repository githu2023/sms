# 短信平台 - Go API 文档

**版本:** 1.0
**基础URL:** `/` (Go服务的根路径)

---

## 1. 通用原则

### 1.1. 响应格式
所有API响应均为JSON格式，并遵循以下统一结构。

**成功响应:**
```json
{
  "code": 200,
  "msg": "Success",
  "data": {
    // 具体接口的响应数据
  }
}
```

**错误响应:**
```json
{
  "code": 60001, 
  "msg": "具体的错误信息",
  "data": null
}
```

### 1.2. 认证方式
本API针对两类客户端使用两种不同的认证方案。

- **程序化API (`/api/v1`)**: 使用 **Bearer Token**。Token通过永久的密钥(secret key)从`/api/v1/get_token`接口获取。后续所有请求都必须在HTTP头中包含: `Authorization: Bearer <token>`。

- **客户端API (`/client/v1`)**: 使用 **JSON Web Token (JWT)**。JWT通过用户名和密码从`/client/v1/login`接口获取。后续所有请求都必须在HTTP头中包含: `Authorization: Bearer <jwt>`。

---

## 2. 程序化API (`/api/v1`)
这组接口为机器间通信而设计。

### 2.1. `POST /api/v1/get_token`
**描述:** 使用永久的客户密钥换取一个短时效的API访问令牌。

**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `secret` | string | 是 | 客户的唯一永久API密钥。 |

**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `token` | string | 生成的访问令牌。 |
| `expires_in`| integer | 令牌的有效时长（秒）。 |

### 2.2. `POST /api/v1/get_phone`
**描述:** 为指定的业务类型批量获取手机号码（支持1-10个）。
**请求头:** `Authorization: Bearer <token>`

**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `business_type` | string | 是 | 业务的唯一代码 (例如: "qq", "wechat")。 |
| `card_type` | string | 是 | 请求的SIM卡类型 (例如: "physical", "virtual", "any")。 |
| `count` | integer | 否 | 批量获取数量，默认为1，最大为10。 |

**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `phones` | array | 手机号信息对象的数组。 |
| `phones[].phone_number` | string | 获取到的手机号码。 |
| `phones[].country_code` | string | 国家代码 (例如: "US", "CN")。 |
| `phones[].cost` | number | 该手机号的费用。 |
| `phones[].valid_until` | string | 有效期至 (ISO 8601格式时间戳)。 |
| `phones[].provider_id` | string | 服务商ID。 |
| `total_cost` | number | 本次获取的总费用。 |
| `remaining_balance` | number | 剩余账户余额。 |
| `success_count` | integer | 成功获取的数量。 |
| `failed_count` | integer | 失败的数量。 |

**示例请求:**
```json
{
  "business_type": "wechat",
  "card_type": "virtual",
  "count": 3
}
```

**示例响应:**
```json
{
  "code": 200,
  "msg": "Success",
  "data": {
    "phones": [
      {
        "phone_number": "+8612345678901",
        "country_code": "CN",
        "cost": 0.10,
        "valid_until": "2024-01-01T12:30:00Z",
        "provider_id": "test"
      },
      {
        "phone_number": "+8612345678902",
        "country_code": "CN",
        "cost": 0.10,
        "valid_until": "2024-01-01T12:30:00Z",
        "provider_id": "test"
      }
    ],
    "total_cost": 0.20,
    "remaining_balance": 99.80,
    "success_count": 2,
    "failed_count": 0
  }
}
```

### 2.3. `POST /api/v1/get_code`
**描述:** 批量获取指定手机号收到的验证码。如果验证码还未获取到，返回等待状态，客户端需要再次请求。支持批量查询（最多10个手机号）。
**请求头:** `Authorization: Bearer <token>`

**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `phone_numbers` | array[string] | 是 | 需要获取验证码的手机号列表，最多10个。 |

**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `codes` | array | 验证码信息对象的数组。 |
| `codes[].phone_number` | string | 手机号码。 |
| `codes[].code` | string | 收到的验证码（如果还未获取到则为空字符串）。 |
| `codes[].message` | string | 消息内容。 |
| `codes[].received_at` | string | 验证码接收时间 (ISO 8601格式时间戳)。 |
| `codes[].provider_id` | string | 服务商ID。 |
| `codes[].status` | string | 状态：`success`（已获取）、`pending`（等待中）、`failed`（失败）。 |
| `success_count` | integer | 成功获取验证码的数量。 |
| `pending_count` | integer | 等待中的数量（验证码还未获取到）。 |
| `failed_count` | integer | 失败的数量。 |

**示例请求:**
```json
{
  "phone_numbers": ["+8612345678901", "+8612345678902"]
}
```

**示例响应:**
```json
{
  "code": 200,
  "msg": "Success",
  "data": {
    "codes": [
      {
        "phone_number": "+8612345678901",
        "code": "123456",
        "message": "Your verification code is 123456",
        "received_at": "2024-01-01T12:30:00Z",
        "provider_id": "test",
        "status": "success"
      },
      {
        "phone_number": "+8612345678902",
        "code": "",
        "message": "Waiting for code",
        "received_at": "2024-01-01T12:30:00Z",
        "provider_id": "test",
        "status": "pending"
      }
    ],
    "success_count": 1,
    "pending_count": 1,
    "failed_count": 0
  }
}
```

### 2.4. `GET /api/v1/balance`
**描述:** 查询当前账户余额。
**请求头:** `Authorization: Bearer <token>`

**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `balance` | number | 当前账户余额，精确到小数点后4位。 |

### 2.5. `GET /api/v1/business_types`
**描述:** 获取平台支持的所有业务类型列表。
**请求头:** `Authorization: Bearer <token>`

**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `business_types` | array | 业务类型对象的数组。 |
| `business_types[].name` | string | 业务的显示名称 (例如: "腾讯QQ")。 |
| `business_types[].code` | string | 业务的唯一代码 (例如: "qq")。 |

### 2.6. `GET /api/v1/phone_status`
**描述:** 查询指定手机号的当前状态和有效期。
**请求头:** `Authorization: Bearer <token>`

**查询参数:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `phone_number` | string | 是 | 需要查询的手机号。 |

**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `phone_number` | string | 手机号码。 |
| `status` | string | 当前状态 (例如: "active")。 |
| `valid_until` | string | 有效期至 (ISO 8601格式时间戳)。 |
| `customer_id` | integer | 客户ID。 |

### 2.7. IP白名单 (`/api/v1/whitelist`)
**描述:** 管理用于API密钥访问的IP白名单。
**请求头:** `Authorization: Bearer <token>`

- **`GET /api/v1/whitelist`**: 列出账户下所有的IP白名单。
  - **成功响应 (`data` 对象):** IP对象的数组。
    ```json
    [
      { "id": 1, "ip_address": "1.2.3.4", "notes": "办公室IP" }
    ]
    ```

- **`POST /api/v1/whitelist`**: 新增一个IP到白名单。
  - **请求体:**
    | 字段 | 类型 | 必需 | 描述 |
    | :--- | :--- | :--- | :--- |
    | `ip_address` | string | 是 | IP地址或CIDR格式的IP段 (例如: "1.2.3.0/24")。 |
    | `notes` | string | 否 | 备注信息。 |
  - **成功响应 (`data` 对象):** 新创建的IP对象。

- **`DELETE /api/v1/whitelist`**: 删除一个IP白名单。
  - **请求体:**
    | 字段 | 类型 | 必需 | 描述 |
    | :--- | :--- | :--- | :--- |
    | `ip_address` | string | 是 | 要删除的IP地址或CIDR格式的IP段。 |
  - **成功响应:** 无`data`对象，仅返回成功消息。

---

## 3. 客户端API (`/client/v1`)
这组接口为Flutter客户端应用而设计。

### 3.1. 认证

#### `POST /client/v1/register`
**描述:** 创建一个新的用户账户。
**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `username` | string | 是 | 唯一的用户名。 |
| `email` | string | 是 | 唯一的邮箱地址。 |
| `password` | string | 是 | 用户密码 (最少8位)。 |
**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `user_id` | integer | 新创建用户的ID。 |
| `username` | string | 新创建用户的用户名。 |

#### `POST /client/v1/login`
**描述:** 认证用户并返回用于会话管理的JWT。
**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `username` | string | 是 | 用户的用户名。 |
| `password` | string | 是 | 用户的密码。 |
**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `token` | string | 生成的JWT。 |
| `expires_in`| integer | JWT的有效时长（秒）。 |

#### `POST /client/v1/password/change`
**描述:** 允许已登录用户修改自己的密码。
**请求头:** `Authorization: Bearer <jwt>`
**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `old_password` | string | 是 | 用户的当前密码。 |
| `new_password` | string | 是 | 用户的新密码。 |
**成功响应:** 无`data`对象，仅返回成功消息。

### 3.2. 核心与账户API
*(以下所有接口都需要 `Authorization: Bearer <jwt>`)*

#### `GET /client/v1/profile`
**描述:** 获取当前登录用户的详细信息。
**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `user_id` | integer | 用户的唯一ID。 |
| `username` | string | 用户名。 |
| `email` | string | 邮箱地址。 |
| `balance` | number | 当前账户余额。 |
| `api_secret_key`| string | 用于程序化API调用的永久密钥。 |
| `registration_ip`| string | 注册时使用的IP。 |
| `last_login_at`| string | 上次登录的ISO 8601格式时间戳。 |

#### `GET /client/v1/assignments`
**描述:** 分页获取用户的手机号分配历史记录。
**查询参数:**
| 字段 | 类型 | 默认值 | 描述 |
| :--- | :--- | :--- | :--- |
| `page` | integer | 1 | 需要获取的页码。 |
| `limit`| integer | 20 | 每页的项目数量。 |
**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `items` | array | 分配记录对象的数组。 |
| `items[].phone` | string | 分配的手机号。 |
| `items[].business_type`| string | 业务类型代码 (例如: "qq")。 |
| `items[].card_type` | string | SIM卡类型 (例如: "physical", "virtual")。 |
| `items[].code` | string | 收到的验证码 (可能为空)。 |
| `items[].cost` | number | 本次分配的费用。 |
| `items[].status` | string | 最终状态 (例如: "completed", "expired")。 |
| `items[].created_at` | string | 创建时间的ISO 8601格式时间戳。 |
| `pagination` | object | 分页元数据。 |
| `pagination.total` | integer | 记录总数。 |
| `pagination.page` | integer | 当前页码。 |
| `pagination.limit` | integer | 每页项目数。 |

#### `POST /client/v1/get_phone`
**描述:** 使用客户端JWT认证，为指定业务类型批量请求新手机号（支持1-10个）。
**请求头:** `Authorization: Bearer <jwt>`

**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `business_type` | string | 是 | 业务的唯一代码 (例如: "qq", "wechat")。 |
| `card_type` | string | 是 | 请求的SIM卡类型 (例如: "physical", "virtual", "any")。 |
| `count` | integer | 否 | 批量获取数量，默认为1，最大为10。 |

**成功响应:** 与 `POST /api/v1/get_phone` 相同。

#### `POST /client/v1/get_code`
**描述:** 批量获取指定手机号收到的验证码。支持批量查询（最多10个手机号）。
**请求头:** `Authorization: Bearer <jwt>`

**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `phone_numbers` | array[string] | 是 | 需要获取验证码的手机号列表，最多10个。 |

**成功响应:** 与 `POST /api/v1/get_code` 相同。

#### `GET /client/v1/phone_status`
**描述:** 查询指定手机号的当前状态和有效期。
**请求头:** `Authorization: Bearer <jwt>`

**查询参数:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `phone_number` | string | 是 | 需要查询的手机号。 |

**成功响应:** 与 `GET /api/v1/phone_status` 相同。

#### 其他接口
`/client/v1/`下的下列接口与其在`/api/v1/`中的对应接口功能相同，但使用JWT认证并代表当前登录用户进行操作：
- `GET /client/v1/balance`
- `GET /client/v1/business_types`
- `GET /client/v1/whitelist`
- `POST /client/v1/whitelist`
- `DELETE /client/v1/whitelist`（使用请求体传递`ip_address`）

---

## 4. 通用错误码
*(部分重要错误码示例)*

| Code | Message (中文) |
| :--- | :--- |
| 40001 | 请求参数错误 |
| 40101 | 未授权 (Token无效或过期) |
| 40301 | 禁止访问 (例如IP不在白名单内) |
| 40401 | 资源不存在 |
| 50001 | 服务器内部错误 |
| 60001 | 客户余额不足 |
| 60101 | 客户密钥错误 |
| 60501 | 验证码不存在或超时 |
