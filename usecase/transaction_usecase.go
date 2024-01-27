package usecase

import (
	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/repository"
)

type TransactionUsecase interface {
	CreateTransaction(data entity.Transactions) (entity.Transactions, error)
}

type transactionUsecase struct {
	repo repository.TransactionRepo
}

// CreateTransaction implements TransactionUsecase.
func (t *transactionUsecase) CreateTransaction(data entity.Transactions) (entity.Transactions, error) {
	return t.repo.Create(data)
}

func NewTransactionUsecase(repo repository.TransactionRepo) TransactionUsecase {
	return &transactionUsecase{repo: repo}
}
