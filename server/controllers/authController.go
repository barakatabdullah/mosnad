package controllers

import (
	"mosnad/server/initializers"
	"mosnad/server/models"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {

	var registerInput models.RegisterInput
	if err := c.ShouldBindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if  registerInput.Password!=registerInput.Passwordconfirmation{
		c.JSON(http.StatusBadRequest, gin.H{"error": "password doesn't match"})
		return
	}
	var userFound models.User
	initializers.DB.Where("email=?", registerInput.Email).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}
	

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: registerInput.Username,
		Email: registerInput.Email,
		Role: "client",
		Phone: int(registerInput.Phone),
		Password: string(passwordHash),
	}
	
	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func Login(c *gin.Context) {

	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email=?", authInput.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}


func GetUserProfile(c *gin.Context) {

	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})
}