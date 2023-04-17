package model

import (
	"time"

	"gorm.io/gorm"
)

type Credential struct {
	UserID    string         `json:"userId"  gorm:"default:uuid_generate_v4();unique;primaryKey"`
	UserName  string         `json:"username" gorm:"unique"`
	Role      string         `json:"role"`
	Contact   string         `json:"contact" gorm:"unique"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
