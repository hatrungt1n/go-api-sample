package main

import (
	"github.com/hatrungt1n/go-api-sample/configs"
	"github.com/hatrungt1n/go-api-sample/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/hatrungt1n/go-api-sample/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()

	// set up gin-swagger middleware
	// swagger ui: http://localhost:8080/swagger/index.html#/
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// run database
	configs.ConnectDB()

	// routes
	routes.Register(router)

	router.Run()
}
