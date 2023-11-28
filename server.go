package main

import (
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init(){
	if err:= godotenv.Load(".env"); err!=nil{
		log.Fatal("Error loading env file.")
	}

	database.ConnectToDatabase()

}

func main() {

	
	db, err  :=  database.DBConn.DB()

	if err != nil{
		panic("Error in postgress connection")
	}

	defer db.Close()
	

	app := fiber.New()

	// controller.GetUserPicture()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	router.Routes(app)
	// log.Fatal(app.ListenTLS(":443", "./127.0.0.1.pem", "./127.0.0.1-key.pem"))
	app.Listen("localhost:8080")
}