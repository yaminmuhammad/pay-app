package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yaminmuhammad/pay-app/dto"
	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/repository"
	"github.com/yaminmuhammad/pay-app/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequest, ctx *gin.Context) (dto.AuthResponse, error)
	LogActivity(activity entity.Activities) error
}

type authUseCase struct {
	customerUC   CustomerUseCase
	customerRepo repository.CustomerRepo
	jwtService   service.JwtService
}

// LogActivity implements AuthUseCase.
func (a *authUseCase) LogActivity(activity entity.Activities) error {
	err := a.customerRepo.LogActivity(activity)
	if err != nil {
		// Penanganan kesalahan saat menyimpan aktivitas
		return fmt.Errorf("failed to log activity: %v", err)
	}
	return nil
}

func (a *authUseCase) Login(payload dto.AuthRequest, ctx *gin.Context) (dto.AuthResponse, error) {
	customer, err := a.customerUC.AuthCustomer(payload.Email, payload.HashPassword)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	token, err := a.jwtService.CreateToken(customer, ctx)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return token, nil
}

func NewAuthUseCase(customerUC CustomerUseCase, jwtService service.JwtService, customerRepo repository.CustomerRepo) AuthUseCase {
	return &authUseCase{customerUC: customerUC, jwtService: jwtService, customerRepo: customerRepo}
}
