package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	UserName  string `json:"user_name" gorm:"not null;column:user_name;size:255"`
	Email     string `json:"email" gorm:"not null;column:email;size:255"`
	Password  string `json:"password" gorm:"not null;column:password;size:255"`
	Name      string `json:"name" gorm:"column:name;size:255"`
	PhotoURL string `json:"photo_url" gorm:"column:photo_url;type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	Posts []Posts `json:"posts"`
	Ruang []Ruang `json:"ruang" gorm:"many2many:anggota;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:ruang_id"`
}


type UserResponse struct{
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	UserName  string `json:"user_name" gorm:"not null;column:user_name;size:255"`
	Email     string `json:"email" gorm:"not null;column:email;size:255"`
	Password  string `json:"-" gorm:"not null;column:password;size:255"`
	Name      string `json:"name" gorm:"column:name;size:255"`
	PhotoURL string `json:"photo_url" gorm:"column:photo_url;type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	Posts []Posts `json:"posts" gorm:"foreignKey:user_id;references:id"`
	Ruang []RuangRespone `json:"ruang" gorm:"many2many:anggota;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:ruang_id"`
}

func (UserResponse) TableName() string{
	return "users"
}



