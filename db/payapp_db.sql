CREATE DATABASE payapp_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE activity_type AS ENUM('Login', 'Payment', 'Logout');
CREATE  TYPE status_type AS ENUM('Success', 'Process', 'Failed');

CREATE TABLE customers (
id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
username VARCHAR(255) NOT NULL,
phone VARCHAR(13) UNIQUE,
email VARCHAR(100) UNIQUE NOT NULL,
hash_password VARCHAR(100) NOT NULL,
token VARCHAR(255) DEFAULT NULL,
created_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE merchants (
id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
name VARCHAR(255) NOT NULL,
phone VARCHAR(13) UNIQUE,
created_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE activities (
id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
customer_id uuid,
activity activity_type DEFAULT NULL,
activity_time TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

CREATE TABLE transactions (
id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
customer_id uuid,
merchant_id uuid,
product VARCHAR(255) DEFAULT NULL,
amount DECIMAL(10, 2),
status status_type DEFAULT NULL,
code VARCHAR(255) DEFAULT NULL,
transaction_time TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (customer_id) REFERENCES customers(id),
FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);


