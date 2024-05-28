package main

import (
	"mosnad/server/middlewares"
	"mosnad/server/controllers"
	"mosnad/server/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvs()
	initializers.ConnectDB()
}


func main(){
	router := gin.Default()

	router.POST("/auth/register", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
	router.Run()
}