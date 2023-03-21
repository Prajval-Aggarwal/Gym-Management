package mod

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// only display for user full information
type Display struct {
	User_Id      string         `json:"user_id"`
	User_Name    string         `json:"user_name"`
	Gender       string         `json:"gender"`
	Amount       float64        `json:"amount"`
	Payment_Type string         `json:"payment_type"`
	Payment_Id   string         `json:"payment_id" `
	Emp_Id       string         `json:"emp_id"`
	Subs_Name    string         `json:"subs_name"`
	StartDate    string         `json:"start_date"`
	EndDate      string         `json:"end_date"`
	Duration     float64        `json:"duration"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
type User struct {
	User_Id    string `json:"user_id" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	User_Name  string `json:"name"`
	Gender     string `json:"gender"`
	Contact_No string `json:"phoneNumber"`

}

type Payment struct {
	Payment_Id   string  `json:"payment_id" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	User_Id      string  `json:"user_id"`                                                        //FK
	User         User    `gorm:"references:User_Id"`
	Amount       float64   `json:"amount"`
	OfferAmount float64   `json:"offer_amount"`
	Offer          string `json:"offer"`
	Payment_Type string  `json:"payment_type"`

}

type Subscription struct {
	Payment_Id string         `json:"payment_id" gorm:"default:null"` //FK
	Payment    Payment        `gorm:"references:Payment_Id"`
	User_Id    string         `json:"user_id"` //Fk
	User       User           `gorm:"references:User_Id"`
	Emp_Id     string         `json:"emp_id" gorm:"default:null"` //FKs
	Emp        GymEmp         `gorm:"references:Emp_Id"`
	Emp_name   string         `json:"emp_name"`
	Subs_Name  string         `json:"subs_name"` //FK
	StartDate  string         `json:"start_date"`
	EndDate    string         `json:"end_date"`
	Duration   float64        `json:"duration"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Slot_id    int64          `json:"slot_id"`
}

type SubsType struct {
	Subs_Name string  `json:"subs_name" gorm:"unique"`
	Price     float64 `json:"price"`
}

type GymEmp struct {
	//gorm.Model
	Emp_Id    string    `json:"emp_id" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	Emp_name  string    `json:"emp_name"`
	Gender    string    `json:"gender"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"updated"`
}

// employee who has users under him count
type EmpWithUser struct {
	Emp_id          string `json:"emp_id"`
	Emp_name        string `json:"emp_name"`
	Alotted_members int    `json:"alotted_members"`
}

// employee types
type EmpTypes struct {
	Role   string  `json:"role" gorm:"unique"`
	Salary float64 `json:"salary"`
}

type Equipment struct {
	// Model_No   string`json:"model_no" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	Equip_Name string `json:"equip_name"`
	Quantity   int64  `json:"quantity"`
}

type UAttendence struct {
	User_Id string `json:"user_id"`
	Date    string `json:"date"`
	Present string `json:"present" gorm:"default:null"`
}
type EmpAttendence struct {
	User_Id string `json:"user_id"`
	Date    string `json:"date"`
	Present string `json:"present" gorm:"default:null"`
}

// slot structs
type Slot struct {
	ID              int    `json:"slot_id" gorm:"primaryKey"`
	Start_time      string `json:"start_time"`
	End_time        string `json:"end_time"`
	Available_space int64  `json:"available_slots" gorm:"default:50"`
}
type Credential struct {
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	//Contact  string `json:"contact" gorm:"unique"`
}
type DbVersion struct {
	Version int `json:"version"`
}
type Claims struct {
	Username string `json:"user_id"`
	jwt.RegisteredClaims
}

