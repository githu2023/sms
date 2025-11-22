
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
    request.PageInfo
}
