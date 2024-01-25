package entity

import "time"

type Transactions struct {
	Id              string    `json:"id"`
	CustomerId      string    `json:"customerId"`
	MerchantId      string    `json:"merchantId"`
	Product         string    `json:"product"`
	Amount          float64   `json:"amount"`
	Status          string    `json:"status"`
	Code            string    `json:"code"`
	TransactionTime time.Time `json:"transactionTime"`
}
