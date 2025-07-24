package client

import (
	"fmt"

	dns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/tiamxu/aliops/config"
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
