package database

import (
	"backend/infrastructure/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func New(cfg config.Config) (Connection, error) {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.DBUser, cfg.Database.DBPass, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot open MYSQL")
	}
	log.Println("Databases Connected...")

	return Connection{
		DB: db,
	}, nil
}
