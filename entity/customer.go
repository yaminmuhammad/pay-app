package entity

import "time"

type Customer struct {
	Id           string    `json:"id"`
	Username     string    `json:"username" binding:"required"`
	Phone        string    `json:"phone" binding:"required"`
	Email        string    `json:"email" binding:"required,email"`
	HashPassword string    `json:"hashPassword" binding:"required"`
	Token        string    `json:"token"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
