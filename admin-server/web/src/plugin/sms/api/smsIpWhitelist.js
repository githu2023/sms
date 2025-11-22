import service from '@/utils/request'
// @Tags SmsIpWhitelist
// @Summary 创建白名单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsIpWhitelist true "创建白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsIpWhitelist/createSmsIpWhitelist [post]
export const createSmsIpWhitelist = (data) => {
  return service({
    url: '/smsIpWhitelist/createSmsIpWhitelist',
    method: 'post',
    data
  })
}

// @Tags SmsIpWhitelist
// @Summary 删除白名单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsIpWhitelist true "删除白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsIpWhitelist/deleteSmsIpWhitelist [delete]
export const deleteSmsIpWhitelist = (params) => {
  return service({
    url: '/smsIpWhitelist/deleteSmsIpWhitelist',
    method: 'delete',
    params
  })
}

// @Tags SmsIpWhitelist
// @Summary 批量删除白名单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsIpWhitelist/deleteSmsIpWhitelist [delete]
export const deleteSmsIpWhitelistByIds = (params) => {
  return service({
    url: '/smsIpWhitelist/deleteSmsIpWhitelistByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsIpWhitelist
// @Summary 更新白名单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsIpWhitelist true "更新白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsIpWhitelist/updateSmsIpWhitelist [put]
export const updateSmsIpWhitelist = (data) => {
  return service({
    url: '/smsIpWhitelist/updateSmsIpWhitelist',
    method: 'put',
    data
  })
}

// @Tags SmsIpWhitelist
// @Summary 用id查询白名单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsIpWhitelist true "用id查询白名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsIpWhitelist/findSmsIpWhitelist [get]
export const findSmsIpWhitelist = (params) => {
  return service({
    url: '/smsIpWhitelist/findSmsIpWhitelist',
    method: 'get',
    params
  })
}

// @Tags SmsIpWhitelist
// @Summary 分页获取白名单列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取白名单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsIpWhitelist/getSmsIpWhitelistList [get]
export const getSmsIpWhitelistList = (params) => {
  return service({
    url: '/smsIpWhitelist/getSmsIpWhitelistList',
    method: 'get',
    params
  })
}
// @Tags SmsIpWhitelist
// @Summary 不需要鉴权的白名单接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsIpWhitelistSearch true "分页获取白名单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsIpWhitelist/getSmsIpWhitelistPublic [get]
export const getSmsIpWhitelistPublic = () => {
  return service({
    url: '/smsIpWhitelist/getSmsIpWhitelistPublic',
    method: 'get',
  })
}
