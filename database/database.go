package database

import (
	"log"

	"github.com/KaisarOrange/smart-office/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectToDatabase(){
	dsn:= "host=localhost user=postgres password=Captain10 dbname=smart_office port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil{
		panic("Koneksi database gagal.")
	}
	log.Println("Koneksi database berhasil!")

	db.AutoMigrate(new(model.User), new(model.Posts), new(model.Ruang), new(model.Comment))

	DBConn = db
}

