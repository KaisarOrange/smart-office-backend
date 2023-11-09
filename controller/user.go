package controller

import (
	"encoding/json"
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

	var users []model.UserResponse

	

	err:=database.DBConn.Preload("Posts").Preload("Ruang").Find(&users).Error

	if err !=nil{
		c.Status(500).JSON(fiber.Map{"err":"tidak dapat mengambil Posts dari database"})
	}
	context["data"] = users

	return c.Status(200).JSON(context)
}


func CreateUser(c *fiber.Ctx) error{
	context:= fiber.Map{
		"status":"creating new user.",
	}

	record := new(model.User)

	record.PhotoURL = GetUserPicture()

	if err:= c.BodyParser(&record);err!=nil{
		log.Printf("Error in parsing Body.")
	}
	record.ID = uuid.New()
	// uuid.New()

	result := database.DBConn.Create(record)

	if result.Error != nil{
		log.Println("Error menyimpan di dalam database")
	}


	context["data"] = record
	context["message"] = "buat user baru sukses"

	return c.Status(201).JSON(context)
}

type UserPicture struct{
	Results []struct{
		Picture struct {
			Medium    string `json:"medium"`
		} `json:"picture"`
	}
}

func GetUserPicture() string{
	

	result :=fiber.Get("https://randomuser.me/api/")
	result.Set("header-token", "value")

	_, data, err := result.Bytes()

	
	
	if err!=nil{
		log.Fatal(err)
	}

	var userPicture UserPicture

	Jsonerr := json.Unmarshal(data, &userPicture)

	if Jsonerr !=nil{
		log.Fatal(Jsonerr)
	}

    resultReturn := userPicture.Results[0].Picture.Medium
	

	return resultReturn

}


