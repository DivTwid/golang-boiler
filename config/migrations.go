package config

import (
	"fmt"

	"github.com/DivTwid/golang-boiler/model"
	"github.com/DivTwid/golang-boiler/seeders"

	"gorm.io/gorm"
)

type Migration interface {
	Migrate()
	Seed()
}
type migration struct {
	db *gorm.DB
}

func NewMigration(db *gorm.DB) Migration {
	return &migration{
		db: db,
	}
}

func (m *migration) Migrate() {
	m.db.AutoMigrate(&model.User{})
}

func (m *migration) Seed() {
	users := seeders.UserSeeder(PqDB)
	fmt.Println("users", users)
}
