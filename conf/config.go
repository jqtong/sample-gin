package conf

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"nucarf.com/store_service/api/conf/initialize"
)

// BiMysql is mysql client
var Mysql *gorm.DB

//BiMongo mongo client
var BiMongo *mongo.Client

// BiRedis is redis client
var Redis *redis.Client

// GeneralLog debug info and so on...
var GeneralLog *logrus.Logger

// AppLog application panic or other error log
var AppLog *logrus.Logger

// AccessLog api calls
var AccessLog *logrus.Logger

//BiMongo mongo clent
var Mongo *mongo.Client

// Load load config and init mysql and redis client
func Load() {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config file error: %s", err)
	}

	initialize.InitServerConf()

	// init mysql and redis
	Mysql = initialize.InitDB()
	Redis = initialize.InitRedis()
	//Mongo = initialize.InitMongo()

	// init all the log targets
	AppLog = initApplicationLog()
}
