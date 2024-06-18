package main

import (
	"github.com/hatrungt1n/go-api-sample/configs"
	"github.com/hatrungt1n/go-api-sample/routes"

	"github.com/gin-gonic/gin"
)

func main() {
		router := gin.Default()

		//run database
		configs.ConnectDB()

		//routes
    routes.UserRoute(router)

		router.Run("localhost:6000") 
}