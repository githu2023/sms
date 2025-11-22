package service

var Service = new(service)

type service struct {
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
