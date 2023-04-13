package db

import (
	"fmt"
	mod "gym-api/Models"
	cons "gym-api/Utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func VersionManger() {
	var dbVersion mod.DbVersion
	err := DB.First(&dbVersion).Error
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("DB version is:", dbVersion.Version)
	if dbVersion.Version < 1 {
		err := DB.AutoMigrate(&mod.User{}, &mod.GymEmp{}, &mod.Payment{}, &mod.Subscription{}, &mod.SubsType{}, &mod.Equipment{}, &mod.UAttendence{}, &mod.EmpAttendence{}, &mod.EmpTypes{})
		if err != nil {
			panic(err)
		}
		DB.Create(&mod.DbVersion{
			Version: 1,
		})
		dbVersion.Version = 1
	}
	if dbVersion.Version < 2 {
		err := DB.AutoMigrate(&mod.Slot{})
		if err != nil {
			panic(err)
		}
		DB.Where("version=?", dbVersion.Version).Updates(&mod.DbVersion{
			Version: 2,
		})
		dbVersion.Version = 2

	}
	if dbVersion.Version < 3 {
		err := DB.AutoMigrate(&mod.Credential{})
		if err != nil {
			panic(err)
		}
		DB.Where("version=?", dbVersion.Version).Updates(&mod.DbVersion{
			Version: 3,
		})
		dbVersion.Version = 3
	}
	if dbVersion.Version < 4 {
		fmt.Println("ajkas", DB.Migrator().HasTable(mod.Equipment{}))
		if DB.Migrator().HasTable(mod.Equipment{}) {
			err := DB.Migrator().AddColumn(&mod.Equipment{}, "Weight")
			if err != nil {
				fmt.Println("eroor is ", err)
			}
		}
		DB.Where("version=?", dbVersion.Version).Updates(&mod.DbVersion{
			Version: 4,
		})
		dbVersion.Version = 4
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
	db.Exec("CREATE SCHEMA IF NOT EXISTS public")
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	db.AutoMigrate(&mod.DbVersion{})

	DB = db
	VersionManger()
	fmt.Println("Successfully connected to database")
	return nil
}
