package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yaminmuhammad/pay-app/config"
	"github.com/yaminmuhammad/pay-app/dto"
	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/shared/model"
)

type JwtService interface {
	CreateToken(customer entity.Customer) (dto.AuthResponse, error)
	ParseToken(tokenHeader string) (jwt.MapClaims, error)
	// SaveCustomerSession(ctx *gin.Context, claims jwt.MapClaims)
}

type jwtService struct {
	config config.TokenConfig
}

// SaveCustomerSession implements JwtService.
// func (j *jwtService) SaveCustomerSession(ctx *gin.Context, claims jwt.MapClaims) {
// 	session := sessions.Default(ctx)
// 	session.Set("customerId", claims["customerId"])
// 	session.Save()
// }

func (j *jwtService) CreateToken(customer entity.Customer) (dto.AuthResponse, error) {
	claims := model.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.IssuerName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.JwtExpiresTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		CustomerId: customer.Id,
	}

	// session := sessions.Default(ctx)
	// session.Set("customerId", claims.CustomerId)
	// session.Save()

	token := jwt.NewWithClaims(j.config.JwtSigningMethod, claims)
	ss, err := token.SignedString(j.config.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("Failed to Create Token :%v", err)
	}
	return dto.AuthResponse{Token: ss}, nil
}

func (j *jwtService) ParseToken(tokenHeader string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		return j.config.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("oops, failed to verify token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("oops, failed to claim token")
	}
	return claims, nil
}

func NewJwtService(config config.TokenConfig) JwtService {
	return &jwtService{config: config}
}
