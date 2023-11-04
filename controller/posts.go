package controller

import (
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error{

	context:= fiber.Map{
		"status": "Get All Post",
		
	}

	var posts []model.Posts

	database.DBConn.Find(&posts)
	context["data"] = posts 

	return c.Status(201).JSON(context)
}

func CreatePost(c *fiber.Ctx) error{

	context:= fiber.Map{
		"status": "Create Post",
	}

	record:= new(model.Posts)

	if err:= c.BodyParser(&record); err !=nil{
		return c.Status(503).JSON(fiber.Map{
			"err":"failed to handle request",
		})
	}

	if record.Judul == ""{
		return c.Status(400).JSON(fiber.Map{
			"err":"tidak ada judul",
		})
	}


	if record.UserID == 0 {
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