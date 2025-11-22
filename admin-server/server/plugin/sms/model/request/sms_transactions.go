package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SmsTransactionsSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	CustomerId     *int        `json:"customerId" form:"customerId"`
	Type           *string     `json:"type" form:"type"`
	request.PageInfo
}
