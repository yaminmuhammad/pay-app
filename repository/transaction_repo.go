package repository

import (
	"database/sql"
	"time"

	"github.com/yaminmuhammad/pay-app/config"
	"github.com/yaminmuhammad/pay-app/entity"
)

type TransactionRepo interface {
	Create(data entity.Transactions) (entity.Transactions, error)
}

type transactionRepo struct {
	db *sql.DB
}

// Create implements TransactionRepo.
func (t *transactionRepo) Create(data entity.Transactions) (entity.Transactions, error) {
	var transaction entity.Transactions

	tx, err := t.db.Begin()
	if err != nil {
		return entity.Transactions{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.QueryRow(config.CreateTransaction,
		data.CustomerId,
		data.MerchantId,
		data.Product,
		data.Amount,
		data.Status,
		data.Code,
		time.Now(),
	).Scan(&transaction.Id)

	if err != nil {
		return entity.Transactions{}, err
	}

	transaction.CustomerId = data.CustomerId
	transaction.MerchantId = data.MerchantId
	transaction.Product = data.Product
	transaction.Amount = data.Amount
	transaction.Status = data.Status
	transaction.Code = data.Code
	transaction.TransactionTime = time.Now()

	_, err = tx.Exec(config.InsertActivity,
		data.CustomerId,
		"Payment",
		time.Now(),
	)
	if err != nil {
		return entity.Transactions{}, err
	}

	return transaction, nil
}

func NewTransactionRepo(db *sql.DB) TransactionRepo {
	return &transactionRepo{db: db}
}
