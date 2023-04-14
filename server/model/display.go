package model

import (
	"gorm.io/gorm"
)

type Display struct {
	User_Id      string         `json:"userId"`
	User_Name    string         `json:"userName"`
	Gender       string         `json:"gender"`
	Amount       float64        `json:"amount"`
	OfferAmount  float64        `json:"offerAmount"`
	Payment_Type string         `json:"paymentType"`
	Payment_Id   string         `json:"paymentId" `
	Emp_Id       string         `json:"empId"`
	Subs_Name    string         `json:"subsName"`
	StartDate    string         `json:"startDate"`
	EndDate      string         `json:"endDate"`
	Duration     float64        `json:"duration"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
