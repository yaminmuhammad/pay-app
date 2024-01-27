package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/yaminmuhammad/pay-app/config"
	"github.com/yaminmuhammad/pay-app/entity"
)

type MerchantRepo interface {
	Register(data entity.Merchant) (entity.Merchant, error)
}

type merchantRepo struct {
	db *sql.DB
}

// Register implements MerchantRepo.
func (m *merchantRepo) Register(data entity.Merchant) (entity.Merchant, error) {
	var merchant entity.Merchant
	err := m.db.QueryRow(config.CreateMerchant,
		data.Name,
		data.Phone,
		time.Now()).Scan(
		&merchant.Id,
		&merchant.CreatedAt,
	)
	if err != nil {
		log.Println("merchantRepository :", err.Error())
		return entity.Merchant{}, err
	}
	merchant.Name = data.Name
	merchant.Phone = data.Phone
	merchant.UpdatedAt = data.UpdatedAt

	return merchant, nil
}

func NewMerchantRepo(db *sql.DB) MerchantRepo {
	return &merchantRepo{db: db}
}
