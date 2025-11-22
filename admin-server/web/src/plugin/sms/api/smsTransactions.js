import service from '@/utils/request'
// @Tags SmsTransactions
// @Summary 创建交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsTransactions true "创建交易记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsTransactions/createSmsTransactions [post]
export const createSmsTransactions = (data) => {
  return service({
    url: '/smsTransactions/createSmsTransactions',
    method: 'post',
    data
  })
}

// @Tags SmsTransactions
// @Summary 删除交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsTransactions true "删除交易记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsTransactions/deleteSmsTransactions [delete]
export const deleteSmsTransactions = (params) => {
  return service({
    url: '/smsTransactions/deleteSmsTransactions',
    method: 'delete',
    params
  })
}

// @Tags SmsTransactions
// @Summary 批量删除交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除交易记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsTransactions/deleteSmsTransactions [delete]
export const deleteSmsTransactionsByIds = (params) => {
  return service({
    url: '/smsTransactions/deleteSmsTransactionsByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsTransactions
// @Summary 更新交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsTransactions true "更新交易记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsTransactions/updateSmsTransactions [put]
export const updateSmsTransactions = (data) => {
  return service({
    url: '/smsTransactions/updateSmsTransactions',
    method: 'put',
    data
  })
}

// @Tags SmsTransactions
// @Summary 用id查询交易记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsTransactions true "用id查询交易记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsTransactions/findSmsTransactions [get]
export const findSmsTransactions = (params) => {
  return service({
    url: '/smsTransactions/findSmsTransactions',
    method: 'get',
    params
  })
}

// @Tags SmsTransactions
// @Summary 分页获取交易记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取交易记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsTransactions/getSmsTransactionsList [get]
export const getSmsTransactionsList = (params) => {
  return service({
    url: '/smsTransactions/getSmsTransactionsList',
    method: 'get',
    params
  })
}
// @Tags SmsTransactions
// @Summary 不需要鉴权的交易记录接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsTransactionsSearch true "分页获取交易记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsTransactions/getSmsTransactionsPublic [get]
export const getSmsTransactionsPublic = () => {
  return service({
    url: '/smsTransactions/getSmsTransactionsPublic',
    method: 'get',
  })
}
