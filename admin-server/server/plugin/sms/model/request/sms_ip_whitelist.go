
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type SmsIpWhitelistSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       CustomerId  *int `json:"customerId" form:"customerId"` 
    request.PageInfo
}
