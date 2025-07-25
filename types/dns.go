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

type DomainRecordAddReq struct {
	DomainName string `json:"domain" form:"domain"`
	RR         string `json:"rr" form:"rr"`
	Type       string `json:"type" form:"type"`
	Value      string `json:"value" form:"value"`
	TTL        int64  `json:"ttl" form:"ttl"`
}

type DomainRecordDelReq struct {
}
