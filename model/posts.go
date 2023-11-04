package model

import (
	"time"
)

type Posts struct {
	ID   uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null"`
	Judul string `json:"judul" gorm:"not null;size:50"`
	Konten string `json:"konten" gorm:"type:text"`
	CreatedAt time.Time `gorm:"not null"`
}

type PostResponse struct{
	ID   uint `json:"id"`
	UserID uint `json:"-"`
	Judul string `json:"judul"`
	Konten string `json:"konten"`
	CreatedAt time.Time 
}

func (PostResponse) TableName() string{
	return "posts"
}