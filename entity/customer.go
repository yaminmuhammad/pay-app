package entity

import "time"

type Customer struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	HashPassword string    `json:"hashPassword"`
	Token        string    `json:"token"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
