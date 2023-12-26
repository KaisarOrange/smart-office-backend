package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Posts struct {
	ID			uint 				`json:"id" gorm:"primaryKey;"`
	UserID		uuid.UUID 			`json:"user_id" gorm:"not null"`
	RuangID		uuid.UUID 			`json:"ruang_id"`
	Judul		string 				`json:"judul" gorm:"not null;size:50"`
	Konten		datatypes.JSON 		`json:"konten" gorm:"type:json"`
	CreatedAt	time.Time 			`json:"created_at" gorm:"not null"`
	User 		UserPostResponse 	`json:"user" gorm:"foreignKey:user_id;references:id"`
	Ruang 		RuangPostResponse 	`json:"ruang" gorm:"foreignKey:ruang_id;references:id"`
	Draft		bool				`json:"draft" gorm:"default:false"`
	Private		bool				`json:"private" gorm:"default:false"`		
	Comment 	Comment				`json:"comment" gorm:"constraint:OnDelete:CASCADE"` 
	LikedByUser []User				`json:"user_like" gorm:"many2many:user_like_posts;foreignKey:id;joinForeignKey:posts_id;references:id;joinReferences:user_id"`						  
	Reminder	Reminder			`json:"reminder" gorm:"foreignKey:posts_id"`			
}

type PostResponse struct{
	ID   uint `json:"id"`
	UserID uuid.UUID `json:"-"`
	RuangID uuid.UUID `json:"-"`
	Judul string `json:"judul"`
	Konten string `json:"konten"`
	CreatedAt time.Time
	LikedByUser []User `json:"user_like" gorm:"many2many:user_like_posts;foreignKey:id;joinForeignKey:posts_id;references:id;joinReferences:user_id"`						  
}

func (PostResponse) TableName() string{
	return "posts"
}

//Comments

type Comment struct {
	gorm.Model
	PostsID		uint								`json:"posts_id"`
	Comments	datatypes.JSON						`json:"comments"`
	// Comments	datatypes.JSONSlice[CommentText]	`json:"comments" gorm:"type:json[]"`

}

// type CommentText struct{
// 	UserName string `json:"user_name"`
// 	UserImage string `json:"user_img"`
// 	Text      string `json:"text"`
// 	Comments *[]CommentText `json:"comments"`
// 	Like	uint `json:"like"`
// }


//Liked
