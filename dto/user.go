package dto

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct{
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}