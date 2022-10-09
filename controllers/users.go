package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/gin-gorm-go/config"
	"github.com/karlbehrensg/gin-gorm-go/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetUser(c *gin.Context) {
	user := models.User{}
	config.DB.Where("id = ?", c.Params.ByName("id")).First(&user)
	c.JSON(200, &user)
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	user := models.User{}
	config.DB.Where("id = ?", c.Params.ByName("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	user := models.User{}
	config.DB.Where("id = ?", c.Params.ByName("id")).Delete(&user)
	c.JSON(200, &user)
}
