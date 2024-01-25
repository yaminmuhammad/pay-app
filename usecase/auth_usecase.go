package usecase

import (
	"github.com/yaminmuhammad/pay-app/dto"
	"github.com/yaminmuhammad/pay-app/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequest) (dto.AuthResponse, error)
}

type authUseCase struct {
	customerUC CustomerUseCase
	jwtService service.JwtService
}

func (a *authUseCase) Login(payload dto.AuthRequest) (dto.AuthResponse, error) {
	customer, err := a.customerUC.AuthCustomer(payload.Email, payload.Password)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	token, err := a.jwtService.CreateToken(customer)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return token, nil
}

func NewAuthUseCase(customerUC CustomerUseCase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{customerUC: customerUC, jwtService: jwtService}
}
