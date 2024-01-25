package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	jwt.RegisteredClaims
	CustomerId string `json:"customerId"`
}
