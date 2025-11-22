import service from '@/utils/request'
// @Tags SmsPlatformBusinessTypes
// @Summary 创建平台业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformBusinessTypes true "创建平台业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsPlatformBusinessTypes/createSmsPlatformBusinessTypes [post]
export const createSmsPlatformBusinessTypes = (data) => {
  return service({
    url: '/smsPlatformBusinessTypes/createSmsPlatformBusinessTypes',
    method: 'post',
    data
  })
}

// @Tags SmsPlatformBusinessTypes
// @Summary 删除平台业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformBusinessTypes true "删除平台业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsPlatformBusinessTypes/deleteSmsPlatformBusinessTypes [delete]
export const deleteSmsPlatformBusinessTypes = (params) => {
  return service({
    url: '/smsPlatformBusinessTypes/deleteSmsPlatformBusinessTypes',
    method: 'delete',
    params
  })
}

// @Tags SmsPlatformBusinessTypes
// @Summary 批量删除平台业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除平台业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsPlatformBusinessTypes/deleteSmsPlatformBusinessTypes [delete]
export const deleteSmsPlatformBusinessTypesByIds = (params) => {
  return service({
    url: '/smsPlatformBusinessTypes/deleteSmsPlatformBusinessTypesByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsPlatformBusinessTypes
// @Summary 更新平台业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformBusinessTypes true "更新平台业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsPlatformBusinessTypes/updateSmsPlatformBusinessTypes [put]
export const updateSmsPlatformBusinessTypes = (data) => {
  return service({
    url: '/smsPlatformBusinessTypes/updateSmsPlatformBusinessTypes',
    method: 'put',
    data
  })
}

// @Tags SmsPlatformBusinessTypes
// @Summary 用id查询平台业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsPlatformBusinessTypes true "用id查询平台业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsPlatformBusinessTypes/findSmsPlatformBusinessTypes [get]
export const findSmsPlatformBusinessTypes = (params) => {
  return service({
    url: '/smsPlatformBusinessTypes/findSmsPlatformBusinessTypes',
    method: 'get',
    params
  })
}

// @Tags SmsPlatformBusinessTypes
// @Summary 分页获取平台业务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取平台业务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsPlatformBusinessTypes/getSmsPlatformBusinessTypesList [get]
export const getSmsPlatformBusinessTypesList = (params) => {
  return service({
    url: '/smsPlatformBusinessTypes/getSmsPlatformBusinessTypesList',
    method: 'get',
    params
  })
}
// @Tags SmsPlatformBusinessTypes
// @Summary 不需要鉴权的平台业务接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsPlatformBusinessTypesSearch true "分页获取平台业务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsPlatformBusinessTypes/getSmsPlatformBusinessTypesPublic [get]
export const getSmsPlatformBusinessTypesPublic = () => {
  return service({
    url: '/smsPlatformBusinessTypes/getSmsPlatformBusinessTypesPublic',
    method: 'get',
  })
}
