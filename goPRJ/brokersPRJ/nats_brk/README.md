## RUN nats_broker
#### 1 安装nats
```
go get github.com/nats-io/gnatsd
```
#### 2 运行nats
```
gnatsd
```
### 3 运行server.go client.go 
运行subscriber
```
go run server.go 
```
运行publish
```
go run client.go  #publish
```
