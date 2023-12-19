package controller

import (
	"encoding/json"
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SendNotification(c *fiber.Ctx) error{
	
	type postIdUserId struct{
		ID uint
		UserID uuid.UUID `json:"user_id"`
		PostID uint `json:"posts_id"`

	}

	post := new(model.Posts)

	ids := postIdUserId{}

	err:= c.BodyParser(&ids)
	
	if err != nil{
		log.Println(err.Error())
		c.Status(503).JSON(fiber.Map{})
	}

	err = database.DBConn.Take(&post, "id = ?", ids.PostID).Error

	
	if err != nil{
		log.Println(err.Error())
		c.Status(503).JSON(fiber.Map{})
	}
	
	notification:= []model.Notification{}

	err = database.DBConn.Find(&notification, "message @> ?", map[string]interface{}{"post_id": ids.PostID, "sender_id": ids.UserID, "reciever_id": post.UserID}).Error


	if err != nil{
		log.Println(err.Error())
		c.Status(503).JSON(fiber.Map{"err": err.Error()})
	}


	return c.Status(200).JSON(fiber.Map{
		"data":notification,
	})
}


func GetNotifs(c *fiber.Ctx) error{
	

	type userInfo struct{
		Message string	`json:"message"`
		PostID uint		`json:"post_id"`
		PostTitle string `json:"post_title"`
		SenderID uuid.UUID	`json:"sender_id"`
		RecieverID uuid.UUID	`json:"reciever_id"`
		Username	string		`json:"username"`
		UserPicture string		`json:"user_picture"`
	}

	
	
	notification := []model.Notification{}
	

	err:= database.DBConn.Where("dibaca = ?", false).Find(&notification, "user_id = ?", c.Params("id")).Error

	for i, v := range notification{

		result := userInfo{}
		user := model.UserResponse{}

		
		json.Unmarshal([]byte(v.Message), &result)

		err = database.DBConn.Take(&user, "id = ?", result.SenderID).Error

		
		if err != nil{
			log.Println(err.Error())
			return c.Status(503).JSON(fiber.Map{"err": err.Error()})
		}

		result.UserPicture = user.PhotoURL
		result.Username = user.UserName

		resJson, err := json.Marshal(result)

		if err != nil{
			log.Fatal(err.Error())
		}

		notification[i].Message = resJson

		
	}


	if err != nil{
		log.Println(err.Error())
		return c.Status(503).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": notification})
}