package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shop_api/goods_web/global"
	"shop_api/goods_web/proto"
)

func InitSrvConn() {

	//连接goods grpc 服务
	consulInfo := global.ServerConfig.ConsulInfo

	goodsConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsServerInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接user服务失败", "msg", err.Error())
	}

	goodsSrvClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsSrvClient
}
