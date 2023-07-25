package main

import (
	"os"

	"github.com/DivTwid/golang-boiler/config"
	"github.com/DivTwid/golang-boiler/router"
)

const (
	path = "env/.env_test"
)

func init() {
	config.LoadENV(path)
	ps := config.NewPostgresDB()
	ps.InitializeDB()
	migrate := config.NewMigration(config.PqDB)
	migrate.Migrate()
	// migrate.Seed()
}

func main() {
	routes := router.SetupRouter()
	routes.Run(":" + os.Getenv("PORT"))
}
