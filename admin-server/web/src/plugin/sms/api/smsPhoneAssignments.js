import service from '@/utils/request'
// @Tags SmsPhoneAssignments
// @Summary 创建号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPhoneAssignments true "创建号码记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsPhoneAssignments/createSmsPhoneAssignments [post]
export const createSmsPhoneAssignments = (data) => {
  return service({
    url: '/smsPhoneAssignments/createSmsPhoneAssignments',
    method: 'post',
    data
  })
}

// @Tags SmsPhoneAssignments
// @Summary 删除号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPhoneAssignments true "删除号码记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsPhoneAssignments/deleteSmsPhoneAssignments [delete]
export const deleteSmsPhoneAssignments = (params) => {
  return service({
    url: '/smsPhoneAssignments/deleteSmsPhoneAssignments',
    method: 'delete',
    params
  })
}

// @Tags SmsPhoneAssignments
// @Summary 批量删除号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除号码记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsPhoneAssignments/deleteSmsPhoneAssignments [delete]
export const deleteSmsPhoneAssignmentsByIds = (params) => {
  return service({
    url: '/smsPhoneAssignments/deleteSmsPhoneAssignmentsByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsPhoneAssignments
// @Summary 更新号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsPhoneAssignments true "更新号码记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsPhoneAssignments/updateSmsPhoneAssignments [put]
export const updateSmsPhoneAssignments = (data) => {
  return service({
    url: '/smsPhoneAssignments/updateSmsPhoneAssignments',
    method: 'put',
    data
  })
}

// @Tags SmsPhoneAssignments
// @Summary 用id查询号码记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsPhoneAssignments true "用id查询号码记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsPhoneAssignments/findSmsPhoneAssignments [get]
export const findSmsPhoneAssignments = (params) => {
  return service({
    url: '/smsPhoneAssignments/findSmsPhoneAssignments',
    method: 'get',
    params
  })
}

// @Tags SmsPhoneAssignments
// @Summary 分页获取号码记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取号码记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsPhoneAssignments/getSmsPhoneAssignmentsList [get]
export const getSmsPhoneAssignmentsList = (params) => {
  return service({
    url: '/smsPhoneAssignments/getSmsPhoneAssignmentsList',
    method: 'get',
    params
  })
}
// @Tags SmsPhoneAssignments
// @Summary 不需要鉴权的号码记录接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsPhoneAssignmentsSearch true "分页获取号码记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsPhoneAssignments/getSmsPhoneAssignmentsPublic [get]
export const getSmsPhoneAssignmentsPublic = () => {
  return service({
    url: '/smsPhoneAssignments/getSmsPhoneAssignmentsPublic',
    method: 'get',
  })
}
