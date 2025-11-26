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
        "phone_number": "15820192778",
        "country_code": "CN",
        "cost": 0.01,
        "valid_until": "2025-11-25T12:25:02+08:00",
        "provider_id": "mqtt"
      }
    ],
    "total_cost": 0.01,
    "remaining_balance": 9996.99,
    "success_count": 1,
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

### 2.4. `GET /api/v1/assignments`
**描述:** 分页获取程序化账户的手机号分配记录，过滤字段与客户端接口一致。
**请求头:** `Authorization: Bearer <token>`

**查询参数:** `page`、`limit`、`status`、`business_type`、`start_date`、`end_date`（含义、格式同 `/client/v1/assignments`）。

**成功响应 (`data` 对象):** 结构与 `/client/v1/assignments` 完全一致。

### 2.5. `GET /api/v1/assignments/recent`
**描述:** 获取最近的手机号获取记录（默认 5 条，可 1-50）。
**查询参数:** `limit`（默认 5）。
**成功响应:** `data.items` 列表与 `/client/v1/assignments/recent` 相同。

### 2.6. `GET /api/v1/balance`
**描述:** 查询当前账户余额。
**请求头:** `Authorization: Bearer <token>`

**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `balance` | number | 当前账户余额，单位元。 |
| `frozen_amount` | number | 当前冻结金额，正在等待结算或释放的额度。 |
| `currency` | string | 货币代码（目前固定为 `USD`）。 |

### 2.7. `GET /api/v1/business_types`
**描述:** 获取平台支持的所有业务类型列表。
**请求头:** `Authorization: Bearer <token>`

**成功响应 (`data` 数组):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `id` | integer | 业务配置ID。 |
| `business_code` | string | 业务的唯一代码，如 `wx`。 |
| `business_name` | string | 业务名称。 |
| `weight` | integer | 权重，用于多渠道路由。 |

### 2.8. `GET /api/v1/phone_status`
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
| `status` | string | 当前状态（示例：`active`）。 |
| `valid_until` | string | 有效期至 (ISO 8601格式时间戳)。 |
| `customer_id` | integer | 客户ID。 |

### 2.9. IP白名单 (`/api/v1/whitelist`)
**描述:** 管理用于API密钥访问的IP白名单。
**请求头:** `Authorization: Bearer <token>`

- **`GET /api/v1/whitelist`**: 列出账户下所有的IP白名单。
  - **成功响应:** `{"total": <int>,"list":[{ "id":1,"customer_id":4,"ip_address":"1.2.3.4","notes":"备注","created_at":1732525210 }] }`
  - *说明：该接口直接返回 JSON 对象，不包含统一的 `code/msg` 包装。*

- **`POST /api/v1/whitelist`**: 新增一个IP到白名单。
  - **请求体:**
    | 字段 | 类型 | 必需 | 描述 |
    | :--- | :--- | :--- | :--- |
    | `ip_address` | string | 是 | IP地址或CIDR格式的IP段 (例如: "1.2.3.0/24")。 |
    | `notes` | string | 否 | 备注信息。 |
  - **成功响应:** `{"message":"添加成功"}`。

- **`DELETE /api/v1/whitelist`**: 删除一个IP白名单。
  - **请求体:**
    | 字段 | 类型 | 必需 | 描述 |
    | :--- | :--- | :--- | :--- |
    | `ip_address` | string | 是 | 要删除的IP地址或CIDR格式的IP段。 |
  - **成功响应:** `{"message":"删除成功"}`。

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

#### `POST /client/v1/change_password`
**描述:** 允许已登录用户修改自己的密码。
**请求头:** `Authorization: Bearer <jwt>`
**请求体:**
| 字段 | 类型 | 必需 | 描述 |
| :--- | :--- | :--- | :--- |
| `old_password` | string | 是 | 用户的当前密码。 |
| `new_password` | string | 是 | 用户的新密码。 |
**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `message` | string | 始终返回“密码修改成功”。 |

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
| `frozen_amount` | number | 当前冻结金额。 |
| `api_secret_key`| string | 用于程序化API调用的永久密钥。 |
| `registration_ip`| string | 注册时使用的IP。 |
| `last_login_at`| string | 上次登录的ISO 8601格式时间戳。 |

#### `GET /client/v1/balance`
**描述:** 查询当前账户余额。
**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `balance` | number | 当前余额，单位元（示例：9996.99）。 |
| `frozen_amount` | number | 当前冻结金额，待结算或待释放的金额。 |
| `currency` | string | 货币代码，当前固定为 `USD`。 |

#### `GET /client/v1/business_types`
**描述:** 返回当前客户可用的业务类型列表。
**成功响应 (`data` 数组):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `id` | integer | 业务配置ID。 |
| `business_code` | string | 业务编码，如 `wx`、`qq`。 |
| `business_name` | string | 业务名称，如 “微信”。 |
| `weight` | integer | 权重，用于多渠道路由。 |

#### `GET /client/v1/assignments`
**描述:** 分页获取用户的手机号分配历史记录。
**查询参数:**
| 字段 | 类型 | 默认值 | 描述 |
| :--- | :--- | :--- | :--- |
| `page` | integer | 1 | 页码。 |
| `limit`| integer | 20 | 每页数量 (1-100)。 |
| `status` | integer | 0 | 1=pending、2=completed、3=expired、4=failed；0 表示不限。 |
| `business_type` | string | 空 | 按业务编码过滤。 |
| `start_date` | string | 空 | 起始日期 `YYYY-MM-DD`。 |
| `end_date` | string | 空 | 结束日期 `YYYY-MM-DD`（包含当天）。 |
**成功响应 (`data` 对象):**
| 字段 | 类型 | 描述 |
| :--- | :--- | :--- |
| `items` | array | 分配记录对象的数组。 |
| `items[].id` | integer | 分配记录ID。 |
| `items[].phone_number` | string | 分配的手机号。 |
| `items[].business_type`| string | 业务类型代码 (例如: "wx")。 |
| `items[].card_type` | string | 卡类型（当前固定返回 "virtual"）。 |
| `items[].verification_code` | string | 最新验证码（若仍在等待则为空字符串）。 |
| `items[].cost` | number | 本次分配的费用（`merchant_fee`）。 |
| `items[].status` | integer | 状态：1=pending、2=completed、3=expired、4=failed。 |
| `items[].created_at` | string | 创建时间 (ISO 8601)。 |
| `items[].provider_name` | string | 渠道名称（若可获取）。 |
| `pagination` | object | 分页元数据。 |
| `pagination.total` | integer | 记录总数。 |
| `pagination.page` | integer | 当前页码。 |
| `pagination.limit` | integer | 每页项目数。 |

> `/api/v1/assignments` 与该接口响应结构一致，只是认证方式不同。

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
- `GET /client/v1/assignments/recent`
- `GET /client/v1/whitelist`
- `POST /client/v1/whitelist`
- `DELETE /client/v1/whitelist`（使用请求体传递`ip_address`）

#### `GET /client/v1/assignments/recent`
**描述:** 获取最近的手机号获取记录（默认最新5条，可自定义1-50之间）。
**请求头:** `Authorization: Bearer <jwt>`

**查询参数:**
| 字段 | 类型 | 默认值 | 描述 |
| :--- | :--- | :--- | :--- |
| `limit` | integer | 5 | 返回的记录条数，范围 1-50。 |

**成功响应 (`data` 对象):**
```json
{
  "items": [
    {
      "id": 1,
      "phone_number": "15820192778",
      "business_type": "wx",
      "card_type": "physical",
      "verification_code": "",
      "cost": 0.01,
      "status": 1,
      "created_at": "2025-11-25T11:25:02+08:00",
      "provider_name": "MQTT"
    }
  ]
}
```
> 与 `/client/v1/assignments` 的单页结果一致，但不返回分页信息，适合刷新时快速显示最近记录。同名接口 `/api/v1/assignments/recent` 也可用于程序化查询。

#### `GET /client/v1/whitelist`
**描述:** 分页查询白名单。
**查询参数:** `page`（默认1）、`limit`（默认20）。
**成功响应:**
```json
{
  "total": 2,
  "list": [
    {
      "id": 10,
      "customer_id": 4,
      "ip_address": "1.2.3.4",
      "notes": "办公室IP",
      "created_at": 1732525210
    }
  ]
}
```
> *说明：白名单接口直接返回 JSON 对象，无统一的 `code/msg` 包装。*

#### `POST /client/v1/whitelist`
**请求体:** `{ "ip_address": "1.2.3.4", "notes": "描述" }`
**成功响应:** `{"message":"添加成功"}`。

#### `DELETE /client/v1/whitelist`
**请求体:** `{ "ip_address": "1.2.3.4" }`
**成功响应:** `{"message":"删除成功"}`。

### 3.3. 客户端典型流程（推荐）

开发客户端时，建议按照以下顺序调用接口：

1. **登录**：`POST /client/v1/login`
   ```http
   POST /client/v1/login
   Authorization: (无)
   Content-Type: application/json

   {
     "username": "testuser_1763969884",
     "password": "TestPassword123!"
   }
   ```
   响应中返回 `token`，后续接口均放在 `Authorization: Bearer <token>` 头中。

2. **获取业务类型**：`GET /client/v1/business_types`
   - 确认可用的 `business_code`（例如 `wx`），避免因为业务不存在导致 `60301 业务类型错误`。

3. **获取手机号**：`POST /client/v1/get_phone`
   ```http
   POST /client/v1/get_phone
   Authorization: Bearer <jwt>
   Content-Type: application/json

   {
     "business_type": "wx",
     "card_type": "physical",
     "count": 1
   }
   ```
   响应示例（真实返回）：
   ```json
   {
     "code": 200,
     "msg": "Success",
     "data": {
       "phones": [
         {
           "phone_number": "15820192778",
           "country_code": "CN",
           "cost": 0.01,
           "valid_until": "2025-11-25T12:25:02+08:00",
           "provider_id": "mqtt"
         }
       ],
       "total_cost": 0.01,
       "remaining_balance": 9996.99,
       "success_count": 1,
       "failed_count": 0
     }
   }
   ```

4. **获取验证码**：`POST /client/v1/get_code`
   ```http
   POST /client/v1/get_code
   Authorization: Bearer <jwt>
   Content-Type: application/json

   {
     "phone_numbers": ["15820192778"]
   }
   ```
   - 如果 `codes[].status` 为 `pending`，继续轮询此接口即可。

5. **可选**：`GET /client/v1/phone_status`、`GET /client/v1/assignments` 查询状态或历史。

> 提示：仓库中的 `tools/api_tester.go` 支持命令行测试，例如 `go run tools/api_tester.go phone`（仅获取手机号）或 `go run tools/api_tester.go code <phone>`（轮询验证码），便于联调。

---

## 4. 资金交易类型

资金相关接口（例如余额查询、交易历史检索、后续将开放的账单导出）都会使用 `sms_transactions.type` 字段来标识每一笔流水的动作。客户端在展示或对账时，可参考下表进行含义映射：

| 类型值 | 名称 | 含义 |
| :--- | :--- | :--- |
| `1` | 充值/入账 (TopUp) | 余额增加，例如人工上分或自动充值。 |
| `2` | 拉号消费 (Deduct) | 直接从可用余额扣款的消费记录。 |
| `3` | 拉号回退 (Refund) | 订单失败或超时导致的退款，金额退回可用余额。 |
| `4` | 上分 (Credit) | 系统/管理员加款，与充值类似（备用）。 |
| `5` | 下分 (Debit) | 系统/管理员扣款。 |
| `6` | 预冻结 (Freeze) | 从可用余额转入冻结余额，用于占用额度等待外部结果。 |
| `7` | 解冻 (Unfreeze) | 将冻结金额释放回可用余额，通常对应失败/超时场景。 |
| `8` | 冻结转实扣 (FreezeToCharge) | 预冻结成功转换为正式消费，不再影响可用余额，仅减少冻结额。 |

> 说明：随着新的余额并发保护方案上线，`6/7/8` 会被广泛使用；旧的直接扣款流程仍暂时保留，便于兼容历史数据。

## 5. 通用错误码
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
| 61202 | 服务商业务类型未配置（请联系平台管理员检查 `sms_providers_business_types` 配置） |
