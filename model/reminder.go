package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reminder struct {
	gorm.Model
	Title 			string `json:"title"`
	CompletedTask 	int `json:"completed_task"`
	TotalTask		int	`json:"total_task"`
	DueTime			time.Time `json:"due_time"`
	RuangID			uuid.UUID `json:"ruang_id"`
	PostsID          uint	`json:"post_id"`
	ReminderUsers   []User	`json:"reminders" gorm:"many2many:users_reminder;foreignKey:id;joinForeignKey:reminder_id;references:id;joinReferences:user_id"`					
}