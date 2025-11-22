
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SmsApiLogs 访问日志 结构体
type SmsApiLogs struct {
    global.GVA_MODEL
  CustomerId  *int64 `json:"customerId" form:"customerId" gorm:"comment:客户ID;column:customer_id;"`  //客户ID
  RequestIp  *string `json:"requestIp" form:"requestIp" gorm:"comment:请求来源IP;column:request_ip;size:45;"`  //请求来源IP
  RequestPath  *string `json:"requestPath" form:"requestPath" gorm:"comment:请求的API路径;column:request_path;size:255;"`  //请求的API路径
  RequestBody  *string `json:"requestBody" form:"requestBody" gorm:"comment:请求体内容;column:request_body;"`  //请求体内容
  ResponseCode  *int32 `json:"responseCode" form:"responseCode" gorm:"comment:HTTP响应状态码;column:response_code;"`  //HTTP响应状态码
  DurationMs  *int32 `json:"durationMs" form:"durationMs" gorm:"comment:请求处理耗时(毫秒);column:duration_ms;"`  //请求处理耗时(毫秒)
}


// TableName 访问日志 SmsApiLogs自定义表名 sms_api_logs
func (SmsApiLogs) TableName() string {
    return "sms_api_logs"
}







