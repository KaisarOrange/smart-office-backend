package controller

import (
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
		return db.Order("created_at desc")}).Preload("Ruang.Posts.Ruang").Preload("Ruang.Posts.Comment").Preload("Ruang.Posts.User").Take(&user,"id = ?", c.Params("id")).Error

	
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
	context["data"] = &posts

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
		return db.Order("created_at desc")}).Preload("Ruang.Posts.Ruang").Preload("Ruang.Posts.User").Take(&user,"id = ?", c.Params("id")).Error

	
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
	context["data"] = &posts

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

func CreateComment(c *fiber.Ctx) error{
	context := fiber.Map{"message":"create comment"}
	comment := new(model.Comment)

	// var commentText = []model.CommentText{}

	if c.Locals("posts") !=nil{
		post:= c.Locals("posts").(*model.Posts)
		comment.PostsID = post.ID
		// comment.Comments = commentText
		comment.Comments = datatypes.JSON([]byte(`[]`))

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
	}

	// commentResult := new(model.Comment)
	// commentResult.Comments = comment.Comments

	if comment.PostsID == 0{
		return c.Status(403).JSON(fiber.Map{
			"err":"post id not found!",
		})
	}

	// comment.Comments = commentResult.Comments
	// commentGo := []model.CommentText{{UserName: "Alif Boyke", UserImage: "https://source.unsplash.com/hr7eefjrekI",
	//  Text: "Ini golang canggih dan rumit", Like: 25, Comments: &[]model.CommentText{{UserName: "Nurdin", UserImage: "https://source.unsplash.com/hr7eefjrekI", Text: "good game!", Like: 3, Comments: &[]model.CommentText{}}}}}
	
	
	//  comment = &model.Comment{
	// 	PostsID: 15,
	// 	Comments: datatypes.JSON(comment.Comments),
		
	// 	// Comments: datatypes.NewJSONSlice(commentGo),

	// }	  

	// comment.Comments = datatypes.JSON(comment.Comments)

	err := database.DBConn.Model(&comment).Where("posts_id = ?", comment.PostsID).Update("comments", comment.Comments).Error

	if err!=nil{
		context["err"]= err.Error()
		log.Println(err.Error())
		return c.Status(503).JSON(context)
	}

	context["data"] = comment
	return c.Status(201).JSON(context)
}

