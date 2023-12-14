package router

import (
	"github.com/KaisarOrange/smart-office/controller"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){
	app.Get("api/users", controller.UserList)
	app.Post("api/users", controller.CreateUser, controller.CreateRuang)
	app.Get("api/user/:id", controller.GetUser)
	app.Get("api/user/login/:id", controller.GetLoggedInUserInfo)



	app.Get("api/posts/:id", controller.GetPosts)
	app.Get("api/posts/:id/draft", controller.GetPostsDraft)
	app.Get("api/posts/:id/like", controller.GetLikePosts)
	app.Post("api/posts", controller.CreatePost, controller.CreateComment)
	app.Post("api/posts/private", controller.CreatePost)
	app.Put("api/posts", controller.UpdatePost)
	app.Delete("api/posts/delete/:id", controller.DeletePost)

	app.Put("api/posts/like", controller.LikePosts)
	app.Get("api/posts/like/:id", controller.GetPostLikeCount)


	app.Put("api/posts/comment", controller.CreateComment)

	app.Get("api/ruang", controller.GetRuangs)
	app.Get("api/ruang/:id", controller.GetRuang)
	app.Post("api/ruang", controller.CreateRuang)
	app.Put("api/ruangupdate", controller.InsertUserIntoRuang)


	app.Get("api/auth/restricted", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("rahasia")},	
	}),controller.Restricted)
	app.Post("api/auth/login", controller.Login)



}


