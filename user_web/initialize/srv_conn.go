package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shop_api/user_web/global"
	"shop_api/user_web/proto"
)

func InitSrvConn() {

	//连接user grpc 服务
	consulInfo := global.ServerConfig.ConsulInfo

	userConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserServerInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接user服务失败", "msg", err.Error())
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
func InitSrvConn2() {
	// 从注册中心获取user grpc服务地址
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	var data map[string]*api.AgentService
	data, err = client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserServerInfo.Name))
	//data, err = client.Agent().ServicesWithFilter(fmt.Sprintf(`Service =="user-srv"`))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Errorw("[GetUserList] [找不到/连接user服务失败]")
		return
	}
	//连接user grpc 服务
	var userConn *grpc.ClientConn
	userConn, err = grpc.NewClient(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接user服务失败", "msg", err.Error())
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient

}
