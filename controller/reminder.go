package controller

import (
	"log"
	"time"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SetReminder(c *fiber.Ctx) error { 
	reminder := new(model.Reminder)

	type ReqReminder struct{
	Title         string    `json:"title"`
    CompletedTask int       `json:"completed_task"`
    TotalTask     int       `json:"total_task"`
    DueTime       time.Time `json:"due_time"`
    MentionedUsers []string `json:"mentioned_users"` 
	RuangID	       string 	`json:"ruang_id"`
    PostsID       uint      `json:"post_id"`
	}

	var req ReqReminder

	err:= c.BodyParser(&req)
	if err!=nil{
		c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	ruangUUID, err := uuid.Parse(req.RuangID)

	if err!=nil{
		c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}
	log.Println(req)
	err = database.DBConn.Take(&reminder, "posts_id = ?", req.PostsID).Error

	if err!=nil{
		c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}
	reminder.CompletedTask = req.CompletedTask
	reminder.Title = req.Title
	reminder.TotalTask = req.TotalTask
	reminder.PostsID = req.PostsID

	if reminder.ID == 0{
		reminder.DueTime = req.DueTime
		reminder.RuangID = ruangUUID
	}



	err = database.DBConn.Save(reminder).Error


	if err!=nil{
		c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	for _,v := range req.MentionedUsers{
		log.Println("hehehehehe", v)

		user:= new(model.UserResponse)

		
		err = database.DBConn.Take(&user, "user_name = ?", v).Error

		if err != nil{
			
			log.Println("hehe: ", err.Error())
			c.Status(503).JSON(fiber.Map{"err":err.Error()})
		}

		log.Println("hehehehehe", user)


		// err = database.DBConn.Model(&user).Association("Reminders").Append(&reminder)
		err = database.DBConn.Table("users_reminder").Create(map[string]interface{}{
			"user_id": user.ID,
			"reminder_id":reminder.ID,
		}).Error

		if err != nil{	
			log.Println("hehe: ", err.Error())
			c.Status(503).JSON(fiber.Map{"err":err.Error()})
		}
	}

	return c.Status(200).JSON(fiber.Map{"data":reminder})
}