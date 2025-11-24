
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type SmsPhoneAssignmentsSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       CustomerId  *int `json:"customerId" form:"customerId"` 
       BusinessTypeId  *int `json:"businessTypeId" form:"businessTypeId"` 
       PhoneNumber  *string `json:"phoneNumber" form:"phoneNumber"` 
       MerchantNo  *string `json:"merchantNo" form:"merchantNo"` 
       BusinessCode  *string `json:"businessCode" form:"businessCode"` 
       Status  *string `json:"status" form:"status"` 
    request.PageInfo
}
