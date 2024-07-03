package config

type GoodsServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
type UserOpServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
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
	Name          string             `mapstructure:"name" json:"name"`
	Port          int                `mapstructure:"port" json:"port"`
	Host          string             `mapstructure:"host" json:"host"`
	Tags          []string           `mapstructure:"tags" json:"tags"`
	JWTInfo       JWTConfig          `mapstructure:"jwt" json:"jwt"`
	GoodsSrvInfo  GoodsServerConfig  `mapstructure:"goods_srv" json:"goods_srv"`
	UserOpSrvInfo UserOpServerConfig `mapstructure:"userop_srv" json:"userop_srv"`
	ConsulInfo    ConsulConfig       `mapstructure:"consul" json:"consul"`
}
