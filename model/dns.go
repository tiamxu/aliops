package model

type DNSRecord struct {
	DomainName string `json:"domain_name"`
	RecordId   string `json:"record_id"`
	RR         string `json:"rr"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	TTL        int64  `json:"ttl"`
	Status     string `json:"status"`
	Locked     bool   `json:"locked"`
}
