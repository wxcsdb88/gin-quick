package config

import (
	"os"

	"github.com/spf13/viper"
)

// LoadConfig load config, support json, toml, yaml, yml, properties, props, prop, hcl
// default is toml and json, if use other please add the related tag
func LoadConfig(file string) *GlobalConfig {
	if "" == file {
		dir, _ := os.Getwd()
		file = dir + "/config/app.toml"
	}

	v := viper.New()
	v.SetConfigFile(file) // auto detect the file suffix

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	c := &GlobalConfig{}

	err = v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	return c
}

// GlobalConfig global config
type GlobalConfig struct {
	Name string

	Mysql     MysqlOptions
	Redis     RedisOptions
	Jsonrpc   JsonrpcOptions
	Websocket WebsocketOptions
	Common    CommonOptions
	Log       LogOptions
	TLS       TLSOptions
}

// JsonrpcOptions json rpc options
type JsonrpcOptions struct {
	Port string
}

// WebsocketOptions web socket options
type WebsocketOptions struct {
	Port string
}

func (c *GlobalConfig) defaultConfig() {

}

// TLSOptions web socket options
type TLSOptions struct {
	Addr     string
	CertFile string
	KeyFile  string
}

// CommonOptions common options
type CommonOptions struct {
}

// LogOptions log options
type LogOptions struct {
}

// MysqlOptions mysql options
type MysqlOptions struct {
	Hostname           string
	Port               string
	User               string
	Password           string
	DBName             string
	TablePrefix        string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    int
	Debug              bool
}

// RedisOptions redis options
type RedisOptions struct {
	Host        string
	Port        string
	Password    string
	IdleTimeout int
	MaxIdle     int
	MaxActive   int
}
