# aliops
### 查询
```
curl -XGET http://localhost:8800/api/dns/records?domain=gopron.cn
```
#### 添加
```
curl -XPOST http://localhost:8800/api/dns/records

{
    "domain": "gopron.cn",
    "rr": "test1",
    "type": "A",
    "value": "100.10.1.100",
    "ttl": 600
}
```

### 删除
```
curl -XDELETE http://localhost:8800/api/dns/records?domain=gopron.cn&rr=www

```

### 更新记录
```
curl -XPUT http://localhost:8800/api/dns/records
{
    "record_id": "1948944545861936128",
    "rr": "test4",
    "type": "A",
    "value": "100.10.1.101",
    "ttl": 600
}
```
### 更新记录状态
```
curl -XPATCH http://localhost:8800/api/dns/records
{
    "domain": "gopron.cn",
    "rr": "test4",
    "status": "enable"
}
```