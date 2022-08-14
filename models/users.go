package models

type Users struct {
	Code int `json:"code"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Role string `json:"role"`
	LastUpdated string `json:"last_updated"`  
}