package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	Payment_Id   string         `json:"payment_id" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	User_Id      string         `json:"user_id"`                                                        //FK
	Amount       float64        `json:"amount"`
	OfferAmount  float64        `json:"offer_amount"`
	Offer        string         `json:"offer"`
	Payment_Type string         `json:"payment_type"`
	Status       string         `json:"status"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Order_id     string         `json:"order_id"`    
	Membership   string         `json:"membership"`   
}

type Subscription struct {
	Payment_Id string         `json:"payment_id" gorm:"default:null"` //FK
	User_Id    string         `json:"user_id"`                        //Fk
	Emp_Id     string         `json:"emp_id" gorm:"default:null"`     //FKs
	Emp_name   string         `json:"emp_name"`
	Subs_Name  string         `json:"subs_name"` //FK
	StartDate  string         `json:"start_date"`
	EndDate    string         `json:"end_date"`
	Duration   float64        `json:"duration"`
	Slot_id    int            `json:"slot_id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type Membership struct {
	MemName   string         `json:"membershipName" gorm:"unique" validate:"required"`
	Price     float64        `json:"price" validate:"required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
