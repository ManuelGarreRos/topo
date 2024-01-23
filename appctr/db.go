package appctr

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func DB() *gorm.DB {
	return &db
}

var db gorm.DB

func prepareDB() {
	str := cfg.GetString("db")
	log.Println("Try to connect DB: ", str)

	d, err := gorm.Open(postgres.Open(str), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error to open db: %v", err)
	}
	Log().Debug("DB is Ok")

	db = *d
}
