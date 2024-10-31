package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqlLiteConnection() (*gorm.DB, error) {

	dbfilePath := "teste.db"

	db, err := gorm.Open(sqlite.Open(dbfilePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
