package main

import (
	"jwt-golang/controllers"
	"jwt-golang/database"
	"jwt-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("root:secure_pass_here@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()

	// Initialize Router
	router := iniRouter()
	router.Run(":8080")

}

func iniRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
