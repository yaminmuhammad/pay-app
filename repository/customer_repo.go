package repository

import (
	"database/sql"
	"github.com/yaminmuhammad/pay-app/config"
	"github.com/yaminmuhammad/pay-app/entity"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type CustomerRepo interface {
	Save(data entity.Customer) (entity.Customer, error)
}

type customerRepo struct {
	db *sql.DB
}

func (c *customerRepo) Save(data entity.Customer) (entity.Customer, error) {
	var customer entity.Customer

	//	Hass Password
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("customerRepo GenerateFromPassword:", err.Error())
		return entity.Customer{}, err
	}

	err = c.db.QueryRow(config.CreateCustomer,
		data.Username,
		data.Phone,
		data.Email,
		string(password),
		time.Now()).Scan(
		&customer.Id,
		&customer.CreatedAt,
	)
	if err != nil {
		log.Println("customerRepo QueryRow", err.Error())
		return entity.Customer{}, err
	}
	customer.Username = data.Username
	customer.Phone = data.Phone
	customer.Email = data.Email
	customer.UpdatedAt = data.UpdatedAt

	return customer, nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepo{db: db}
}
