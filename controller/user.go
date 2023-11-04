package controller

import (
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserList(c *fiber.Ctx) error{
	context:= fiber.Map{
		"status": "getUserList",
	}

	record:= []model.User{}

	err:=database.DBConn.Preload("Posts").Find(&record).Error

	if err !=nil{
		c.Status(500).JSON(fiber.Map{"err":"tidak dapat mengambil Posts dari database"})
	}
	context["data"] = record

	return c.Status(200).JSON(context)
}


func CreateUser(c *fiber.Ctx) error{
	context:= fiber.Map{
		"status":"creating new user.",
	}

	record := new(model.User)

	if err:= c.BodyParser(&record);err!=nil{
		log.Printf("Error in parsing Body.")
	}
	record.ID = uint(uuid.New().ID())
	// uuid.New()

	result := database.DBConn.Create(record)

	if result.Error != nil{
		log.Println("Error menyimpan di dalam database")
	}


	context["data"] = record
	context["message"] = "buat user baru sukses"

	return c.Status(201).JSON(context)
}