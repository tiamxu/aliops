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
		c.JSON(http.StatusBadRequest, RespError(c, err, ""))
		return
	}
	recordId, err := h.service.Add(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RespError(c, err, "添加DNS记录失败"))
		return
	}
	c.JSON(http.StatusOK, RespSuccess(c, recordId))

}

func (h *DNSHandler) Delete(c *gin.Context) {
	var req types.RecordIDGetReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, RespError(c, err, ""))

		return
	}
	if err := h.service.Delete(req.DomainName, req.RR); err != nil {
		c.JSON(http.StatusInternalServerError, RespError(c, err, "删除DNS记录失败"))

		return
	}
	c.JSON(http.StatusOK, RespSuccess(c, "删除成功"))

}
func (h *DNSHandler) Update(c *gin.Context) {
	// recordId := c.Param("id")
	var record types.DomainRecordUpdateReq
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, RespError(c, err, ""))
		return
	}
	// record.RecordId = recordId

	if err := h.service.Update(&record); err != nil {
		c.JSON(http.StatusInternalServerError, RespError(c, err, "更新DNS记录失败"))

		return
	}

	c.JSON(http.StatusOK, RespSuccess(c, ""))
}
func (h *DNSHandler) SetStatus(c *gin.Context) {
	var req types.DomainRecordStatusUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, RespError(c, err, ""))
		return
	}
	if err := h.service.SetStatus(&req); err != nil {
		c.JSON(http.StatusBadRequest, RespError(c, err, "更新DNS状态失败"))

		return
	}
	c.JSON(http.StatusOK, RespSuccess(c, "更新状态成功"))
}
func (h *DNSHandler) List(c *gin.Context) {
	domain := c.Query("domain")
	if domain == "" {
		c.JSON(http.StatusBadRequest, RespError(c, nil, "domain参数必填"))
		return
	}

	result, err := h.service.List(domain)
	if err != nil {
		c.JSON(http.StatusBadRequest, RespError(c, err, "查询DNS记录失败"))
		return
	}
	c.JSON(http.StatusOK, RespSuccess(c, result))
}

func (h *DNSHandler) QueryRecordID(c *gin.Context) {
	var req types.RecordIDGetReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, RespError(c, err, ""))
		return
	}

	recordId, err := h.service.QueryRecordID(req.DomainName, req.RR)
	if err != nil {
		c.JSON(http.StatusNotFound, RespError(c, err, "查询DNS记录不存在"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"domain":    req.DomainName,
		"rr":        req.RR,
		"record_id": recordId,
	})
}
