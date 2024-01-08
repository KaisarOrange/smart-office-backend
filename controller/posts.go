package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/KaisarOrange/smart-office/database"
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)



func GetPosts(c *fiber.Ctx) error{
	db := database.DBConn

	context:= fiber.Map{
		"status": "Get All Post",
		
	}

	var ruang []model.RuangRespone

	err:=database.DBConn.Take(&ruang).Error
	
	if err !=nil{
		log.Println(err.Error())
		context["err"]= err.Error()
		c.Status(503).JSON(context)
	}

	var user model.UserGetPostAllRuang

	// err:=db.Preload("User").Preload("Ruang").Order("created_at desc").Find(&posts).Error
	err = db.Preload("Ruang").Preload("Ruang.Posts","private <> true AND draft <> true", func(db *gorm.DB) *gorm.DB{
		return db.Order("created_at desc")}).Preload("Ruang.Posts.Ruang").Preload("Ruang.Posts.Comment").Preload("Ruang.Posts.User").Preload("Ruang.Posts.LikedByUser","id = ?", c.Params("id")).Take(&user,"id = ?", c.Params("id")).Error

	
	if err !=nil{
		log.Println(err.Error())
		context["err"]= err.Error()
		c.Status(503).JSON(context)
	
	}

	var posts []model.Posts

	for _, users := range user.Ruang{
		posts = append(posts, users.Posts...)
	}


	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
	context["data"] = posts

	return c.Status(201).JSON(context)
}

func GetPost(c *fiber.Ctx) error{
	db := database.DBConn

	context:= fiber.Map{
		"status": "Get Post",
	}

	posts := model.Posts{}

	err:= db.Take(&posts, c.Params("id")).Error
	
	
	if err !=nil{
		log.Println(err.Error())
		context["err"]= err.Error()
		c.Status(503).JSON(context)
	}

	userId := posts.UserID.String()

	fmt.Println("ini strign: ", userId)

	err = db.Preload("Comment").Preload("User").Preload("LikedByUser", "id = ?" , userId).Take(&posts,"id = ?", c.Params("id")).Error
  
	if err !=nil{
		log.Println(err.Error())
		context["err"]= err.Error()
		c.Status(503).JSON(context)
	
	}

	context["data"] = posts

	return c.Status(201).JSON(context)
}



func GetPostsDraft(c *fiber.Ctx) error{
	db := database.DBConn

	context:= fiber.Map{
		"status": "Get All Post",
		
	}

	var ruang []model.RuangRespone

	err:=database.DBConn.Take(&ruang).Error
	
	if err !=nil{
		log.Println(err.Error())
		context["err"]= err.Error()
		c.Status(503).JSON(context)
	}

	var user model.UserGetPostAllRuang

	// err:=db.Preload("User").Preload("Ruang").Order("created_at desc").Find(&posts).Error
	err = db.Preload("Ruang").Preload("Ruang.Posts.Comment").Preload("Ruang.Posts","draft = true", func(db *gorm.DB) *gorm.DB{
		return db.Order("created_at desc")}).Preload("Ruang.Posts.Ruang").Preload("Ruang.Posts.User").Preload("Ruang.Posts.LikedByUser").Take(&user,"id = ?", c.Params("id")).Error

	
	if err !=nil{
		log.Println(err.Error())
		context["err"]= err.Error()
		c.Status(503).JSON(context)
	
	}

	var posts []model.Posts

	for _, users := range user.Ruang{
		posts = append(posts, users.Posts...)
	}


	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
	context["data"] = posts

	return c.Status(201).JSON(context)
}

func CreatePost(c *fiber.Ctx) error{

	context:= fiber.Map{
		"status": "Create Post",
	}

	record:= new(model.Posts)

	if err:= c.BodyParser(&record); err !=nil{
		return c.Status(400).JSON(fiber.Map{
			"err":"request can't be processed, failed to parse response into struct",
			"message": err.Error(),
		})
	}
	if record.RuangID == uuid.Nil{
		return c.Status(400).JSON(fiber.Map{
			"err":"Ruang belum dipilih!",
		})
	}

	if record.Judul == ""{
		return c.Status(400).JSON(fiber.Map{
			"err":"tidak ada judul",
		})
	}


	if record.UserID == uuid.Nil {
		return c.Status(403).JSON(fiber.Map{
			"err":"user id not found!",
		})
	} 

	if record.Private {
		record.RuangID = record.UserID
	}
		
	result := database.DBConn.Create(record)

	if result.Error != nil{
		log.Println("Error menyimpan di dalam database")
		context["err"] = result.Error.Error()
		return c.Status(400).JSON(context)
	}



	context["data"] = record
	context["message"] = "buat post baru sukses"

	c.Locals("posts", record)

	return c.Next()
}

func UpdatePost(c *fiber.Ctx) error{
	context := fiber.Map{"message":"UpdatePost"}

	post := new(model.Posts)

	err:=c.BodyParser(&post)

	if err != nil{
		context["err"] = err.Error()
		return c.Status(503).JSON(context)
	}

	err = database.DBConn.Model(&post).Where("id = ?", post.ID).Update("konten", post.Konten).Error

	if err!=nil{
		context["err"] = err.Error()
		return c.Status(503).JSON(context)
	}

	context["data"] = post

	return c.Status(200).JSON(context)
}

func CreateComment(c *fiber.Ctx) error{
	context := fiber.Map{"message":"create comment"}
	comment := new(model.Comment)

	// var commentText = []model.CommentText{}

	if c.Locals("posts") !=nil{
		post:= c.Locals("posts").(*model.Posts)
		comment.PostsID = post.ID
		// comment.Comments = commentText
		comment.Comments = datatypes.JSON([]byte(`[]`))

		database.DBConn.Create(&comment)

	}else{
		if err:= c.BodyParser(&comment);err!=nil{
			context["err"] = err.Error()
			context["message"] = "couldnt not handle request"
			log.Println(err.Error())

			return c.Status(403).JSON(context)
		}
		// err := database.DBConn.Select("").First(&comment, "comments.posts_id = ?",comment.PostsID).Error
		// if err != nil{
		// 	context["err"] = err.Error()
		// 	log.Println(err.Error())
		// 	return c.Status(503).JSON(context)
		// }

			
		err := database.DBConn.Model(&comment).Where("posts_id = ?", comment.PostsID).Update("comments", comment.Comments).Error

		if err!=nil{
			context["err"]= err.Error()
			log.Println(err.Error())
			return c.Status(503).JSON(context)
		}
	}

	// commentResult := new(model.Comment)
	// commentResult.Comments = comment.Comments

	// comment.Comments = commentResult.Comments
	// commentGo := []model.CommentText{{UserName: "Alif Boyke", UserImage: "https://source.unsplash.com/hr7eefjrekI",
	//  Text: "Ini golang canggih dan rumit", Like: 25, Comments: &[]model.CommentText{{UserName: "Nurdin", UserImage: "https://source.unsplash.com/hr7eefjrekI", Text: "good game!", Like: 3, Comments: &[]model.CommentText{}}}}}
	
	
	//  comment = &model.Comment{
	// 	PostsID: 15,
	// 	Comments: datatypes.JSON(comment.Comments),
		
	// 	// Comments: datatypes.NewJSONSlice(commentGo),

	// }	  

	// comment.Comments = datatypes.JSON(comment.Comments)


	context["data"] = comment
	return c.Status(201).JSON(context)
}

func DeletePost(c *fiber.Ctx) error{
	context:= fiber.Map{
		"message":"delete Post",
	}

	post:= new(model.Posts)
 
	err:= database.DBConn.First(&post, "id = ?", c.Params("id")).Error

	if err != nil{
		context["err"] = err.Error()
		c.Status(503).JSON(context)
		log.Println(err.Error(), "1")
	}

	// comment := model.Comment{}

	// err = database.DBConn.First(&comment, "posts_id = ?", c.Params("id")).Error

	// if err !=nil{
	// 	if err != nil{
	// 		context["err"] = err.Error()
	// 		log.Println(err.Error(), "2")
	// 		c.Status(503).JSON(context)
	// 	}	
	// }

	// err = database.DBConn.Unscoped().Model(&post).Association("Comment").Unscoped().Clear()
 
	err = database.DBConn.Select("Comment","LikedByUser").Delete(&post).Error
	if err != nil{
		context["err"] = err.Error()
		log.Println(err.Error(), "3")

		c.Status(503).JSON(context)
	}
	return c.Status(200).JSON(context)
}

func LikePosts(c *fiber.Ctx) error{
	context:= fiber.Map{"status":"like post"}

	user := model.User{}
	post := model.Posts{}

	type postIdUserId struct{
		ID uint
		UserID uuid.UUID `json:"user_id"`
		PostID uint `json:"posts_id"`

	}

	ids := postIdUserId{}

	err:= c.BodyParser(&ids)
	
	if err != nil{
		context["err"]= err.Error()
		log.Println(err.Error())
		c.Status(503).JSON(context)
	}

	err = database.DBConn.Take(&post, "id = ?", ids.PostID).Error
	
	if err != nil{
		context["err"]= err.Error()
		log.Println(err.Error())
		c.Status(503).JSON(context)
	}


	err = database.DBConn.Model(&post).Association("LikedByUser").Find(&user, "id = ?", ids.UserID)


	if err != nil{
		context["err"]= err.Error()
		log.Println("hehe: ", err.Error())
		c.Status(503).JSON(context)
	}

	if user.ID != uuid.Nil{
	
		err = database.DBConn.Take(&user, "id = ?", ids.UserID).Error

		if err != nil{
			context["err"]= err.Error()
			log.Println(err.Error())
			c.Status(503).JSON(context)
		}
	
		err = database.DBConn.Model(&post).Association("LikedByUser").Delete(&user)
	
		if err != nil{
			context["err"]= err.Error()
			log.Println(err.Error())
			c.Status(503).JSON(context)
		}
	
	
	
		context["result ganteng"] = user
		return c.Status(200).JSON(context)
	}

	err = database.DBConn.Take(&user, "id = ?",ids.UserID).Error

	if err != nil{
		context["err"]= err.Error()
		log.Println(err.Error())
		c.Status(503).JSON(context)
	}

	err = database.DBConn.Take(&post, "id = ?",ids.PostID).Error

	if err != nil{
		context["err"]= err.Error()
		log.Println(err.Error())
		c.Status(503).JSON(context)
	}
	err = database.DBConn.Model(&post).Association("LikedByUser").Append(&user)

	if err != nil{
		context["err"]= err.Error()
		log.Println(err.Error())
		c.Status(503).JSON(context)
	}

	//Send Notifications


	notification:= []model.Notification{}

	err = database.DBConn.Find(&notification, "message @> ?", map[string]interface{}{"post_id": ids.PostID, "sender_id": ids.UserID, "reciever_id": post.UserID}).Error


	if err != nil{
		context["err"]= err.Error()
		log.Println(err.Error())
		c.Status(503).JSON(context)
	}
	log.Println("ini notif", notification, len(notification))

	if len(notification) > 0 {

		return c.Status(200).JSON(context)
	}

	
	type like struct{
		SenderID uuid.UUID		`json:"sender_id"`
		RecieverID uuid.UUID	`json:"reciever_id"` 
		MessageNotif string		`json:"message"`
		PostTitle	string 		`json:"post_title"`
		PostID   uint    		`json:"post_id"`
		RuangID   uuid.UUID    	`json:"ruang_id"`
	}

	likeNotif :=like{
		SenderID: ids.UserID,
		RecieverID: post.UserID,
		PostTitle: post.Judul,
		MessageNotif: "User like your post!",
		PostID: ids.PostID,
		RuangID: post.RuangID,
	}

	res, _ := json.Marshal(&likeNotif)

	notif := model.Notification{
		UserID: post.UserID,
		Dibaca: false,
		Type: "like",
		Message: res,
	}

	result := database.DBConn.Create(&notif)

	if result.Error != nil{
		log.Println("Error menyimpan di dalam database")
		context["err"] = result.Error.Error()
		return c.Status(400).JSON(context)
	}



	// context["data"] = post
	context["user"] = user
	context["notif"] = notif


	return c.Status(200).JSON(context)
}


func GetPostLikeCount(c *fiber.Ctx) error{
	context:= fiber.Map{
		"status" :"Get Post Likes Count",
	}

	var count int64 

	err := database.DBConn.Table("user_like_posts").Where("posts_id = ?", c.Params("id")).Count(&count).Error
	
	
	if err != nil{
		context["err"]= err.Error()
		log.Println(err.Error())
		c.Status(503).JSON(context)
	}

	context["like_count"] = count

	return c.Status(200).JSON(context)
}

func GetLikePosts(c *fiber.Ctx) error{
	db := database.DBConn

	context:= fiber.Map{
		"status": "Get All Post",
		
	}

	var user model.UserGetLikePost

	// err:=db.Take(&user).Error
	
	// if err !=nil{
	// 	log.Println(err.Error())
	// 	context["err"]= err.Error()
	// 	c.Status(503).JSON(context)
	// }

	// err:=db.Preload("User").Preload("Ruang").Order("created_at desc").Find(&posts).Error
	err := db.Preload("LikePosts").Preload("LikePosts.Comment").Preload("LikePosts.LikedByUser").Preload("LikePosts.Ruang").Preload("LikePosts.User").Take(&user,"id = ?", c.Params("id")).Error

	log.Println("params: ", c.Params("id"))
	if err !=nil{
		log.Println(err.Error())
		context["err"]= err.Error()
		c.Status(503).JSON(context)
	
	}

	context["data"] = user

	return c.Status(201).JSON(context)
}

func IsUserAllowedToEdit(c *fiber.Ctx) error{

post := new(model.Posts)

err:= database.DBConn.Preload("Reminder.ReminderUsers", "id", c.Params("id")).Take(&post, c.Params("postId")).Error


if err!=nil{
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err":err.Error()})
}

if len(post.Reminder.ReminderUsers) > 0{
	return c.Status(200).JSON(fiber.Map{"message":"is user allowed to edit", "data":true})

}else{
	return c.Status(200).JSON(fiber.Map{"message":"is user allowed to edit", "data":false})
}
}