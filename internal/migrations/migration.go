package migrations

import (
	"TOPO/appctr"
	"TOPO/internal/models"
)

func MigrateSchema() {
	db := appctr.DB()
	err := db.AutoMigrate(&models.UserEntity{})
	if err != nil {
		return
	}
}
