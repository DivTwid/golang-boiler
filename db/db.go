package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Global Postgres DB instance
var PqDB *gorm.DB

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
		panic("failed to connect database")
	}
	log.Println("Postgres connection success")
	PqDB = db
}
