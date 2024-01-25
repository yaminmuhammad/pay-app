package config

const (
	CreateCustomer = `INSERT INTO customers
    (username, phone, email, password, updated_at) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	GetCustomerById = `SELECT id, username, phone, email, created_at, updated_at
	FROM customers WHERE id = $1`

	GetCustomerByEmail = `SELECT id, username, phone, email
	FROM customers WHERE email = $1`
)
