package model

import (
	"time"
)

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	UserName  string `json:"user_name" gorm:"not null;column:user_name;size:255"`
	Email     string `json:"email" gorm:"not null;column:email;size:255"`
	Password  string `json:"password" gorm:"not null;column:password;size:255"`
	Name      string `json:"name" gorm:"column:name;size:255"`
	Photo_URL string `json:"photo_url" gorm:"column:photo_url;type:text"`
	CreatedAt time.Time `gorm:"not null"`
	Posts []PostResponse `json:"posts"`
}