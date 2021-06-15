package main

import (
	"os"

	"authentication/config"
	"authentication/router"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		router.AuthRouter(client.Group("/auth"))
		router.AdminRouter(client.Group("/admin"))
		router.UserRouter(client.Group("/user"))
	}
	
	return r
}

func main() {
	port := os.Getenv("PORT");
	if port == "" {
		port = "8088"
	}
	router := setupRouter()
	config.ConnectDatabase();
	router.Run(":" + port)
}