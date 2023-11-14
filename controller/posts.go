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

	dbFetchError:=db.Preload("User").Preload("Ruang").Order("created_at desc").Find(&posts).Error

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
			"message": err.Error(),
		})
	}
	if record.RuangID == uuid.Nil{
		return c.Status(400).JSON(fiber.Map{
			"err":"Ruang belum dipilih!",
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

	if record.Private {
		record.RuangID = record.UserID
	}
		
	result := database.DBConn.Create(record)

	if result.Error != nil{
		log.Println("Error menyimpan di dalam database")
		context["err"] = result.Error.Error()
		return c.Status(400).JSON(context)
	}


	context["data"] = record
	context["message"] = "buat post baru sukses"

	return c.Status(201).JSON(context)
}

