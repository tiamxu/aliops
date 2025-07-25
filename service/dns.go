package service

import (
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/tiamxu/aliops/client"
	"github.com/tiamxu/aliops/model"
	"github.com/tiamxu/aliops/types"
)

type DNSService struct {
	client *client.DNSClient
}

func NewDNSService(client *client.DNSClient) *DNSService {
	return &DNSService{client: client}
}

func (s *DNSService) ListAllRecords(domain string) error {
	records, err := s.client.DescribeAllRecords(domain)
	if err != nil {
		return fmt.Errorf("获取DNS记录失败: %w", err)
	}

	fmt.Printf("域名 %s 的所有解析记录:\n", domain)
	for i, record := range records {
		fmt.Printf("%d. 记录ID: %s, 主机记录: %s, 类型: %s, 值: %s, TTL: %d, 状态: %s\n",
			i+1,
			tea.StringValue(record.RecordId),
			tea.StringValue(record.RR),
			tea.StringValue(record.Type),
			tea.StringValue(record.Value),
			tea.Int64Value(record.TTL),
			tea.StringValue(record.Status))
	}
	fmt.Println(records)
	return nil
}

func (s *DNSService) ListRecordsByType(domain, recordType string) error {
	records, err := s.client.DescribeRecordsByType(domain, recordType)
	if err != nil {
		return fmt.Errorf("获取%s记录失败: %w", recordType, err)
	}

	fmt.Printf("域名 %s 的%s记录:\n", domain, recordType)
	for i, record := range records {
		fmt.Printf("%d. %s -> %s (TTL: %d)\n",
			i+1,
			tea.StringValue(record.RR),
			tea.StringValue(record.Value),
			tea.Int64Value(record.TTL))
	}

	return nil
}
func (s *DNSService) Add(req *types.DomainRecordAddReq) (string, error) {
	resp, err := s.client.AddDomainRecord(req)
	if err != nil {
		return "", fmt.Errorf("添加dns解析记录失败", err)
	}
	return tea.StringValue(resp.Body.RecordId), nil
}
func (s *DNSService) Delete(recordId string) error {
	return nil
}
func (s *DNSService) Update(record model.DNSRecord) error {
	return nil
}
func (s *DNSService) List(domain string) (types.ListResponse, error) {
	records, err := s.client.DescribeAllRecords(domain)
	if err != nil {
		return types.ListResponse{}, err
	}
	var result []model.DNSRecord
	for _, r := range records {
		result = append(result, model.DNSRecord{
			RecordId:   tea.StringValue(r.RecordId),
			DomainName: domain,
			RR:         tea.StringValue(r.RR),
			Type:       tea.StringValue(r.Type),
			Value:      tea.StringValue(r.Value),
			TTL:        tea.Int64Value(r.TTL),
			Status:     tea.StringValue(r.Status),
		})
	}

	return types.ListResponse{
		Records: result,
		Total:   len(result),
	}, nil
}
