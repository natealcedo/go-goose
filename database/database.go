package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) Close() {
	sqlDb, err := db.DB.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDb.Close()
	if err != nil {
		panic(err)
	}
}

func CreateDatabaseClient() (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{DB: gormDb}, nil
}
