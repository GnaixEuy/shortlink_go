package repo

import (
	"log"
	"shortLink/internal/config"
	"shortLink/internal/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() error {
	dsn := config.C.MySQL.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if config.C.MySQL.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(config.C.MySQL.MaxOpenConns)
	}
	if config.C.MySQL.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.C.MySQL.MaxIdleConns)
	}
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	// 自动建表
	if err := db.AutoMigrate(&model.ShortLink{}); err != nil {
		return err
	}
	DB = db
	log.Println("[mysql] connected and migrated")
	return nil
}
