CREATE DATABASE payapp_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE customers (
                           id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                           username VARCHAR(255) NOT NULL,
                           phone VARCHAR(13) UNIQUE,
                           email VARCHAR(100) UNIQUE,
                           password VARCHAR(100),
                           created_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE merchants (
                           id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           phone VARCHAR(13) UNIQUE,
                           email VARCHAR(100) UNIQUE,
                           password VARCHAR(100),
                           created_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
                              id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                              customer_id uuid,
                              merchant_id uuid,
                              amount DECIMAL(10, 2),
                              transaction_type VARCHAR(10),
                              transaction_time TIMESTAMPTZ(0) DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (customer_id) REFERENCES customers(id),
                              FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);


INSERT INTO
    users(name, email, username, address, hash_password, role)
VALUES
    (
        'Iqi Tes',
        'iqi@mail.com',
        'iqi',
        'Cirebon',
        'password',
        'participant'
    );

INSERT INTO
    participants(
    date_of_birth,
    place_of_birth,
    last_education,
    user_id,
    role
)
VALUES
    (
        '1999-10-10',
        'Jakarta',
        'Universitas Gadjah Mada',
        '37881fb4-5ca4-4939-a1ad-d8a36fbb23a6',
        'Advance'
    );

INSERT INTO
    schedules(
    activity,
    date,
    trainer_id,
    participant_id
)
VALUES
    (
        'Training',
        '2023-12-23',
        'e0da3c80-17b4-41df-8087-5e2fbc19f654',
        'c112bc5c-29c8-4921-a85a-505b21b97b1d'
    );
