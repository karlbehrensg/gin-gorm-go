package config

import (
	"github.com/karlbehrensg/gin-gorm-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=gin_gorm_go port=5432"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
