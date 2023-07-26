package main

import (
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

	//Initialize Postgres DB
	postgres := db.PostgresDB{
		Host:     config.Database.Host,
		User:     config.Database.Username,
		Password: config.Database.Password,
		Port:     config.Database.Port,
		DB:       config.Database.Name,
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
	migration.AlterMigrate()
	// migration.Rollback()

	//Seed Data
	seed := seeders.Seeder{
		DB: db.PqDB,
	}
	seed.UserSeeder()

	//SetupRoutes
	routes := router.SetupRouter()
	routes.Run(":" + config.App.Port)
}
