package model

import (
	"time"

	"gorm.io/gorm"
)

type GymEmp struct {
	//gorm.Model
	Emp_Id    string         `json:"empId" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	Emp_name  string         `json:"empName"`
	Gender    string         `json:"gender"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

// employee who has users under him count
type EmpWithUser struct {
	Emp_id          string         `json:"empId"`
	Emp_name        string         `json:"empName"`
	Alotted_members int            `json:"alottedMembers"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

// employee types
type EmpTypes struct {
	Role      string         `json:"role" gorm:"unique"`
	Salary    float64        `json:"salary"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type EmpAttendence struct {
	User_Id   string         `json:"userId"`
	Date      string         `json:"date"`
	Present   string         `json:"present" gorm:"default:null"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
