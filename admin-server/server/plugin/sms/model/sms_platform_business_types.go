
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsPlatformBusinessTypes 平台业务 结构体
type SmsPlatformBusinessTypes struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:平台业务名称;column:name;size:255;"`  //平台业务名称
  Code  *string `json:"code" form:"code" gorm:"comment:平台业务编码;column:code;size:50;"`  //平台业务编码
  Description  *string `json:"description" form:"description" gorm:"comment:业务描述;column:description;size:500;"`  //业务描述
  Status  *bool `json:"status" form:"status" gorm:"comment:启用状态;column:status;"`  //启用状态
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:500;"`  //备注
}


// TableName 平台业务 SmsPlatformBusinessTypes自定义表名 sms_platform_business_types
func (SmsPlatformBusinessTypes) TableName() string {
    return "sms_platform_business_types"
}







