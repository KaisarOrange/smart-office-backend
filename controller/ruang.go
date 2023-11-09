package controller

import (
	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateRuang(c *fiber.Ctx) error{
	context:=fiber.Map{
		"status":"create ruang baru",
	}

	ruang:= new(model.Ruang)

	if err:= c.BodyParser(&ruang);err!=nil{
		context["err"] = "couldnt not handle request"
		return c.Status(400).JSON(context)
	}

	if ruang.UserID == uuid.Nil {
		context["err"] = "required user id to proceed"
		return c.Status(403).JSON(context)
	}

	ruang.ID = uuid.New()
	
	result:= database.DBConn.Create(ruang)

	if result.Error!=nil{
		context["err"] = "Gagal menyimpan dalam database"
		return c.Status(503).JSON(context)

	}

	anggota:= new(model.Anggota)

	anggota.RuangID = ruang.ID
	anggota.UserID = ruang.UserID

	database.DBConn.Create(&anggota)

	context["data"] = ruang
	return c.Status(200).JSON(context)

}

func GetRuang(c *fiber.Ctx) error{

	context := fiber.Map{
		"status":"Get Ruangs",
	}

 	var ruang []model.RuangRespone

	err:= database.DBConn.Preload("Posts").Preload("Users").Find(&ruang).Error

	if err !=nil{
		context["error_message"] = err;
		return c.Status(503).JSON(context)
	}

	 context["data"] = ruang

	 return c.Status(200).JSON(context);
}