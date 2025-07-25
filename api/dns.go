package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiamxu/aliops/service"
	"github.com/tiamxu/aliops/types"
)

type DNSHandler struct {
	service *service.DNSService
}

func NewDNSHandler(service *service.DNSService) *DNSHandler {
	return &DNSHandler{service: service}
}

func (h *DNSHandler) Add(c *gin.Context) {
	var req types.DomainRecordAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recordId, err := h.service.Add(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"record_id": recordId,
		"message":   "记录创建成功",
	})
}

func (h *DNSHandler) Delete(c *gin.Context) {

}
func (h *DNSHandler) Update(c *gin.Context) {

}
func (h *DNSHandler) List(c *gin.Context) {
	domain := c.Query("domain")
	if domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "domain参数必填"})
		return
	}

	result, err := h.service.List(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
