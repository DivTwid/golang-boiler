package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PqDB *gorm.DB

type PostgresDB struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
	timezone string
}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_DATABASE"),
		port:     os.Getenv("DB_PORT"),
		sslmode:  "disable",
		timezone: "Asia/Shanghai",
	}
}

func (ps PostgresDB) InitializeDB() {
	dsn := "host=" + ps.host + " user=" + ps.user + " password=" + ps.password + " dbname=" + ps.dbname + " port=" + ps.port + " sslmode=" + ps.sslmode + " TimeZone=" + ps.timezone
	fmt.Println("dsn", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	PqDB = db
}
