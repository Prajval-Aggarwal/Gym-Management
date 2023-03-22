package db

import (
	"fmt"
	mod "gym-api/Models"
	cons "gym-api/Utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Fun() {
	var dbVersion mod.DbVersion
	DB.First(&dbVersion)

	if dbVersion.Version < 1 {
		DB.AutoMigrate(&mod.Subscription{}, &mod.Payment{}, &mod.SubsType{}, &mod.User{}, &mod.GymEmp{}, &mod.Equipment{}, &mod.UAttendence{}, &mod.EmpAttendence{}, &mod.EmpTypes{})
		DB.Create(&mod.DbVersion{
			Version: 1,
		})
		//DB.Create(&dbVersion)
	} else if dbVersion.Version < 2 {
		DB.AutoMigrate(&mod.Slot{})
		DB.Where("version=?", dbVersion.Version).Updates(&mod.DbVersion{
			Version: 2,
		})
		//DB.Create(&dbVersion)
	} else if dbVersion.Version < 3 {
		DB.AutoMigrate(&mod.Credential{})
		DB.Where("version=?", dbVersion.Version).Updates(&mod.DbVersion{
			Version: 3,
		})
	} else {
		fmt.Println("Databse is currently in its latest version")
	}

}
func Connect() error {
	fmt.Println("Connecting to database...")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", cons.Host, cons.Port, cons.User, cons.Password, cons.Dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database:", err)
		return err
	}
	db.AutoMigrate(&mod.DbVersion{})

	DB = db
	Fun()
	fmt.Println("Successfully connected to database")
	return nil
}
