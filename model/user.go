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
	LikePosts	[]Posts `json:"user_like" gorm:"many2many:user_like_posts;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:posts_id"`
}


type UserResponse struct{
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	UserName  string `json:"user_name" gorm:"not null;column:user_name;size:255"`
	Email     string `json:"email" gorm:"not null;column:email;size:255"`
	Password  string `json:"-" gorm:"not null;column:password;size:255"`
	Name      string `json:"name" gorm:"column:name;size:255"`
	PhotoURL string `json:"photo_url" gorm:"column:photo_url;type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	Posts 		[]Posts `json:"posts" gorm:"foreignKey:user_id;references:id"`
	Ruang 		[]RuangRespone `json:"ruang" gorm:"many2many:anggota;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:ruang_id"`
	LikePosts	[]Posts `json:"user_like" gorm:"many2many:user_like_posts;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:posts_id"`
}

//Types for get All post
type UserGetPostAllRuang struct{
	ID        uuid.UUID `json:"-" gorm:"primaryKey"`
	Ruang []RuangUserGetPostAllRuang `json:"ruang" gorm:"many2many:anggota;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:ruang_id"`
	LikePosts	[]Posts `json:"user_like" gorm:"many2many:user_like_posts;foreignKey:id;joinForeignKey:posts_id;references:id;joinReferences:user_id"`

}

type RuangUserGetPostAllRuang struct{
	ID			uuid.UUID `json:"-"`
	Posts		[]Posts  `json:"posts" gorm:"foreignKey:ruang_id;references:id"`
}

func (RuangUserGetPostAllRuang) TableName() string{
	return "ruangs"
}

//---------------------------------------------------------------------------------//

type UserGetLikePost struct{
	ID        uuid.UUID `json:"-" gorm:"primaryKey"`
	LikePosts	[]Posts `json:"user_like" gorm:"many2many:user_like_posts;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:posts_id"`
}


//----//

type UserPostResponse struct{
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	UserName  string `json:"user_name" gorm:"not null;column:user_name;size:255"`
	Email     string `json:"email" gorm:"not null;column:email;size:255"`
	Name      string `json:"name" gorm:"column:name;size:255"`
	PhotoURL string `json:"photo_url" gorm:"column:photo_url;type:text"`
}


type UserID struct{
	ID uuid.UUID
}



func (UserResponse) TableName() string{
	return "users"
}

func (UserGetLikePost) TableName() string{
	return "users"
}

func (UserGetPostAllRuang) TableName() string{
	return "users"
}

func (UserPostResponse) TableName() string{
	return "users"
}




