module userwebPRJ

go 1.12

require (
	github.com/LiuBaiSMD/goProPlugins/user/proto v0.0.0-20190812061729-8d1f72939b39
	github.com/go-redis/redis v6.15.2+incompatible // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gorilla/websocket v1.4.0
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.9.0
	github.com/micro/micro v1.9.0 // indirect
	github.com/nats-io/nats-server/v2 v2.0.2 // indirect
)

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
