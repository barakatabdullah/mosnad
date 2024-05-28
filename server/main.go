package main

import (
	"mosnad/server/middlewares"
	"mosnad/server/controllers"
	"mosnad/server/initializers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func init(){
	initializers.LoadEnvs()
	initializers.ConnectDB()
}


func main(){
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://127.0.0.1:8080"},
        AllowMethods:     []string{"PUT", "PATCH"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "https://github.com"
        },
        MaxAge: 12 * time.Hour,
    }))
	router.POST("/auth/register", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
	router.Run()
}