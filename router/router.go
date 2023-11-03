package router

import (
	"github.com/KaisarOrange/smart-office/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){
	app.Get("/", controller.UserList)
	app.Post("/", controller.CreateUser)

}