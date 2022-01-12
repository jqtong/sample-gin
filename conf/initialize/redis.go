package initialize

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type redisCache struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database int    `mapstructure:"database"`
	Password string `mapstructure:"password"`
}

var redisConf = &redisCache{}

func InitRedis() *redis.Client {

	if err := viper.UnmarshalKey("redis.store_platform", redisConf); err != nil {
		log.Fatalf("Parse config.redis.bi segment error: %s", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.Database,
	})

	_, err := rdb.Ping().Result()

	if err != nil {
		log.Fatalf("Connect redis.bi error: %s", err)
	}

	return rdb
}