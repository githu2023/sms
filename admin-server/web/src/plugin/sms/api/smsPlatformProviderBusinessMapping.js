import service from '@/utils/request'
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 创建平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformProviderBusinessMapping true "创建平台子业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsPlatformProviderBusinessMapping/createSmsPlatformProviderBusinessMapping [post]
export const createSmsPlatformProviderBusinessMapping = (data) => {
  return service({
    url: '/smsPlatformProviderBusinessMapping/createSmsPlatformProviderBusinessMapping',
    method: 'post',
    data
  })
}

// @Tags SmsPlatformProviderBusinessMapping
// @Summary 删除平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformProviderBusinessMapping true "删除平台子业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMapping [delete]
export const deleteSmsPlatformProviderBusinessMapping = (params) => {
  return service({
    url: '/smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMapping',
    method: 'delete',
    params
  })
}

// @Tags SmsPlatformProviderBusinessMapping
// @Summary 批量删除平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除平台子业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMapping [delete]
export const deleteSmsPlatformProviderBusinessMappingByIds = (params) => {
  return service({
    url: '/smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMappingByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsPlatformProviderBusinessMapping
// @Summary 更新平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPlatformProviderBusinessMapping true "更新平台子业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsPlatformProviderBusinessMapping/updateSmsPlatformProviderBusinessMapping [put]
export const updateSmsPlatformProviderBusinessMapping = (data) => {
  return service({
    url: '/smsPlatformProviderBusinessMapping/updateSmsPlatformProviderBusinessMapping',
    method: 'put',
    data
  })
}

// @Tags SmsPlatformProviderBusinessMapping
// @Summary 用id查询平台子业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsPlatformProviderBusinessMapping true "用id查询平台子业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsPlatformProviderBusinessMapping/findSmsPlatformProviderBusinessMapping [get]
export const findSmsPlatformProviderBusinessMapping = (params) => {
  return service({
    url: '/smsPlatformProviderBusinessMapping/findSmsPlatformProviderBusinessMapping',
    method: 'get',
    params
  })
}

// @Tags SmsPlatformProviderBusinessMapping
// @Summary 分页获取平台子业务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取平台子业务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsPlatformProviderBusinessMapping/getSmsPlatformProviderBusinessMappingList [get]
export const getSmsPlatformProviderBusinessMappingList = (params) => {
  return service({
    url: '/smsPlatformProviderBusinessMapping/getSmsPlatformProviderBusinessMappingList',
    method: 'get',
    params
  })
}
// @Tags SmsPlatformProviderBusinessMapping
// @Summary 不需要鉴权的平台子业务接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsPlatformProviderBusinessMappingSearch true "分页获取平台子业务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsPlatformProviderBusinessMapping/getSmsPlatformProviderBusinessMappingPublic [get]
export const getSmsPlatformProviderBusinessMappingPublic = () => {
  return service({
    url: '/smsPlatformProviderBusinessMapping/getSmsPlatformProviderBusinessMappingPublic',
    method: 'get',
  })
}
