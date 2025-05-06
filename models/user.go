package models

type User struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Email       string `json:"email"  binding:"required" db:"email"`
	Password    string `json:"password"  binding:"required"`
	Subsription int    `json:"subscription" db:"sub"`
	UserType    string `json:"user_type" db:"user_type"`
}
type AuthUser struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
