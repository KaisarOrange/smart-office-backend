package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Posts struct {
	ID   uint `json:"id" gorm:"primaryKey;aut"`
	UserID uuid.UUID `json:"user_id" gorm:"not null"`
	Judul string `json:"judul" gorm:"not null;size:50"`
	Konten datatypes.JSON `json:"konten" gorm:"type:json"`
	CreatedAt time.Time `gorm:"not null"`
}

type PostResponse struct{
	ID   uint `json:"id"`
	UserID uuid.UUID `json:"-"`
	Judul string `json:"judul"`
	Konten string `json:"konten"`
	CreatedAt time.Time 
}

func (PostResponse) TableName() string{
	return "posts"
}