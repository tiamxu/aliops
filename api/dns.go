package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiamxu/aliops/service"
)

type DNSHandler struct {
	service *service.DNSService
}

func NewDNSHandler(service *service.DNSService) *DNSHandler {
	return &DNSHandler{service: service}
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
