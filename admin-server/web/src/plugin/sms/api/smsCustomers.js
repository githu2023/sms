import service from '@/utils/request'
// @Tags SmsCustomers
// @Summary 创建商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsCustomers true "创建商户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /smsCustomers/createSmsCustomers [post]
export const createSmsCustomers = (data) => {
  return service({
    url: '/smsCustomers/createSmsCustomers',
    method: 'post',
    data
  })
}

// @Tags SmsCustomers
// @Summary 删除商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsCustomers true "删除商户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsCustomers/deleteSmsCustomers [delete]
export const deleteSmsCustomers = (params) => {
  return service({
    url: '/smsCustomers/deleteSmsCustomers',
    method: 'delete',
    params
  })
}

// @Tags SmsCustomers
// @Summary 批量删除商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smsCustomers/deleteSmsCustomers [delete]
export const deleteSmsCustomersByIds = (params) => {
  return service({
    url: '/smsCustomers/deleteSmsCustomersByIds',
    method: 'delete',
    params
  })
}

// @Tags SmsCustomers
// @Summary 更新商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SmsCustomers true "更新商户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smsCustomers/updateSmsCustomers [put]
export const updateSmsCustomers = (data) => {
  return service({
    url: '/smsCustomers/updateSmsCustomers',
    method: 'put',
    data
  })
}

// @Tags SmsCustomers
// @Summary 用id查询商户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SmsCustomers true "用id查询商户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smsCustomers/findSmsCustomers [get]
export const findSmsCustomers = (params) => {
  return service({
    url: '/smsCustomers/findSmsCustomers',
    method: 'get',
    params
  })
}

// @Tags SmsCustomers
// @Summary 分页获取商户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smsCustomers/getSmsCustomersList [get]
export const getSmsCustomersList = (params) => {
  return service({
    url: '/smsCustomers/getSmsCustomersList',
    method: 'get',
    params
  })
}
// @Tags SmsCustomers
// @Summary 不需要鉴权的商户接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SmsCustomersSearch true "分页获取商户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /smsCustomers/getSmsCustomersPublic [get]
export const getSmsCustomersPublic = () => {
  return service({
    url: '/smsCustomers/getSmsCustomersPublic',
    method: 'get',
  })
}

// @Tags SmsCustomers
// @Summary 充值/扣费
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.CreditDebitSmsCustomersReq true "充值/扣费"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /smsCustomers/creditDebit [post]
export const creditDebitSmsCustomers = (data) => {
  return service({
    url: '/smsCustomers/creditDebit',
    method: 'post',
    data
  })
}

// @Tags SmsCustomers
// @Summary 配置商户业务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.ConfigureBusinessReq true "配置商户业务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"配置成功"}"
// @Router /smsCustomers/configureBusiness [post]
export const configureBusinessSmsCustomers = (data) => {
  return service({
    url: '/smsCustomers/configureBusiness',
    method: 'post',
    data
  })
}

// @Tags SmsCustomers
// @Summary 调整冻结金额
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.AdjustFrozenAmountReq true "调整冻结金额"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"调整成功"}"
// @Router /smsCustomers/adjustFrozenAmount [post]
export const adjustFrozenAmountSmsCustomers = (data) => {
  return service({
    url: '/smsCustomers/adjustFrozenAmount',
    method: 'post',
    data
  })
}

// @Tags SmsCustomers
// @Summary 获取商户业务配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param customerId query int true "商户ID"
// @Success 200 {object} response.Response{data=[]model.SmsCustomerBusinessConfig} "获取成功"
// @Router /smsCustomers/getBusinessConfig [get]
export const getBusinessConfigSmsCustomers = (params) => {
  return service({
    url: '/smsCustomers/getBusinessConfig',
    method: 'get',
    params
  })
}
