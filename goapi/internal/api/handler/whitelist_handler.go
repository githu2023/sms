package handler

import (
	"net/http"
	"strconv"

	"sms-platform/goapi/internal/dto"
	"sms-platform/goapi/internal/service"

	"github.com/gin-gonic/gin"
)

type WhitelistHandler struct {
	service service.WhitelistService
}

func NewWhitelistHandler(s service.WhitelistService) *WhitelistHandler {
	return &WhitelistHandler{service: s}
}

// AddWhitelist 添加白名单
func (h *WhitelistHandler) AddWhitelist(c *gin.Context) {
	var req dto.WhitelistCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	customerID, err := getCustomerIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}
	if err := h.service.AddWhitelist(c.Request.Context(), customerID, req.IPAddress, req.Notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// DeleteWhitelist 删除白名单
func (h *WhitelistHandler) DeleteWhitelist(c *gin.Context) {
	var req dto.WhitelistDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	customerID, err := getCustomerIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}
	if err := h.service.DeleteWhitelist(c.Request.Context(), customerID, req.IPAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// ListWhitelists 查询白名单列表
func (h *WhitelistHandler) ListWhitelists(c *gin.Context) {
	var req dto.WhitelistListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 20
	}
	customerID, err := getCustomerIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}
	list, total, err := h.service.ListWhitelists(c.Request.Context(), customerID, req.Page, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]dto.WhitelistResponse, 0, len(list))
	for _, wl := range list {
		resp = append(resp, dto.WhitelistResponse{
			ID:         wl.ID,
			CustomerID: wl.CustomerID,
			IPAddress:  wl.IPAddress,
			Notes:      wl.Notes,
			CreatedAt:  wl.CreatedAt.Unix(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"total": total, "list": resp})
}

// getCustomerIDFromContext 从上下文获取客户ID（假设已实现）
func getCustomerIDFromContext(c *gin.Context) (int64, error) {
	// 这里假设客户ID存储在 JWT 或 session 中，实际项目请根据认证实现
	idStr := c.GetHeader("X-Customer-ID")
	if idStr == "" {
		return 0, http.ErrNoCookie
	}
	return strconv.ParseInt(idStr, 10, 64)
}
