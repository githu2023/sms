
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsIpWhitelist 白名单 结构体
type SmsIpWhitelist struct {
    global.GVA_MODEL
  CustomerId  *int64 `json:"customerId" form:"customerId" gorm:"comment:客户ID;column:customer_id;"`  //客户ID
  IpAddress  *string `json:"ipAddress" form:"ipAddress" gorm:"comment:白名单IP或IP段;column:ip_address;size:45;"`  //白名单IP或IP段
  Notes  *string `json:"notes" form:"notes" gorm:"comment:备注, 例如 "办公室IP";column:notes;size:255;"`  //备注, 例如 "办公室IP"
}


// TableName 白名单 SmsIpWhitelist自定义表名 sms_ip_whitelist
func (SmsIpWhitelist) TableName() string {
    return "sms_ip_whitelist"
}







