package initialize

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // load mysql driver
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/url"
	"time"
)

type db struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

var dbConfig = &db{}

func InitDB() *gorm.DB {
	if err := viper.UnmarshalKey("mysql.store_platform", dbConfig); err != nil {
		log.Fatalf("Parse config.mysql.bi segment error: %s\n", err)
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
		url.QueryEscape(ServerConf.Timezone),
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Connect mysql.bi error: %s", err)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(time.Minute * 30)

	return db
}

func InitMongo() *mongo.Client {
	// mongo 192.168.1.150:20000/project_lt -u user_lt -p 'wode@lt'
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		viper.GetString("mongo.bi.user"),
		viper.GetString("mongo.bi.pass"),
		viper.GetString("mongo.bi.host"),
		viper.GetString("mongo.bi.port"),
	)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal("Connecting mongo error...")
	}

	return db
}
