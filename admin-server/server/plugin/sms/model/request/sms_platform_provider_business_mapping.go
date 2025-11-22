
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type SmsPlatformProviderBusinessMappingSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       PlatformBusinessTypeId  *int `json:"platformBusinessTypeId" form:"platformBusinessTypeId"` 
       PlatformBusinessCode  *string `json:"platformBusinessCode" form:"platformBusinessCode"` 
       ProviderBusinessTypeId  *int `json:"providerBusinessTypeId" form:"providerBusinessTypeId"` 
       ProviderCode  *string `json:"providerCode" form:"providerCode"` 
    request.PageInfo
}
