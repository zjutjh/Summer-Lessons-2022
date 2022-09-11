package config

import (
	"log"

	"github.com/spf13/viper"
)

type server struct {
	Port         string   `mapstructure:"port"`
	AllowOrigins []string `mapstructure:"allow_origins"`
}
type db struct {
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	DBName   string `mapstructure:"db_name"`
}
type jwt struct {
	Secret  string `mapstructure:"secret"`
	Expires uint   `mapstructure:"expires"`
	Issuer  string `mapstructure:"issuer"`
}
type config struct {
	Dev    bool   `mapstructure:"dev"`
	Server server `mapstructure:"server"`
	DB     db     `mapstructure:"db"`
	Jwt    jwt    `mapstructure:"jwt"`
}

var Config config

func InitConfig() {
	var config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")
	err := config.ReadInConfig()
	if err != nil {
		log.Panicln("Config Error: ", err)
	}
	config.Unmarshal(&Config)
}
