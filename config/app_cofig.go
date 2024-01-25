package config

const (
	ApiGroup = "/api/v1"

	CreateCustomer = `INSERT INTO customers
    (username, phone, email, password, updated_at) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
)
