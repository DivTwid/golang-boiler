package main

import (
	"fmt"
	"log"

	"github.com/DivTwid/golang-boiler/config"
	"github.com/DivTwid/golang-boiler/db"
	"github.com/DivTwid/golang-boiler/db/migration"
	"github.com/DivTwid/golang-boiler/db/seeders"
	"github.com/DivTwid/golang-boiler/router"
)

func main() {
	// Load Config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config", err)
	}
	fmt.Println("Database Host:", config.DatabasePostgres.Host)
	//Initialize Postgres DB
	postgres := db.PostgresDB{
		Host:     config.DatabasePostgres.Host,
		User:     config.DatabasePostgres.Username,
		Password: config.DatabasePostgres.Password,
		Port:     config.DatabasePostgres.Port,
		DB:       config.DatabasePostgres.Name,
		Sslmode:  "disable",
		Timezone: "Asia/Shanghai",
	}
	postgres.Init()

	//Run Migrations
	migration := migration.Migration{
		DB:       db.PqDB,
		Migrator: db.PqDB.Migrator(),
	}
	migration.Migrate()
	// migration.AlterMigrate()
	// migration.Rollback()

	//Seed Data
	seed := seeders.Seeder{
		DB: db.PqDB,
	}
	seed.UserSeeder()

	//Initialize Mysql
	mysql := db.MySqlDB{
		Host:     config.DatabaseMysql.Host,
		Port:     config.DatabaseMysql.Port,
		User:     config.DatabaseMysql.Username,
		Password: config.DatabaseMysql.Password,
		DB:       config.DatabaseMysql.Name,
	}

	mysql.Init()

	//SetupRoutes
	routes := router.SetupRouter()
	routes.Run(":" + config.App.Port)

}
