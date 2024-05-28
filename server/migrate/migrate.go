
package main

import (
	"mosnad/server/initializers"
	"mosnad/server/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	
     initializers.DB.AutoMigrate(&models.User{})
}