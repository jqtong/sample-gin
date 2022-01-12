package initialize

import (
	"log"

	"github.com/spf13/viper"
)

type server struct {
	RunMode  string `mapstructure:"runMode"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Timezone string `mapstructure:"timezone"`
}

// ServerConf is config of http server
var ServerConf = &server{}

func InitServerConf() {

	if err := viper.UnmarshalKey("server", ServerConf); err != nil {
		log.Fatalf("Parse config.server segment error: %s", err)
	}
}
