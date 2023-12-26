package controller

import (
	"log"
	"time"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginRequest struct{
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}




func Restricted(c *fiber.Ctx) error {

	


	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status":"Restricted"})
}

func Login(c *fiber.Ctx) error{

	userLoginInfo := UserLoginRequest{}
	
	err:= c.BodyParser(&userLoginInfo)

	if err !=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}

	validate := validator.New()

	isValidateErr := validate.Struct(userLoginInfo)

	if isValidateErr != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": isValidateErr.Error()})

	}

	user:= model.User{}

	err = database.DBConn.Take(&user, "user_name = ?", userLoginInfo.UserName).Error

	if err !=nil{
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"err": err.Error()})
	}

	isUserValid := CheckHashPassword(userLoginInfo.Password, user.Password)
	

	if !isUserValid{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message":"wrong credentials"})
	}

	claims := jwt.MapClaims{
		"user_name": user.UserName,
		"user_id":user.ID,
		"exp": time.Now().Add(time.Minute * 60).Unix(),
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("rahasia"))

	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	// cookie := fiber.Cookie{
    //     Name: "token",
    //     Value: signedToken,
    //     HTTPOnly: true,
    // }

	// 	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": signedToken})
}


func HashPassword(password string) (string,error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password),14)

	if err !=nil{
		return "", err
	}

	return string(hash), err
}

func CheckHashPassword(password string, hashPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	return err == nil
}