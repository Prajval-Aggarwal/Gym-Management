package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	User_Id    string         `json:"user_id" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	User_Name  string         `json:"name"`
	Gender     string         `json:"gender"`
	Contact_No string         `json:"phoneNumber"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type UAttendence struct {
	User_Id string `json:"user_id"`
	Date    string `json:"date"`
	Present string `json:"present" gorm:"default:null"`
}
