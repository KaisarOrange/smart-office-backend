package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Posts struct {
	ID			uint `json:"id" gorm:"primaryKey;"`
	UserID		uuid.UUID `json:"user_id" gorm:"not null"`
	RuangID		uuid.UUID `json:"ruang_id"`
	Judul		string `json:"judul" gorm:"not null;size:50"`
	Konten		datatypes.JSON `json:"konten" gorm:"type:json"`
	CreatedAt	time.Time `json:"created_at" gorm:"not null"`
	User 		UserPostResponse `json:"user" gorm:"foreignKey:user_id;references:id"`
	Ruang 		RuangPostResponse `json:"ruang" gorm:"foreignKey:ruang_id;references:id"`
}

type PostResponse struct{
	ID   uint `json:"id"`
	UserID uuid.UUID `json:"-"`
	RuangID uuid.UUID `json:"-"`
	Judul string `json:"judul"`
	Konten string `json:"konten"`
	CreatedAt time.Time 
}

func (PostResponse) TableName() string{
	return "posts"
}