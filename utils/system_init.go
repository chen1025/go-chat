package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB          *gorm.DB
	Conf        *Config
	RedisClient *redis.Client
	SN          *Node
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		err := viper.Unmarshal(&Conf)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func InitSN() {
	node, err := NewNode(1)
	if err != nil {
		panic(interface{}("failed to create snowflake"))
	}
	SN = node
}

type Config struct {
	Mysql   MysqlConfig
	Redis   RedisConfig
	Timeout TimeOutConfig
	Log     LogConfig
}

type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host        string
	Port        string
	Password    string
	DB          int
	PoolSize    int
	MinIdleConn int
}

type TimeOutConfig struct {
	DelayHeartbeat   int
	HeartbeatHz      int
	HeartbeatMaxTime int
	RedisOnlineTime  int
}

type LogConfig struct {
	Level string
	Path  string
}

func InitMysql() {
	// 自定义日志，打印日志
	mysqlLog := logger.New(
		log.New(os.Stdin, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢sql
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		})
	my := Conf.Mysql
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", my.Username, my.Password, my.Host, my.Port, my.Database)
	client, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: mysqlLog})
	if err != nil {
		panic(interface{}("failed to connect database"))
	}
	fmt.Println("successfully initialized mysql")
	sqlDB, err := client.DB()
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	DB = client
}

func InitRedis() {
	re := Conf.Redis
	// 自定义日志，打印日志
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         re.Host + ":" + re.Port,
		Password:     re.Password, // no password set
		DB:           re.DB,       // use default DB
		MinIdleConns: re.MinIdleConn,
		PoolSize:     re.PoolSize,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(interface{}("failed to connect redis"))
	} else {
		fmt.Println("successfully initialized redis")
	}

}
