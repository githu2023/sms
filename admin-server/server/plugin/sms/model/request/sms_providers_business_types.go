package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SmsProvidersBusinessTypesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ProviderCode   string     `json:"providerCode" form:"providerCode" `
	BusinessCode   string     `json:"businessCode" form:"businessCode" `
	Status         *bool      `json:"status" form:"status" `
	request.PageInfo
}
