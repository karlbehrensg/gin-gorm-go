package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/gin-gorm-go/controllers"
)

func UsersRoutes(router *gin.Engine) {
	router.GET("/", controllers.GetUsers)
	router.GET("/:id", controllers.GetUser)
	router.POST("/", controllers.CreateUser)
	router.PUT("/:id", controllers.UpdateUser)
	router.DELETE("/:id", controllers.DeleteUser)
}
