package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "sms", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(new(model.SmsApiLogs), new(model.SmsBusinessTypes), new(model.SmsCustomers), new(model.SmsIpWhitelist), new(model.SmsPhoneAssignments), new(model.SmsProviders), new(model.SmsTransactions), new(model.SmsProvidersBusinessTypes), new(model.SmsPlatformBusinessTypes), //go:generate go mod tidy
		//go:generate go mod download
		//go:generate go run gen.go

		new(model.SmsPlatformProviderBusinessMapping),
	)
	g.Execute()
}
