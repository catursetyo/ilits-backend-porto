package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("yahhhh... gagal connect ke database: %v", err)
	}

	log.Println("database SQLite udah connect!")
	return DB
}
