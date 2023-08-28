package main

import (
	"fmt"
	"lucky-draw/controllers"
	"lucky-draw/docs"
	"lucky-draw/initializers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// initializers.LoadEnvVariables() // only use for local
	initializers.ConnectToDB()
}

// Draw-system-api
//
//	@title			Lucky Draw System Api
//	@version		1.0.0
//	@description	A lucky draw system api.
//	@host			localhost:8080
func main() {

	// Create a new gin router
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	api := r.Group("/api")

	// Add Swagger documentation endpoint
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.GET("/draw/:customerId", controllers.Draw)
	api.POST("/redeem/:customerId", controllers.Redeem)

	// Start the server
	fmt.Println("Running on port 8080")
	fmt.Println("Visit swagger doc at http://localhost:8080/api/swagger/index.html")

	r.Run(":8080")
}
