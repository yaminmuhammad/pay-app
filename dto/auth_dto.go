package dto

type AuthRequest struct {
	Email        string `json:"email" binding:"required,email"`
	HashPassword string `json:"hashPassword" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
