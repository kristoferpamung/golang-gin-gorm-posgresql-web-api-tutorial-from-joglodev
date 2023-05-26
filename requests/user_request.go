package requests

import "time"

type UserRequest struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required"`
	Address  string    `json:"address" binding:"required"`
	BornDate time.Time `json:"born_date" binding:"required"`
}
