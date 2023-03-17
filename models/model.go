package mod

import "gorm.io/gorm"

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
	User_Id   string `json:"user_id" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	User_Name string `json:"name"`
	Gender    string `json:"gender"`
}

type Payment struct {
	Payment_Id   string  `json:"payment_id" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	User_Id      string  `json:"user_id"`                                                        //FK
	User         User    `gorm:"references:User_Id"`
	Amount       float64 `json:"amount"`
	Payment_Type string  `json:"payment_type"`
}

type Subscription struct {
	Payment_Id string         `json:"payment_id" gorm:"default:null"` //FK
	Payment    Payment        `gorm:"references:Payment_Id"`
	User_Id    string         `json:"user_id"` //Fk
	User       User           `gorm:"references:User_Id"`
	Emp_Id     string         `json:"emp_id"`    //FKs
	Subs_Name  string         `json:"subs_name"` //FK
	StartDate  string         `json:"start_date"`
	EndDate    string         `json:"end_date"`
	Duration   float64        `json:"duration"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type SubsType struct {
	Subs_Name string  `json:"subs_name" gorm:"unique"`
	Price     float64 `json:"price"`
}

type GymEmp struct {
	Emp_Id   string `json:"emp_id" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	Emp_name string `json:"emp_name"`
	Gender   string `json:"gender"`
	Role     string `json:"role"`
}

type Equipment struct {
	Model_No   string `json:"model_no" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
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
type Slot struct {
	ID              int    `gorm:"primaryKey"`
	Start_time      string `json:"start_time"`
	End_time        string `json:"end_time"`
	Available_slots int64  `json:"available_slots" gorm:"default:50"`
}
type Credentials struct {
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
