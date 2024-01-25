package dto

type AuthRequest struct {
	Email        string `json:"email"`
	HashPassword string `json:"hashPassword"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
