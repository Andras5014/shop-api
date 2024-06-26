package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"shop_api/goods_web/global"
	"shop_api/goods_web/initialize"
	"shop_api/goods_web/utils"
	"shop_api/goods_web/utils/register/consul"
	"syscall"
)

func main() {

	// 初始化日志
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()

	// 初始化routers
	router := initialize.Routers()

	// 初始化翻译
	_ = initialize.InitTrans("zh")

	// 初始化srv连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	debug := viper.GetBool("SHOP_DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		if err != nil {
			zap.S().Panic("获取空闲端口失败", zap.Error(err))
		}
		global.ServerConfig.Port = port
	}

	// 服务注册
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.New())

	err := register_client.Register(global.ServerConfig.Name, serviceId, global.ServerConfig.Host, global.ServerConfig.Port,
		global.ServerConfig.Tags)
	if err != nil {
		zap.S().Panic("服务注册失败", zap.Error(err))
	}
	zap.S().Info("goods_web启动成功，端口：", global.ServerConfig.Port)

	go func() {
		if err := router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败", zap.Error(err))
		}

	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := register_client.DeRegister(serviceId); err != nil {
		zap.S().Panic("注销失败", zap.Error(err))
	}
	zap.S().Info("注销成功")

}
