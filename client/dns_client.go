package client

import (
	"fmt"

	dns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/tiamxu/aliops/config"
	"github.com/tiamxu/aliops/types"
)

type DNSClient struct {
	client *dns.Client
	config *config.AliyunConfig
}

func NewDNSClient(cfg *config.AliyunConfig) (*DNSClient, error) {
	openapiConfig := &openapi.Config{
		AccessKeyId:     tea.String(cfg.AccessKeyId),
		AccessKeySecret: tea.String(cfg.AccessKeySecret),
		RegionId:        tea.String(cfg.RegionId),
	}
	c, err := dns.NewClient(openapiConfig)
	if err != nil {
		return nil, fmt.Errorf("创建客户端失败: %w", err)
	}
	return &DNSClient{
		client: c,
		config: cfg,
	}, nil
}

func (c *DNSClient) AddDomainRecord(req *types.DomainRecordAddReq) (*dns.AddDomainRecordResponse, error) {
	resp, err := c.client.AddDomainRecord(&dns.AddDomainRecordRequest{
		DomainName: tea.String(req.DomainName),
		RR:         tea.String(req.RR),
		Type:       tea.String(req.Type),
		Value:      tea.String(req.Value),
		TTL:        tea.Int64(req.TTL),
	})
	return resp, err
}

func (c *DNSClient) DeleteDomainRecord(recordId *string) (*dns.DeleteDomainRecordResponse, error) {
	resp, err := c.client.DeleteDomainRecord(&dns.DeleteDomainRecordRequest{
		RecordId: recordId,
	})
	return resp, err
}
func (c *DNSClient) UpdateDomainRecord(req *types.DomainRecordUpdateReq) (*dns.UpdateDomainRecordResponse, error) {
	resp, err := c.client.UpdateDomainRecord(&dns.UpdateDomainRecordRequest{
		RecordId: &req.RecordId,
		RR:       &req.RR,
		Type:     &req.Type,
		Value:    &req.Value,
		TTL:      &req.TTL,
	})
	return resp, err
}
func (c *DNSClient) SetDomainRecordStatus(recordId, status *string) (*dns.SetDomainRecordStatusResponse, error) {
	resp, err := c.client.SetDomainRecordStatus(&dns.SetDomainRecordStatusRequest{
		RecordId: recordId,
		Status:   status,
	})
	return resp, err
}
func (c *DNSClient) DescribeAllRecords(domain string) ([]*dns.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	req := &dns.DescribeDomainRecordsRequest{
		DomainName: tea.String(domain),
		PageSize:   tea.Int64(100),
	}

	resp, err := c.client.DescribeDomainRecords(req)
	if err != nil {
		return nil, fmt.Errorf("查询DNS记录失败: %w", err)
	}

	if resp.Body == nil || resp.Body.DomainRecords == nil {
		return nil, fmt.Errorf("未获取到有效响应数据")
	}

	return resp.Body.DomainRecords.Record, nil
}

func (c *DNSClient) DescribeRecordsByType(domain, recordType string) ([]*dns.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	req := &dns.DescribeDomainRecordsRequest{
		DomainName: tea.String(domain),
		Type:       tea.String(recordType),
		PageSize:   tea.Int64(100),
	}

	resp, err := c.client.DescribeDomainRecords(req)
	if err != nil {
		return nil, fmt.Errorf("查询%s记录失败: %w", recordType, err)
	}

	return resp.Body.DomainRecords.Record, nil
}
