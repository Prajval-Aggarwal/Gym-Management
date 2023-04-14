package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	fmt.Println("Initalising Database...")
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"),
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	Execute(db)
	AutoMigrateDatabase(db)
	return db

}
