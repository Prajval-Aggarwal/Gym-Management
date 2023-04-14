package db

import (
	"gym/server/model"

	"gorm.io/gorm"
)

func Execute(db *gorm.DB) {
	db.Exec("CREATE SCHEMA IF NOT EXISTS public")
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	db.AutoMigrate(&model.DbVersion{})
}
