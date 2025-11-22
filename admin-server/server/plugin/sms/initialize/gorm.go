package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(
		model.SmsApiLogs{},
		model.SmsCustomers{},
		model.SmsCustomerBusinessConfig{},
		model.SmsIpWhitelist{},
		model.SmsPhoneAssignments{},
		model.SmsProviders{},
		model.SmsTransactions{},
		model.SmsProvidersBusinessTypes{},
		model.SmsPlatformBusinessTypes{},
		model.SmsPlatformProviderBusinessMapping{},
	)
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
