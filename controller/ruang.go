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

	userJWT := UserResponseToken{}


	if c.Locals("user") !=nil{

		user := c.Locals("user").(UserResponseToken)
		ruang.ID = user.ID
		ruang.UserID = user.ID
		ruang.Name = user.Username
		userJWT.ID = user.ID
		userJWT.Token = user.Token
		userJWT.Username = user.Username

	}else{
		
		if err:= c.BodyParser(&ruang);err!=nil{
			context["err"] = err.Error()
			context["message"] = "couldnt not handle request"
			log.Println(err.Error())

			return c.Status(403).JSON(context)
		}

		ruang.ID = uuid.New()
	}

	if ruang.UserID == uuid.Nil {
		context["err"] = "required user id to proceed"
		return c.Status(403).JSON(context)
	}

	if ruang.Name == "" {
		context["err"] = "required name to proceed"
		return c.Status(403).JSON(context)
	}

	err:= database.DBConn.Create(ruang).Error

	if err!=nil{
		context["message"] = "Gagal menyimpan dalam database"
		context["err"] = err.Error()
		log.Println(err.Error())

		return c.Status(503).JSON(context)
	}

	var user model.UserResponse

	err = database.DBConn.Take(&user, "id = ?", ruang.UserID).Error 
	
	if err!=nil{
		context["err"] = "couldn not processed request"
		context["message"] = err.Error()
		log.Println(err.Error())

		return c.Status(500).JSON(context)
	}

	err = database.DBConn.Model(&ruang).Association("Users").Append(&user)

	if err !=nil{
		context["err"] = "couldn not processed request"
		context["message"] = err.Error()
		log.Println(err.Error())

		return c.Status(500).JSON(context)
	}

	context["data"] = ruang
	context["token"] = userJWT.Token
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

	err:= database.DBConn.Preload("Posts" ,"draft <> true AND private <> true", func(db *gorm.DB) *gorm.DB{
		return db.Order("created_at desc")
	}).Preload("Posts.Ruang").Preload("Posts.User").Preload("Posts.LikedByUser").Preload("Users").Preload("Reminders").Preload("Posts.Comment").First(&ruang, "id = ?",c.Params("id")).Error

	if err !=nil{
		context["error_message"] = err.Error();
		log.Println(err.Error())
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

	err:= database.DBConn.Take(&ruang, "id = ?", c.Params("ruang")).Error
	if err !=nil{
		context["err"] = err
		return c.Status(500).JSON(context)
	}
	err = database.DBConn.Take(&user, "user_name = ?", c.Params("user")).Error

	if err !=nil{
		context["err"] = err
		return c.Status(500).JSON(context)
	}


	err =database.DBConn.Model(&ruang).Association("Users").Append(&user)

	if err !=nil{
		context["err"] = err.Error()
		context["message"] = "couldn not processed request"

		log.Println(err.Error())

		return c.Status(500).JSON(context)
	}
	log.Println("okkkkkkkk")

	context["data"] = ruang

	return c.Status(200).JSON(context);
	
}
