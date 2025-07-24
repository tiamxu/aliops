package types

import "github.com/tiamxu/aliops/model"

type UpdateStatusRequest struct {
	RecordId string `json:"record_id" binding:"required"`
	Status   string `json:"status" binding:"required,oneof=ENABLE DISABLE"`
}

type ListResponse struct {
	Records []model.DNSRecord `json:"records"`
	Total   int               `json:"total"`
}
