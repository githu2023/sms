
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type SmsApiLogsSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       CustomerId  *int `json:"customerId" form:"customerId"` 
    request.PageInfo
}
