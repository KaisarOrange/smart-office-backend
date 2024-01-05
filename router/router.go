package router

import (
	"time"

	w "github.com/KaisarOrange/smart-office/pkg/webrtc"

	"github.com/KaisarOrange/smart-office/controller"
	"github.com/KaisarOrange/smart-office/webrtc"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Routes(app *fiber.App){
	app.Get("api/users", controller.UserList)
	app.Post("api/users", controller.CreateUser, controller.CreateRuang)
	app.Get("api/user/:id", controller.GetUser)
	app.Get("api/user/login/:id", controller.GetLoggedInUserInfo)



	app.Get("api/posts/:id", controller.GetPosts)
	app.Get("api/post/:id", controller.GetPost)
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


	app.Post("api/auth/notif", controller.SendNotification)
	app.Get("api/user/notif/:id", controller.GetNotifs)
	app.Post("api/user/notif/mention", controller.SendMentionNotif)

	app.Put("api/user/reminder", controller.SetReminder)


	
//LiveCollab



//WEBRTC

	app.Get("/", webrtc.Welcome)
	app.Get("/room/create", webrtc.RoomCreate)
	app.Get("/room/:uuid", webrtc.Room)
	app.Get("/room/:uuid/websocket", websocket.New(webrtc.RoomWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	// app.Get("/room/:uuid/chat", webrtc.RoomChat)
	// app.Get("/room/:uuid/chat/websocket", websocket.New(webrtc.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(webrtc.RoomViewerWebsocket))
	app.Get("/stream/:suuid", webrtc.Stream)
	app.Get("/stream/:suuid/websocket", websocket.New(webrtc.StreamWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	// app.Get("/stream/:suuid/chat/websocket", websocket.New(webrtc.StreamChatWebsocket))
	app.Get("/stream/:suuid/viewer/websocket", websocket.New(webrtc.StreamViewerWebsocket))
	// app.Static("/", "./assets")

	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)
	go dispatchKeyFrames()

	//Notifs SSE


}
func dispatchKeyFrames() {
	for range time.NewTicker(time.Second * 3).C {
		for _, room := range w.Rooms {
			room.Peers.DispatchKeyFrame()
		}
	}
}



