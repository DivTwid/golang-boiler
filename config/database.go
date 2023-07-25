package config

import (
	"fmt"

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
		// host: "localhost",
		host:     "host.docker.internal", //use while building docker image
		user:     "root",
		password: "root",
		dbname:   "test",
		port:     "5432",
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
