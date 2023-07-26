package migration

import (
	"fmt"

	"github.com/DivTwid/golang-boiler/model"
	"gorm.io/gorm"
)

type Migration struct {
	DB       *gorm.DB
	Migrator gorm.Migrator
}

func (m Migration) Migrate() {
	m.DB.AutoMigrate(&model.User{})
}

func (m Migration) Rollback() {
	m.DB.Migrator().DropTable(&model.User{})
	fmt.Println("Rollback completed successfully.")
}

func (m Migration) AlterMigrate() {
	m.Migrator.RenameColumn(&model.User{}, "Fullname", "Name")
	m.Migrator.AddColumn(&model.User{}, "Age")
	fmt.Println("Alter migrations completed successfully.")
}
