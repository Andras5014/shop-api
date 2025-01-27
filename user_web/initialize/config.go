package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop_api/user_web/config"
	"shop_api/user_web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)

}

func InitConfig() {
	debug := GetEnvInfo("SHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user_web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user_web/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	global.NacosConfig = &config.NacosConfig{}
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	fmt.Println("配置信息：", global.NacosConfig)

	zap.S().Infof("配置信息：%v", global.NacosConfig)

	// 从nacos 获取配置信息
	InitNacosConfig()

}

func InitNacosConfig() {
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   uint64(global.NacosConfig.Port),
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		zap.S().Fatal("nacos 配置中心初始化失败")
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})
	fmt.Println("配置信息：", content)
	if err != nil {
		zap.S().Fatal("nacos 配置中心初始化失败")
		panic(err)
	}
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatal("读取nacos 配置中心初始化失败")
		return
	}
	fmt.Println(global.ServerConfig)
}
