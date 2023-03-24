package db

import (
	"fmt"
	mod "gym-api/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Fun() {
	var dbVersion mod.DbVersion
	DB.First(&dbVersion)

	if dbVersion.Version < 1 {
		err := DB.AutoMigrate(&mod.User{}, &mod.GymEmp{}, &mod.Payment{}, &mod.Subscription{}, &mod.SubsType{}, &mod.Equipment{}, &mod.UAttendence{}, &mod.EmpAttendence{}, &mod.EmpTypes{})
		if err != nil {
			panic(err)
		}
		DB.Create(&mod.DbVersion{
			Version: 1,
		})

	}
	if dbVersion.Version < 2 {
		err := DB.AutoMigrate(&mod.Slot{})
		if err != nil {
			panic(err)
		}
		DB.Where("version=?", dbVersion.Version).Updates(&mod.DbVersion{
			Version: 2,
		})

	} 
	if dbVersion.Version < 3 {
		err := DB.AutoMigrate(&mod.Credential{})
		if err != nil {
			panic(err)
		}
		DB.Where("version=?", dbVersion.Version).Updates(&mod.DbVersion{
			Version: 3,
		})
	} else {
		fmt.Println("Database is currently in its latest version")
	}

}
func Connect() error {
	fmt.Println("Connecting to database...")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", os.Getenv("DB_Host"),os.Getenv("DB_Port"), os.Getenv("DB_User"), os.Getenv("DB_Password"), os.Getenv("Dbname"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database:", err)
		return err
	}
	db.Exec("CREATE SCHEMA IF NOT EXISTS public")
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	db.AutoMigrate(&mod.DbVersion{})

	DB = db
	Fun()
	fmt.Println("Successfully connected to database")
	return nil
}
