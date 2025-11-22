package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service"

var (
	Api                                       = new(api)
	serviceSmsApiLogs                         = service.Service.SmsApiLogs
	serviceSmsCustomers                       = service.Service.SmsCustomers
	serviceSmsIpWhitelist                     = service.Service.SmsIpWhitelist
	serviceSmsPhoneAssignments                = service.Service.SmsPhoneAssignments
	serviceSmsProviders                       = service.Service.SmsProviders
	serviceSmsTransactions                    = service.Service.SmsTransactions
	serviceSmsProvidersBusinessTypes          = service.Service.SmsProvidersBusinessTypes
	serviceSmsPlatformBusinessTypes           = service.Service.SmsPlatformBusinessTypes
	serviceSmsPlatformProviderBusinessMapping = service.Service.SmsPlatformProviderBusinessMapping
)

type api struct {
	SmsApiLogs                         smsApiLogs
	SmsCustomers                       smsCustomers
	SmsIpWhitelist                     smsIpWhitelist
	SmsPhoneAssignments                smsPhoneAssignments
	SmsProviders                       smsProviders
	SmsTransactions                    smsTransactions
	SmsProvidersBusinessTypes          smsProvidersBusinessTypes
	SmsPlatformBusinessTypes           smsPlatformBusinessTypes
	SmsPlatformProviderBusinessMapping smsPlatformProviderBusinessMapping
}
