package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	public.Use()
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	router.Router.SmsApiLogs.Init(public, private)
	router.Router.SmsCustomers.Init(public, private)
	router.Router.SmsIpWhitelist.Init(public, private)
	router.Router.SmsPhoneAssignments.Init(public, private)
	router.Router.SmsProviders.Init(public, private)
	router.Router.SmsTransactions.Init(public, private)
	router.Router.SmsProvidersBusinessTypes.Init(public, private)
	router.Router.SmsPlatformBusinessTypes.Init(public, private)
	router.Router.SmsPlatformProviderBusinessMapping.Init(public, private)
}
