package main

import (
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	
	app.Use(logger.New())



	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://192.168.100.35:5173, http://localhost:5173",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
		AllowCredentials:true,
	}))




	router.Routes(app)
	// log.Fatal(app.ListenTLS("192.168.100.35:8080", "./127.0.0.1.pem", "./127.0.0.1-key.pem"))
	app.Listen("127.0.0.1:8080")
}