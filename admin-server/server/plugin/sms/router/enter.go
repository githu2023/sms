package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api"

var (
	Router                                = new(router)
	apiSmsApiLogs                         = api.Api.SmsApiLogs
	apiSmsCustomers                       = api.Api.SmsCustomers
	apiSmsIpWhitelist                     = api.Api.SmsIpWhitelist
	apiSmsPhoneAssignments                = api.Api.SmsPhoneAssignments
	apiSmsProviders                       = api.Api.SmsProviders
	apiSmsTransactions                    = api.Api.SmsTransactions
	apiSmsProvidersBusinessTypes          = api.Api.SmsProvidersBusinessTypes
	apiSmsPlatformBusinessTypes           = api.Api.SmsPlatformBusinessTypes
	apiSmsPlatformProviderBusinessMapping = api.Api.SmsPlatformProviderBusinessMapping
)

type router struct {
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
