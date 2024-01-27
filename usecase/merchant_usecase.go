package usecase

import (
	"errors"
	"time"

	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/repository"
)

type MerchantUseCase interface {
	RegisterMerchant(data entity.Merchant) (entity.Merchant, error)
}

type merchantUsecase struct {
	repo repository.MerchantRepo
}

// RegisterMerchant implements MerchantUseCase.
func (m *merchantUsecase) RegisterMerchant(data entity.Merchant) (entity.Merchant, error) {
	if data.Name == "" || data.Phone == "" {
		return entity.Merchant{}, errors.New("name required")

	}
	data.UpdatedAt = time.Now()
	merchant, err := m.repo.Register(data)
	if err != nil {
		return entity.Merchant{}, errors.New("failed to register merchant")
	}

	return merchant, nil
}

func NewMerchantUseCase(repo repository.MerchantRepo) MerchantUseCase {
	return &merchantUsecase{repo: repo}
}
