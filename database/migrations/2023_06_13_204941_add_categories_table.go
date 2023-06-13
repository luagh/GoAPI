package migrations

import (
	"GOHUB/app/models"
	"GOHUB/pkg/migrate"
	"database/sql"
	"gorm.io/gorm"
)

func init() {
	type Category struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2023_06_13_204941_add_categories_table", up, down)
}
