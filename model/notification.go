package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Notification struct{
	gorm.Model
	Type string		`json:"type"`
	Dibaca bool		`json:"dibaca" gorm:"not null; default:false"`
	Message datatypes.JSON	`json:"message" gorm:"not null;size:50"`
	UserID uuid.UUID  `json:"user_id" gorm:"not null"` 
	
}