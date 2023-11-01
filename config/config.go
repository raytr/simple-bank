package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBConfig  DBConfig       `mapstructure:"database"`
	Server    ServerConfig   `mapstructure:"server"`
	SecConfig SecurityConfig `mapstructure:"security"`
}

type DBConfig struct {
	Host                  string        `mapstructure:"host"`
	Port                  int           `mapstructure:"port"`
	Username              string        `mapstructure:"username"`
	Password              string        `mapstructure:"password"`
	DBName                string        `mapstructure:"dbname"`
	SchemaName            string        `mapstructure:"schemaname"`
	MaxIdleConnection     int           `mapstructure:"max-idle-connections" default:"20"`
	MaxOpenConnection     int           `mapstructure:"max-open-connections" default:"100"`
	ConnectionMaxLifeTime time.Duration `mapstructure:"connection-max-lifetime" default:"1200"`
	ConnectionMaxIdleTime time.Duration `mapstructure:"connection-max-idletime" default:"1"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type SecurityConfig struct {
	AccessTokenDuration     time.Duration `mapstructure:"access-token-duration"`
	RefreshTokenDuration    time.Duration `mapstructure:"refresh-token-duration"`
	Type                    string        `mapstructure:"type"`
	SecurityToken           string        `mapstructure:"security-token"`
	PasetoTokenSymmetricKey string        `mapstructure:"paseto-token-symmetric-key"`
	PasswordPepper          string        `mapstructure:"pepper"`
	PasswordSaltLength      int           `mapstructure:"salt-length"`
	AuthSkip                bool          `mapstructure:"auth-skip"`
}

func Init(configName, configType string) *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}
