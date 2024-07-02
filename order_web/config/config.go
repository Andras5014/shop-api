package config

type OrderServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
type GoodsServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
type InventoryServerConfig struct {
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
type AlipayConfig struct {
	AppID        string `mapstructure:"app_id" json:"app_id"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnURL    string `mapstructure:"return_url" json:"return_url"`
}
type ServerConfig struct {
	Name                string                `mapstructure:"name" json:"name"`
	Host                string                `mapstructure:"host" json:"host"`
	Port                int                   `mapstructure:"port" json:"port"`
	Tags                []string              `mapstructure:"tags" json:"tags"`
	JWTInfo             JWTConfig             `mapstructure:"jwt" json:"jwt"`
	OrderServerInfo     OrderServerConfig     `mapstructure:"order_srv" json:"order_srv"`
	GoodsServerInfo     GoodsServerConfig     `mapstructure:"goods_srv" json:"goods_srv"`
	InventoryServerInfo InventoryServerConfig `mapstructure:"inventory_srv" json:"inventory_srv"`
	ConsulInfo          ConsulConfig          `mapstructure:"consul" json:"consul"`
	AliPayInfo          AlipayConfig          `mapstructure:"alipay" json:"alipay"`
	JaegerInfo          JaegerConfig          `mapstructure:"jaeger" json:"jaeger"`
}
