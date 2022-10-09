package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/gin-gorm-go/config"
	"github.com/karlbehrensg/gin-gorm-go/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UsersRoutes(router)
	router.Run(":8080")
}
