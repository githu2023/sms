import service from '@/utils/request'
// @Tags SmsApiLogs
// @Summary 创建访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsApiLogs true "创建访问日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsApiLogs/createSmsApiLogs [post]
export const createSmsApiLogs = (data) => {
  return service({
    url: '/smsApiLogs/createSmsApiLogs',
    method: 'post',
    data
  })
}

// @Tags SmsApiLogs
// @Summary 删除访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsApiLogs true "删除访问日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsApiLogs/deleteSmsApiLogs [delete]
export const deleteSmsApiLogs = (params) => {
  return service({
    url: '/smsApiLogs/deleteSmsApiLogs',
    method: 'delete',
    params
  })
}

// @Tags SmsApiLogs
// @Summary 批量删除访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除访问日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsApiLogs/deleteSmsApiLogs [delete]
export const deleteSmsApiLogsByIds = (params) => {
  return service({
    url: '/smsApiLogs/deleteSmsApiLogsByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsApiLogs
// @Summary 更新访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsApiLogs true "更新访问日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsApiLogs/updateSmsApiLogs [put]
export const updateSmsApiLogs = (data) => {
  return service({
    url: '/smsApiLogs/updateSmsApiLogs',
    method: 'put',
    data
  })
}

// @Tags SmsApiLogs
// @Summary 用id查询访问日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsApiLogs true "用id查询访问日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsApiLogs/findSmsApiLogs [get]
export const findSmsApiLogs = (params) => {
  return service({
    url: '/smsApiLogs/findSmsApiLogs',
    method: 'get',
    params
  })
}

// @Tags SmsApiLogs
// @Summary 分页获取访问日志列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取访问日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsApiLogs/getSmsApiLogsList [get]
export const getSmsApiLogsList = (params) => {
  return service({
    url: '/smsApiLogs/getSmsApiLogsList',
    method: 'get',
    params
  })
}
// @Tags SmsApiLogs
// @Summary 不需要鉴权的访问日志接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsApiLogsSearch true "分页获取访问日志列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsApiLogs/getSmsApiLogsPublic [get]
export const getSmsApiLogsPublic = () => {
  return service({
    url: '/smsApiLogs/getSmsApiLogsPublic',
    method: 'get',
  })
}
