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

func SendMentionNotif(c *fiber.Ctx) error{
	type mentionJSON struct{
		SenderID uuid.UUID		`json:"sender_id"`
		RecieverID uuid.UUID	`json:"reciever_id"` 
		MessageNotif string		`json:"message"`
		PostTitle	string 		`json:"post_title"`
		PostID   uint    		`json:"post_id"`
		RuangID   uuid.UUID    	`json:"ruang_id"`
		
	}

	type mention struct{
		MentioneUsers []string `json:"mentioned_users"` 
		SenderID uuid.UUID		`json:"sender_id"`
		PostTitle	string 		`json:"post_title"`
		PostID   uint    		`json:"posts_id"`
		RuangID   uuid.UUID    	`json:"ruang_id"`
	}

	mentionInfo := mention{}
	
	err:= c.BodyParser(&mentionInfo)
	
	if err!=nil{
		log.Println("Error parsing data")		
		return c.Status(400).JSON(fiber.Map{"err":err.Error()})
	}

	for _, user := range mentionInfo.MentioneUsers{

		type User struct{
			ID uuid.UUID 	`json:"user_id"`
			Username string `json:"user_name"`
		}

		
		userInfo:= new(User)
		
		err:= database.DBConn.Take(&userInfo, "user_name = ?", user).Error

		if err!=nil{
			log.Println("Error parsing data")		
			return c.Status(400).JSON(fiber.Map{"err":err.Error()})
		}

		mentionJson:= new(mentionJSON)

		mentionJson.MessageNotif = "User mentioned you in a post!"
		mentionJson.PostID = mentionInfo.PostID
		mentionJson.PostTitle = mentionInfo.PostTitle
		mentionJson.RecieverID = userInfo.ID
		mentionJson.RuangID = mentionInfo.RuangID
		mentionJson.SenderID = mentionInfo.SenderID

		res, _ := json.Marshal(&mentionJson)


		notif := model.Notification{
			UserID: userInfo.ID,
			Dibaca: false,
			Type: "mention",
			Message: res,
		}
	
		result := database.DBConn.Create(&notif)
	
		if result.Error != nil{
			log.Println("Error menyimpan di dalam database")		
			return c.Status(400).JSON(fiber.Map{"err":result.Error.Error()})
		}
	} 


	
	return c.Status(200).JSON(fiber.Map{"message":"sent mention notifs"})
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
		RuangID		uuid.UUID	`json:"ruang_id"`
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