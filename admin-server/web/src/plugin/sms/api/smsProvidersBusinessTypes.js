import service from '@/utils/request'
// @Tags SmsProvidersBusinessTypes
// @Summary 创建三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProvidersBusinessTypes true "创建三方业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsProvidersBusinessTypes/createSmsProvidersBusinessTypes [post]
export const createSmsProvidersBusinessTypes = (data) => {
  return service({
    url: '/smsProvidersBusinessTypes/createSmsProvidersBusinessTypes',
    method: 'post',
    data
  })
}

// @Tags SmsProvidersBusinessTypes
// @Summary 删除三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProvidersBusinessTypes true "删除三方业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypes [delete]
export const deleteSmsProvidersBusinessTypes = (params) => {
  return service({
    url: '/smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypes',
    method: 'delete',
    params
  })
}

// @Tags SmsProvidersBusinessTypes
// @Summary 批量删除三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除三方业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypes [delete]
export const deleteSmsProvidersBusinessTypesByIds = (params) => {
  return service({
    url: '/smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypesByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsProvidersBusinessTypes
// @Summary 更新三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsProvidersBusinessTypes true "更新三方业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsProvidersBusinessTypes/updateSmsProvidersBusinessTypes [put]
export const updateSmsProvidersBusinessTypes = (data) => {
  return service({
    url: '/smsProvidersBusinessTypes/updateSmsProvidersBusinessTypes',
    method: 'put',
    data
  })
}

// @Tags SmsProvidersBusinessTypes
// @Summary 用id查询三方业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsProvidersBusinessTypes true "用id查询三方业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsProvidersBusinessTypes/findSmsProvidersBusinessTypes [get]
export const findSmsProvidersBusinessTypes = (params) => {
  return service({
    url: '/smsProvidersBusinessTypes/findSmsProvidersBusinessTypes',
    method: 'get',
    params
  })
}

// @Tags SmsProvidersBusinessTypes
// @Summary 分页获取三方业务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取三方业务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsProvidersBusinessTypes/getSmsProvidersBusinessTypesList [get]
export const getSmsProvidersBusinessTypesList = (params) => {
  return service({
    url: '/smsProvidersBusinessTypes/getSmsProvidersBusinessTypesList',
    method: 'get',
    params
  })
}
// @Tags SmsProvidersBusinessTypes
// @Summary 不需要鉴权的三方业务接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsProvidersBusinessTypesSearch true "分页获取三方业务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsProvidersBusinessTypes/getSmsProvidersBusinessTypesPublic [get]
export const getSmsProvidersBusinessTypesPublic = () => {
  return service({
    url: '/smsProvidersBusinessTypes/getSmsProvidersBusinessTypesPublic',
    method: 'get',
  })
}
