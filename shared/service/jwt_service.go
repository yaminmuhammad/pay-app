package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yaminmuhammad/pay-app/config"
	"github.com/yaminmuhammad/pay-app/dto"
	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/shared/model"
	"time"
)

type JwtService interface {
	CreateToken(customer entity.Customer) (dto.AuthResponse, error)
	ParseToken(tokenHeader string) (jwt.MapClaims, error)
}

type jwtService struct {
	config config.TokenConfig
}

func (j *jwtService) CreateToken(customer entity.Customer) (dto.AuthResponse, error) {
	claims := model.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.IssuerName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.JwtExpiresTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		CustomerId: customer.Id,
	}

	token := jwt.NewWithClaims(j.config.JwtSigningMethod, claims)
	ss, err := token.SignedString(j.config.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("Failed to Create Token :%v", err)
	}
	return dto.AuthResponse{Token: ss}, nil
}

func (j *jwtService) ParseToken(tokenHeader string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("oops, unexpected signing method: %v", token.Header["alg"])
		}
		return j.config.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("oops, failed to verify token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("oops, failed to parse token claims")
	}
	return claims, nil
}

func NewJwtService(config config.TokenConfig) JwtService {
	return &jwtService{config: config}
}
