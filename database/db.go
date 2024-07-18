package database

import (
	"fmt"
	"os"

	"canarails.dev/database/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	godotenv.Load("../.env")
	databaseUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(
		postgres.Open(databaseUrl),
		&gorm.Config{},
	)
	if err != nil {
		panic(fmt.Errorf("open database error: %w", err))
	}

	err = db.AutoMigrate(
		&models.App{},
		&models.AppVariant{},
		&models.AppDeploy{},
	)
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
