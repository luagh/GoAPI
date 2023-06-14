package seeders

import (
	"GOHUB/database/factories"
	"GOHUB/pkg/console"
	"GOHUB/pkg/logger"
	"GOHUB/pkg/seed"
	"fmt"
	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedCategoriesTable", func(db *gorm.DB) {

		categories := factories.MakeCategories(10)

		result := db.Table("categories").Create(&categories)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
