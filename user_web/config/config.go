package config

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
type AliSmsConfig struct {
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret"`
}
type RedisConfig struct {
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	Expiration int    `mapstructure:"expiration" json:"expiration"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
type JaegerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
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
	Name           string           `mapstructure:"name" json:"name"`
	Port           int              `mapstructure:"port" json:"port"`
	Host           string           `mapstructure:"host" json:"host"`
	Tags           []string         `mapstructure:"tags" json:"tags"`
	JWTInfo        JWTConfig        `mapstructure:"jwt" json:"jwt"`
	UserServerInfo UserServerConfig `mapstructure:"user_srv" json:"user_srv"`
	AliSmsInfo     AliSmsConfig     `mapstructure:"ali_sms" json:"ali_sms"`
	RedisInfo      RedisConfig      `mapstructure:"redis" json:"redis"`
	ConsulInfo     ConsulConfig     `mapstructure:"consul" json:"consul"`
	JaegerInfo     JaegerConfig     `mapstructure:"jaeger" json:"jaeger"`
}
