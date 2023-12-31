package db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global Postgres DB instance
var PqDB *gorm.DB
var MysqlDB *gorm.DB
var Redis *redis.Client

type PostgresDB struct {
	Host     string
	User     string
	Password string
	DB       string
	Port     string
	Sslmode  string
	Timezone string
}

func (ps PostgresDB) Init() {
	dsn := "host=" + ps.Host + " user=" + ps.User + " password=" + ps.Password + " dbname=" + ps.DB + " port=" + ps.Port + " sslmode=" + ps.Sslmode + " TimeZone=" + ps.Timezone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect PostgreSql")
	}
	log.Println("Postgres connection success")
	PqDB = db
}

type MySqlDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func (ms MySqlDB) Init() {
	dsn := "" + ms.User + ":" + ms.Password + "@tcp(" + ms.Host + ":" + ms.Port + ")/" + ms.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect Mysql")
	}
	log.Print("Mysql Connection success")
	MysqlDB = db
}

type RedisClient struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func (rs RedisClient) Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rs.Host + ":" + rs.Port,
		Password: rs.Password,
		DB:       rs.DB,
	})
	Redis = rdb
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil || pong != "PONG" {
		panic("Redis connection failure")
	} else {
		fmt.Println("Redis connection success")
	}
}
