package model

import "github.com/google/uuid"

type Ruang struct {
	ID			uuid.UUID `json:"id" gorm:"not null"`
	Name		string `json:"name" gorm:"not null"`
	RuangImgURL string `json:"ruang_img_url" gorm:"type:text"`
	Posts		[]Posts  `json:"posts"`
	UserID		uuid.UUID `json:"user_id" form:"user_id" gorm:"-"`
}

type RuangRespone struct{
	ID			uuid.UUID `json:"id" gorm:"not null"`
	Name		string `json:"name" gorm:"not null"`
	RuangImgURL string `json:"ruang_img_url" gorm:"type:text"`
	Posts		[]Posts  `json:"posts" gorm:"foreignKey:ruang_id;references:id"`
	UserID		uuid.UUID `json:"-" form:"user_id"`
	Users 		[]UserResponse `gorm:"many2many:anggota;foreignKey:id;joinForeignKey:ruang_id;References:id;joinReferences:user_id"`


}

type RuangPostResponse struct{
	ID uuid.UUID `json:"-"`
	Name		string `json:"name" gorm:"not null"`
	RuangImgURL string `json:"ruang_img_url" gorm:"type:text"`
}

func (RuangRespone) TableName() string{
	return "ruangs"
}


func (RuangPostResponse) TableName() string{
	return "ruangs"
}