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
	// set defaul config
	c.defaultConfig()

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
	c.Log.Depth = 8
	c.Log.Level = "info"
	c.Log.Write = false
	c.Log.MaxAge = 24 * 7   // 7 days
	c.Log.RotationTime = 24 // 24 hours
}

// TLSOptions web socket options
type TLSOptions struct {
	Addr     string
	CertFile string
	KeyFile  string
	Disable  bool
}

// CommonOptions common options
type CommonOptions struct {
	TempFolder string // temp file dir
}

// LogOptions log options
type LogOptions struct {
	Level string
	Depth int

	LogFilePrefix  string
	LogFileName    string
	LogDir         string
	DisableConsole bool
	Write          bool
	WithCallerHook bool

	MaxAge       int // rotatelogs max age, unit hour
	RotationTime int // rotatelogs rotation time, unit hour
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
