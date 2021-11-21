module pay_center

go 1.15

require (
	github.com/douyu/jupiter v0.2.7
	github.com/gin-gonic/gin v1.7.0
	github.com/golang/protobuf v1.5.2
	github.com/shopspring/decimal v1.2.0
	github.com/wechatpay-apiv3/wechatpay-go v0.2.5
	go.etcd.io/etcd/client/v3 v3.5.1
	google.golang.org/grpc v1.38.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
