package controller

import (
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)



func GetPosts(c *fiber.Ctx) error{
	db := database.DBConn

	context:= fiber.Map{
		"status": "Get All Post",
		
	}

	var posts []model.Posts

	dbFetchError:=db.Order("created_at desc").Find(&posts).Error

	if dbFetchError !=nil{
		log.Println(dbFetchError.Error())
	
	}
	context["data"] = posts 

	return c.Status(201).JSON(context)
}

func CreatePost(c *fiber.Ctx) error{

	context:= fiber.Map{
		"status": "Create Post",
	}

	record:= new(model.Posts)

	if err:= c.BodyParser(&record); err !=nil{
		return c.Status(400).JSON(fiber.Map{
			"err":"request can't be processed, failed to parse response into struct",
		})
	}

	if record.Judul == ""{
		return c.Status(400).JSON(fiber.Map{
			"err":"tidak ada judul",
		})
	}


	if record.UserID == uuid.Nil {
		return c.Status(403).JSON(fiber.Map{
			"err":"user id not found!",
		})
	} 
		
	result := database.DBConn.Create(record)

	if result.Error != nil{
		log.Println("Error menyimpan di dalam database")
	}


	context["data"] = record
	context["message"] = "buat post baru sukses"

	return c.Status(201).JSON(context)
}