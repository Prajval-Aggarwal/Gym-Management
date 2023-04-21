package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	AdminId   string         `json:"adminId"  gorm:"default:uuid_generate_v4();unique;primaryKey"`
	Name      string         `json:"name" gorm:"unique"`
	Role      string         `json:"role"`
	Contact   string         `json:"contact" gorm:"unique"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
