package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shop_api/goods_web/global"
	"shop_api/goods_web/proto"
	"shop_api/goods_web/utils/otgrpc"
)

func InitSrvConn() {

	//连接goods grpc 服务
	consulInfo := global.ServerConfig.ConsulInfo

	goodsConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsServerInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接goods服务失败", "msg", err.Error())
	}

	inventoryConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.InventoryInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接inventory服务失败", "msg", err.Error())
	}
	inventorySrvClient := proto.NewInventoryClient(inventoryConn)
	global.InventorySrvClient = inventorySrvClient
	goodsSrvClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsSrvClient
}
