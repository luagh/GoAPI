// Package link 模型
package link

import (
	"GOHUB/app/models"
	"GOHUB/pkg/database"
)

type Link struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	URL  string `json:"URL,omitempty"`

	models.CommonTimestampsField
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}
