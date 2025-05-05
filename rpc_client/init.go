package shorturlservice

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

var ShortUrlClient ShortUrlService

func InitGrpcClient() {

	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{conf.RpcHost}, // 你的 gRPC 服务地址
		NonBlock:  true,
	})

	ShortUrlClient = NewShortUrlService(cli)

}
