package router

import (
	"github.com/KaisarOrange/smart-office/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){
	app.Get("api/users", controller.UserList)
	app.Post("api/users", controller.CreateUser, controller.CreateRuang)
	app.Get("api/user/:id", controller.GetUser)


	app.Get("api/posts/:id", controller.GetPosts)
	app.Post("api/posts", controller.CreatePost)
	app.Post("api/posts/private", controller.CreatePost)


	app.Get("api/ruang", controller.GetRuangs)
	app.Get("api/ruang/:id", controller.GetRuang)
	app.Post("api/ruang", controller.CreateRuang)
	app.Put("api/ruangupdate", controller.InsertUserIntoRuang)
}