package db

import (
	"log"

	"github.com/Devin-Edwards8/DevHub/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=dev_user password=Buckeyes16 dbname=devhub port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func Migrate() {
	DB.AutoMigrate(&model.Task{})
}
