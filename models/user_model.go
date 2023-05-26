package models

import "time"

type User struct {
	ID       *int       `json:"id"`
	Name     *string    `json:"name"`
	Address  *string    `json:"address"`
	Email    *string    `json:"email"`
	Password *string    `json:"password"`
	BornDate *time.Time `json:"born_date"`
}
