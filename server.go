package main

import (
	"log"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/router"
	"github.com/gofiber/fiber/v2"
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

	router.Routes(app)

	app.Listen("localhost:8080")
}