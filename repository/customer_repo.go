package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/yaminmuhammad/pay-app/config"
	"github.com/yaminmuhammad/pay-app/entity"
	"golang.org/x/crypto/bcrypt"
)

type CustomerRepo interface {
	Register(data entity.Customer) (entity.Customer, error)
	Get(id string) (entity.Customer, error)
	GetCustomer(email string) (entity.Customer, error)
	LogActivity(activity entity.Activities) error
}

type customerRepo struct {
	db *sql.DB
}

// LogActivity implements CustomerRepo.
func (c *customerRepo) LogActivity(activity entity.Activities) error {
	_, err := c.db.Exec(
		config.InsertActivity,
		activity.CustomerId,
		activity.Activity,
		activity.ActivityTime,
	)
	return err
}

func (c *customerRepo) GetCustomer(email string) (entity.Customer, error) {
	var customer entity.Customer

	err := c.db.QueryRow(config.GetCustomerByEmail, email).Scan(
		&customer.Id,
		&customer.Username,
		&customer.Phone,
		&customer.Email,
		&customer.HashPassword,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Customer{}, fmt.Errorf("customer not found with email: %s", email)
		}
		log.Println("customerRepository :", err.Error())
		return entity.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepo) Get(id string) (entity.Customer, error) {
	var customer entity.Customer
	err := c.db.QueryRow(config.GetCustomerById, id).Scan(
		&customer.Id,
		&customer.Username,
		&customer.Phone,
		&customer.Email,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		log.Println("customerRepository.Get.QueryRow:", err.Error())
		return entity.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepo) Register(data entity.Customer) (entity.Customer, error) {
	var customer entity.Customer

	//	Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.HashPassword), bcrypt.DefaultCost)

	if err != nil {
		log.Println("customerRepo GenerateFromPassword:", err.Error())
		return entity.Customer{}, err
	}

	err = c.db.QueryRow(config.CreateCustomer,
		data.Username,
		data.Phone,
		data.Email,
		string(hashedPassword),
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
