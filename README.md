# aliops

### 添加接口
```
JSON请求
http://localhost:8800/api/dns/records

Body:
{
    "domain": "gopron.cn",
    "rr": "test1",
    "type": "A",
    "value": "100.10.1.100",
    "ttl": 600
}
```