package database

import (
	"fmt"
	"os"

	"canarails.dev/database/models"
	"canarails.dev/services/envsvc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	databaseUrl := os.Getenv(envsvc.DATABASE_URL)

	db, err := gorm.Open(
		postgres.Open(databaseUrl),
		&gorm.Config{},
	)
	if err != nil {
		panic(fmt.Errorf("open database error: %w", err))
	}

	err = db.AutoMigrate(models.Models...)
	if err != nil {
		panic(fmt.Errorf("auto migrate error: %w", err))
	}

	return db
}

var db *gorm.DB

func GetDb() *gorm.DB {
	if db == nil {
		db = initDb()
	}

	return db
}
