package router

import (
	"github.com/KaisarOrange/smart-office/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){
	app.Get("/users", controller.UserList)
	app.Post("/users", controller.CreateUser)

	app.Get("api/posts", controller.GetPosts)
	app.Post("api/posts", controller.CreatePost)


}