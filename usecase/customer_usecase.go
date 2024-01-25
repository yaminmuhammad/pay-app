package usecase

import (
	"errors"
	"fmt"
	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type CustomerUseCase interface {
	RegisterCustomer(data entity.Customer) (entity.Customer, error)
	FindCustomerByID(id string) (entity.Customer, error)
	AuthCustomer(email string, password string) (entity.Customer, error)
}

type customerUsecase struct {
	repo repository.CustomerRepo
}

func (c *customerUsecase) AuthCustomer(email string, password string) (entity.Customer, error) {
	customer, err := c.repo.GetUser(email)
	if err != nil {
		return entity.Customer{}, err
	}
	if password == "" {
		log.Println("customerUsecase.AuthCustomer: Empty password provided")
		return entity.Customer{}, errors.New("password required")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		log.Println("customerUsecase.AuthCustomer: Password verification failed")
		return entity.Customer{}, errors.New("password verification failed")
	}

	return customer, nil
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
