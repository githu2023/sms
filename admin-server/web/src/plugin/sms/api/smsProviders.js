import service from '@/utils/request'
// @Tags SmsProviders
// @Summary 创建服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProviders true "创建服务端"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsProviders/createSmsProviders [post]
export const createSmsProviders = (data) => {
  return service({
    url: '/smsProviders/createSmsProviders',
    method: 'post',
    data
  })
}

// @Tags SmsProviders
// @Summary 删除服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProviders true "删除服务端"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsProviders/deleteSmsProviders [delete]
export const deleteSmsProviders = (params) => {
  return service({
    url: '/smsProviders/deleteSmsProviders',
    method: 'delete',
    params
  })
}

// @Tags SmsProviders
// @Summary 批量删除服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除服务端"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsProviders/deleteSmsProviders [delete]
export const deleteSmsProvidersByIds = (params) => {
  return service({
    url: '/smsProviders/deleteSmsProvidersByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsProviders
// @Summary 更新服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProviders true "更新服务端"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsProviders/updateSmsProviders [put]
export const updateSmsProviders = (data) => {
  return service({
    url: '/smsProviders/updateSmsProviders',
    method: 'put',
    data
  })
}

// @Tags SmsProviders
// @Summary 用id查询服务端
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsProviders true "用id查询服务端"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsProviders/findSmsProviders [get]
export const findSmsProviders = (params) => {
  return service({
    url: '/smsProviders/findSmsProviders',
    method: 'get',
    params
  })
}

// @Tags SmsProviders
// @Summary 分页获取服务端列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取服务端列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsProviders/getSmsProvidersList [get]
export const getSmsProvidersList = (params) => {
  return service({
    url: '/smsProviders/getSmsProvidersList',
    method: 'get',
    params
  })
}
// @Tags SmsProviders
// @Summary 不需要鉴权的服务端接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsProvidersSearch true "分页获取服务端列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsProviders/getSmsProvidersPublic [get]
export const getSmsProvidersPublic = () => {
  return service({
    url: '/smsProviders/getSmsProvidersPublic',
    method: 'get',
  })
}
