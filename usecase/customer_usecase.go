package usecase

import (
	"fmt"
	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/repository"
	"time"
)

type CustomerUseCase interface {
	RegisterCustomer(data entity.Customer) (entity.Customer, error)
	FindCustomerByID(id string) (entity.Customer, error)
}

type customerUsecase struct {
	repo repository.CustomerRepo
}

func (c *customerUsecase) FindCustomerByID(id string) (entity.Customer, error) {
	return c.repo.Get(id)
}

func (c *customerUsecase) RegisterCustomer(data entity.Customer) (entity.Customer, error) {
	if data.Username == "" || data.Phone == "" || data.Email == "" || data.Password == "" {
		return entity.Customer{}, fmt.Errorf("Oops, all fields must be filled")
	}
	data.UpdatedAt = time.Now()
	customer, err := c.repo.Register(data)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("Failed to save data customer :%v", err.Error())
	}

	return customer, nil
}

func NewCustomerUseCase(repo repository.CustomerRepo) CustomerUseCase {
	return &customerUsecase{repo: repo}
}
