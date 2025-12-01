# 商户业务配置Cost字段添加说明

## 修改日期
2024-12-01

## 修改目的
为 `sms_customer_business_config` 表添加 `cost` 字段，用于配置商户使用特定业务的单价成本。这样在给商户配置业务时，可以为不同商户设置不同的价格策略。

## 修改内容

### 1. 数据库模型修改

#### 1.1 Admin Server 模型 (`admin-server/server/plugin/sms/model/sms_customer_business_config.go`)
```go
Cost float64 `json:"cost" form:"cost" gorm:"column:cost;comment:业务成本（单价）;type:decimal(10,4);not null;default:0.0000;"`
```

#### 1.2 GoAPI Domain 模型 (`goapi/internal/domain/customer_business_config.go`)
```go
Cost float64 `gorm:"type:decimal(10,4);not null;default:0.0000;comment:业务成本（单价）" json:"cost"`
```

### 2. 数据库迁移

创建了迁移SQL文件: `admin-server/server/plugin/sms/migrations/add_cost_to_customer_business_config.sql`

```sql
ALTER TABLE `sms_customer_business_config` 
ADD COLUMN `cost` DECIMAL(10,4) NOT NULL DEFAULT 0.0000 COMMENT '业务成本（单价）' 
AFTER `business_name`;
```

### 3. 后端接口修改

#### 3.1 请求DTO (`admin-server/server/plugin/sms/model/request/sms_customers.go`)
在 `BusinessConfigItem` 结构体中添加:
```go
Cost float64 `json:"cost" binding:"required,min=0"`
```

#### 3.2 Service层 (`admin-server/server/plugin/sms/service/sms_customers.go`)
在 `ConfigureBusiness` 方法中添加Cost字段的保存:
```go
config := &model.SmsCustomerBusinessConfig{
    CustomerID:             req.CustomerID,
    PlatformBusinessTypeID: item.PlatformBusinessTypeID,
    BusinessCode:           item.BusinessCode,
    BusinessName:           item.BusinessName,
    Cost:                   item.Cost,  // 新增
    Weight:                 item.Weight,
    Status:                 item.Status,
}
```

### 4. 前端界面修改

#### 4.1 商户管理页面 (`admin-server/web/src/plugin/sms/view/smsCustomers.vue`)

添加了Cost字段的输入表单:
```vue
<el-form-item :label="`业务成本（单价）`" :prop="`businessConfig.${index}.cost`">
  <el-input-number 
    v-model="item.cost" 
    :min="0"
    :precision="4"
    :step="0.01"
    style="width: 100%"
  />
  <div class="text-xs text-gray-500 mt-1">设置该商户使用此业务的单次费用</div>
</el-form-item>
```

在业务类型选择时，默认cost为0:
```javascript
businessConfigForm.value.businessConfig.push({
  platformBusinessTypeId: type.ID,
  businessCode: type.code,
  businessName: type.name,
  cost: 0.0000,  // 新增默认值
  weight: 1,
  status: true
})
```

## 字段说明

- **字段名**: `cost`
- **字段类型**: `DECIMAL(10,4)`
- **默认值**: `0.0000`
- **是否必填**: 是
- **说明**: 业务成本（单价），表示商户使用该业务类型时每次的费用
- **精度**: 支持4位小数，可精确到0.0001
- **最小值**: 0（不允许负数）

## 使用场景

1. **差异化定价**: 可以为不同商户设置不同的业务价格
2. **成本核算**: 系统可以根据设置的cost计算商户的实际消费
3. **灵活计费**: 支持精确到小数点后4位的定价策略

## 影响范围

### 需要执行的操作

1. **数据库迁移**:
   ```bash
   # 在MySQL中执行迁移脚本
   mysql -u your_user -p sms_platform < admin-server/server/plugin/sms/migrations/add_cost_to_customer_business_config.sql
   ```

2. **重启后端服务**:
   - Admin Server 需要重启以加载新的模型定义
   - GoAPI 需要重启以加载新的domain定义

3. **前端重新构建** (如果需要):
   ```bash
   cd admin-server/web
   npm run build
   ```

### API变化

`POST /smsCustomers/configureBusiness` 接口的请求参数增加了 `cost` 字段:

**请求示例**:
```json
{
  "customerId": 1,
  "businessConfig": [
    {
      "platformBusinessTypeId": 1,
      "businessCode": "wx",
      "businessName": "微信",
      "cost": 0.5000,
      "weight": 10,
      "status": true
    }
  ]
}
```

## 注意事项

1. ✅ cost字段为必填项，最小值为0
2. ✅ 支持4位小数精度
3. ✅ 默认值为0.0000
4. ⚠️ 需要更新所有现有的商户业务配置，为其设置适当的cost值
5. ⚠️ GoAPI中如果有使用该配置的逻辑，需要相应更新以使用cost字段进行计费

## 测试建议

1. 测试创建新的商户业务配置，验证cost字段是否正确保存
2. 测试更新商户业务配置，验证cost字段是否正确更新
3. 测试cost字段的边界值（0, 0.0001, 9999999.9999）
4. 验证前端表单验证是否正常工作
5. 检查现有数据是否正确添加了默认值0.0000

## 后续工作建议

1. 在GoAPI的手机号获取逻辑中，使用customer_business_config中的cost进行实际扣费
2. 在商户报表中展示基于cost的消费统计
3. 考虑添加批量更新cost的功能
4. 添加cost历史记录，追踪价格变动
