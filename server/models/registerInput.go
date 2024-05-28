package models

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone int `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Passwordconfirmation string `json:"passwordconfirmation" binding:"required"`
}