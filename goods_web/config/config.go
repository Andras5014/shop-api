package config

type GoodsServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type JaegerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"data_id" json:"data_id"`
	Group     string `mapstructure:"group" json:"group"`
}
type ServerConfig struct {
	Name            string            `mapstructure:"name" json:"name"`
	Host            string            `mapstructure:"host" json:"host"`
	Port            int               `mapstructure:"port" json:"port"`
	Tags            []string          `mapstructure:"tags" json:"tags"`
	JWTInfo         JWTConfig         `mapstructure:"jwt" json:"jwt"`
	GoodsServerInfo GoodsServerConfig `mapstructure:"goods_srv" json:"goods_srv"`
	ConsulInfo      ConsulConfig      `mapstructure:"consul" json:"consul"`
	JaegerInfo      JaegerConfig      `mapstructure:"jaeger" json:"jaeger"`
}
