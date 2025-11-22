package dto

// WhitelistCreateRequest 创建白名单请求
// @Description 创建IP白名单
// @Param ip_address string true "IP地址或CIDR"
// @Param notes string false "备注"
type WhitelistCreateRequest struct {
	IPAddress string `json:"ip_address" binding:"required"`
	Notes     string `json:"notes"`
}

// WhitelistDeleteRequest 删除白名单请求
type WhitelistDeleteRequest struct {
	IPAddress string `json:"ip_address" binding:"required"`
}

// WhitelistListRequest 查询白名单列表请求
type WhitelistListRequest struct {
	Page  int `form:"page" binding:"min=1"`
	Limit int `form:"limit" binding:"min=1,max=100"`
}

// WhitelistResponse 白名单响应
type WhitelistResponse struct {
	ID         int64  `json:"id"`
	CustomerID int64  `json:"customer_id"`
	IPAddress  string `json:"ip_address"`
	Notes      string `json:"notes"`
	CreatedAt  int64  `json:"created_at"`
}
