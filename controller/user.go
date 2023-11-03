package controller

import (
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error{
	context:= fiber.Map{
		"status": "getUserList",
	}

record:= new(model.User)

	database.DBConn.Find(&record)

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

	result := database.DBConn.Create(record)

	if result.Error != nil{
		log.Println("Error menyimpan di dalam database")
	}


	context["data"] = record
	context["message"] = "buat user baru sukses"

	return c.Status(201).JSON(context)
}