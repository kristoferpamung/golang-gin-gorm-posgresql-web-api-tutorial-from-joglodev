package responses

import "time"

type UserResponse struct {
	Name     *string    `json:"name"`
	Address  *string    `json:"address"`
	Email    *string    `json:"email"`
	Password *string    `json:"password"`
	BornDate *time.Time `json:"born_date"`
}
