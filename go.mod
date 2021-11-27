module pay_center

go 1.15

require (
	github.com/douyu/jupiter v0.2.7
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.7.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/wechatpay-apiv3/wechatpay-go v0.2.5
	go.uber.org/zap v1.17.0 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.38.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
