package routes

import (
	"github.com/hatrungt1n/go-api-sample/controllers"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1

func Register(router *gin.Engine)  {
    router.GET("/users", controllers.GetAllUsers())    
    router.POST("/user", controllers.CreateUser())
}