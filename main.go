package main

import (
	"log"

	"github.com/DivTwid/golang-boiler/config"
	"github.com/DivTwid/golang-boiler/cron"
	"github.com/DivTwid/golang-boiler/db"
	"github.com/DivTwid/golang-boiler/db/migration"
	"github.com/DivTwid/golang-boiler/logs"
	"github.com/DivTwid/golang-boiler/router"
)

func main() {
	// Load Config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config", err)
	}
	//Initialize Postgres DB
	InitializePostgres(*config)
	//Initialize Mysql DB
	InitializeMysql(*config)
	//Initialize Redis
	InitializeRedis(*config)
	//Initialize Logs
	logs.InitLogger()
	//Cron
	cron.Cron()

	//SetupRoutes
	routes := router.SetupRouter()
	routes.Run(":" + config.App.Port)

}

func InitializePostgres(config config.Config) {
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
	// seed := seeders.Seeder{
	// 	DB: db.PqDB,
	// }
	// seed.UserSeeder()
}

func InitializeMysql(config config.Config) {
	//Initialize Mysql
	mysql := db.MySqlDB{
		Host:     config.DatabaseMysql.Host,
		Port:     config.DatabaseMysql.Port,
		User:     config.DatabaseMysql.Username,
		Password: config.DatabaseMysql.Password,
		DB:       config.DatabaseMysql.Name,
	}

	mysql.Init()

	// migration := migration.Migration{
	// 	DB:       db.MysqlDB,
	// 	Migrator: db.MysqlDB.Migrator(),
	// }
	// migration.Migrate()

	//Seed Data
	// seed := seeders.Seeder{
	// 	DB: db.MysqlDB,
	// }
	// seed.UserSeeder()
}

func InitializeRedis(config config.Config) {
	redis := db.RedisClient{
		Host:     config.Redis.Host,
		Password: config.Redis.Password,
		Port:     config.Redis.Port,
		DB:       config.Redis.DB,
	}
	redis.Init()
}
