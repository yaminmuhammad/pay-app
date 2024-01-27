package config

const (
	CreateCustomer = `INSERT INTO customers (username, phone, email, hash_password, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	CreateMerchant = `INSERT INTO merchants (name, phone, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at`

	GetCustomerById = `SELECT id, username, phone, email, created_at, updated_at
	FROM customers WHERE id = $1`

	GetCustomerByEmail = `SELECT id, username, phone, email, hash_password FROM customers WHERE email = $1`

	InsertActivity = `INSERT INTO activities (customer_id, activity, activity_time) VALUES ($1, $2, $3)`
)
