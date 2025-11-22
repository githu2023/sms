
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsPlatformProviderBusinessMapping 平台子业务 结构体
type SmsPlatformProviderBusinessMapping struct {
    global.GVA_MODEL
  PlatformBusinessTypeId  *int64 `json:"platformBusinessTypeId" form:"platformBusinessTypeId" gorm:"comment:平台业务ID（关联sms_platform_business_types表的ID）;column:platform_business_type_id;"`  //平台业务ID
  PlatformBusinessCode  *string `json:"platformBusinessCode" form:"platformBusinessCode" gorm:"comment:平台业务编码;column:platform_business_code;size:50;"`  //平台业务编码
  ProviderBusinessTypeId  *int64 `json:"providerBusinessTypeId" form:"providerBusinessTypeId" gorm:"comment:三方业务ID（关联sms_providers_business_types表的ID）;column:provider_business_type_id;"`  //三方业务ID
  ProviderCode  *string `json:"providerCode" form:"providerCode" gorm:"comment:三方编码;column:provider_code;size:50;"`  //三方编码
  BusinessCode  *string `json:"businessCode" form:"businessCode" gorm:"comment:三方业务编码;column:business_code;size:50;"`  //三方业务编码
  Weight  *int32 `json:"weight" form:"weight" gorm:"comment:权重（用于随机选择，权重越高被选中概率越大）;column:weight;"`  //权重
  Status  *bool `json:"status" form:"status" gorm:"comment:是否启用该映射;column:status;"`  //是否启用该映射
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:500;"`  //备注
}


// TableName 平台子业务 SmsPlatformProviderBusinessMapping自定义表名 sms_platform_provider_business_mapping
func (SmsPlatformProviderBusinessMapping) TableName() string {
    return "sms_platform_provider_business_mapping"
}







