# 基于ws的消息发布订阅

### 连接websocket
ws://127.0.0.1:80000/ws

### 添加订阅
```
{
"action": "setChannel",
"data": {
 "channels": [ "频道编号"]
 }
}
```

### 发布消息
 - 广播消息给所有用户 
```
#数据结构

{
    "data":{"event":"paySuccess","orderId":"1"}
}

#地址 http://127.0.0.1:8000/broadcast
```
```
curl -i -X POST -H 'Content-Type: application/json; charset=UTF-8' -d '{
    "data":{"event":"paySuccess","orderId":"1"}
}' http://127.0.0.1:8000/broadcast
```
- 广播消息给特定频道
```
#数据结构
{
    "channels":["频道1","频道2"],
    "data":{"event":"paySuccess","orderId":"1"}
}
#地址 http://127.0.0.1:8000/broadcast-channel

```

```
curl -i -X POST -H 'Content-Type: application/json; charset=UTF-8' -d '{
    "channels":["12","34"],
    "data":{"event":"paySuccess","orderId":"1"}
}' http://127.0.0.1:8000/broadcast-channel
```


