package db

import (
	"fmt"

	"gym/server/model"

	"gorm.io/gorm"
)

func AutoMigrateDatabase(db *gorm.DB) {

	var dbVersion model.DbVersion
	err := db.First(&dbVersion).Error
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("db version is:", dbVersion.Version)
	if dbVersion.Version < 1 {
		err := db.AutoMigrate(&model.User{}, &model.GymEmp{}, &model.Payment{}, &model.Subscription{}, &model.Membership{}, &model.Equipment{}, &model.UAttendence{}, &model.EmpAttendence{}, &model.EmpTypes{})
		if err != nil {
			panic(err)
		}
		db.Create(&model.DbVersion{
			Version: 1,
		})
		dbVersion.Version = 1
	}
	if dbVersion.Version < 2 {
		err := db.AutoMigrate(&model.Slot{})

		if err != nil {
			panic(err)
		}
		db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
			Version: 2,
		})
		dbVersion.Version = 2

	}
	if dbVersion.Version < 3 {
		err := db.AutoMigrate(&model.Credential{})
		if err != nil {
			panic(err)
		}
		db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
			Version: 3,
		})
		dbVersion.Version = 3
	}

}
