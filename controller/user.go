package controller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func GetUser(c *fiber.Ctx) error{
	context:= fiber.Map{
		"status":"get User",
	}

	var user model.UserResponse

	err:=database.DBConn.Preload("Posts", func(db *gorm.DB) *gorm.DB{
		return db.Order("created_at desc")
	}).Preload("Posts.User").Preload("Posts.Ruang").Preload("Posts.Comment").Preload("Posts.LikedByUser").Preload("Ruang", "id <> ?", c.Params("id")).Order("created_at desc").Find(&user, "id = ?", c.Params("id")).Error

	if err !=nil{
		context["err"] = "tidak dapat mengambil user data"
		c.Status(500).JSON(context)
	}





	context["data"] = user
	return c.Status(200).JSON(context)
}

func GetLoggedInUserInfo(c *fiber.Ctx) error{
	context:= fiber.Map{
		"status":"get User",
	}

	var user model.UserResponse

	err:=database.DBConn.Find(&user, "id = ?", c.Params("id")).Error

	if err !=nil{
		context["err"] = "tidak dapat mengambil user data"
		c.Status(500).JSON(context)
	}

	context["data"] = user
	return c.Status(200).JSON(context)
}



func Testo(c *fiber.Ctx) error{

	return c.Status(200).JSON(fiber.Map{
		"message": c.Locals("user"),
	})
}

type UserResponseToken struct{
	Username string
	ID       uuid.UUID
	Token 	string
}

func CreateUser(c *fiber.Ctx) error{
	

	record := new(model.User)

	record.PhotoURL = GetUserPicture()

	if err:= c.BodyParser(&record);err!=nil{
		log.Printf("Error in parsing Body.")
	}
	record.ID = uuid.New()

	hash, err:= HashPassword(record.Password)

	record.Password = hash
	if err != nil{
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"internal server error",
		})
	}

	result := database.DBConn.Create(record)

	if result.Error != nil{
		log.Println("Error menyimpan di dalam database")
	}

	claims := jwt.MapClaims{
		"user_name": record.UserName,
		"user_id":record.ID,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("rahasia"))

	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	returnedUserInfo := UserResponseToken{
		Username: record.UserName,
		ID: record.ID,
		Token: signedToken,
	}


	log.Println("return user: ",returnedUserInfo)
	
	c.Locals("user", returnedUserInfo)

	return c.Next()
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


