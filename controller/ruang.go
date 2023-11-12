package controller

import (
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	var user model.UserResponse

	err:= database.DBConn.Take(&user, "id = ?", ruang.UserID).Error;if err!=nil{
		context["err"] = "couldn not processed request"
		log.Println(err)
		return c.Status(500).JSON(context)
	}

	errApp:=database.DBConn.Model(&ruang).Association("Users").Append(&user)

	if errApp !=nil{
		context["err"] = "couldn not processed request"
		log.Println(errApp)

		return c.Status(500).JSON(context)
	}

	

	// anggota:= new(model.Anggota)

	// anggota.RuangID = ruang.ID
	// anggota.UserID = ruang.UserID

	// database.DBConn.Create(&anggota)

	context["data"] = ruang
	return c.Status(200).JSON(context)

}

func GetRuangs(c *fiber.Ctx) error{

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

func GetRuang(c *fiber.Ctx) error{

	context := fiber.Map{
		"status":"Get Ruangs",
	}

 	var ruang model.RuangRespone

	err:= database.DBConn.Preload("Posts", func(db *gorm.DB) *gorm.DB{
		return db.Order("created_at desc")
	}).Preload("Posts.Ruang").Preload("Posts.User").Preload("Users").First(&ruang, "id = ?",c.Params("id")).Error

	if err !=nil{
		context["error_message"] = err;
		return c.Status(503).JSON(context)
	}

	 context["data"] = ruang

	 return c.Status(200).JSON(context);
}


func InsertUserIntoRuang(c *fiber.Ctx) error{
	context:= fiber.Map{
		"status":"insert user into ruang",
	}
	// id, _:= uuid.Parse("e82f960b-574f-424a-a644-0034af6766a2")
	var ruang model.RuangRespone
	var user model.UserResponse

	err:= database.DBConn.Take(&ruang, "id = ?", "afeefecb-cd78-44da-92b5-5fb4e314d705").Error
	if err !=nil{
		context["err"] = err
		log.Println("1")
		return c.Status(500).JSON(context)
	}
	errRuang:= database.DBConn.Take(&user, "id = ?", "1eda295b-dba9-4a0e-ba5f-dc6cedc1ece5").Error

	if errRuang !=nil{
		context["err"] = err
		log.Println("2")
		return c.Status(500).JSON(context)
	}

	// var user = &model.User{}

	// err:= database.DBConn.Preload("Users").Take(&ruang,"id = ?", id).Error

	errApp:=database.DBConn.Model(&ruang).Association("Users").Append(&user)

	if errApp !=nil{
		context["err"] = "couldn not processed request"
		log.Println(errApp)

		return c.Status(500).JSON(context)
	}

	context["data"] = ruang

	// if err !=nil {
	// 	context["err"] = err
	// 	return c.Status(500).JSON(context)
	// }

	return c.Status(200).JSON(context);
	
}