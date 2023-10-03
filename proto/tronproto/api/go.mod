module github.com/wangyuche/usdtmonitor/proto/tronproto/api

go 1.20

require (
	github.com/tronprotocol/grpc-gateway/core v1.3.1-0.20180628072903-5e70d2d524cf
	google.golang.org/genproto/googleapis/api v0.0.0-20230920204549-e6e6cdab5c13
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230913181813-007df8e322eb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230920183334-c177e329c48b // indirect
)

replace github.com/tronprotocol/grpc-gateway/core => ../core
