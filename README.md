
# Golang PayApp

## About the project

This is a simple project API to connect between merchant customers and banks. with payment login and logout features. and customer activities are recorded in history.

### API docs

This project does not provide a complete api spec but there is a postman collection that can be tried. the file is located in the postman folder in the project
<!-- Getting Started -->
# Getting Started

<!-- Run Locally -->
### Database
To create a table, the query has been provided in the db folder in the project with the file name is <b>payapp_db.sql</b>.
<!-- Run Locally -->
### Run Locally

Clone the project

```bash
  git clone https://github.com/yaminmuhammad/pay-app.git
```

Go to the project directory

```bash
  cd pay-app
```
<!-- Env Variables -->
### Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_HOST`
`DB_PORT`
`DB_USER`
`DB_PASSWORD`
`DB_NAME`
`DB_DRIVER`
`API_PORT`
`TOKEN_ISSUE`
`TOKEN_SECRET`
`TOKEN_EXPIRE`


Start the server

```bash
  go run . 
```
or
```bash
  go run main.go 
```


<!-- Deployment -->
### Test Postman

To test the project in postman, you need to import the postman json file collection in the postman folder of this project first.
``

## Todo

* [x] Register customer
* [x] Register Merchant
* [x] Login Customer with jwt
* [x] Transaction / Payment
* [x] Record activity Login on table activity (automatic)
* [x] Record activity Payment on table activity (automatic)
* [ ] Record activity Logout on table activity (automatic)


### Layout

```tree
├── .idea
├── .gitignore
├── .env.example
├── go.mod
├── go.sum
├── main.go
├── config
│   ├── app_config.go
│   ├── config.go
│   └── raw_query.go
├── db
│   └── payapp_db.sql
├── dto
│   ├── auth_dto.go
|
├── entity
│   ├── activity.go
│   ├── customer.go
│   └── merchant.go
│   └── transaction.go
|
├── handler
│   ├── controller
│   │   |── auth_controller.go
│   |   |── customer_controller.go
│   |   |── merchant_controller.go
│   |   |── transaction_controller.go
│   |   
│   ├── middleware
│   │   └── auth_middleware.go
│   └── server.go
├── postman
│   |
│   └── Pay-app.postman_collection.json
├── repository
│   ├── customer_repo.go
│   ├── merchant_repo.go
│   └── transaction_repo.go
├── shared
│   ├── common
│   |     └── response.go
│   ├── model
│   │   ├── custom_claims.go
│   │   |
│   │   └── json_model.go
│   └── service
│       └── jwt_service.go
└── usecase
    ├── auth_usecase.go
    ├── customer_usecase.go
    ├── merchant_usecase.go
    └── transaction_usecase.go
```