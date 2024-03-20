package storage

import (
	"taxi_service/storage/postgres"
	"taxi_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Taxi() repo.TaxiStorageI
}

type StoragePg struct {
	db       *sqlx.DB
	taxiRepo repo.TaxiStorageI
}

func NewStoragePg(db *sqlx.DB) *StoragePg {
	return &StoragePg{
		db:       db,
		taxiRepo: postgres.NewTaxiRepo(db),
	}
}

func (s StoragePg) Taxi() repo.TaxiStorageI {
	return s.taxiRepo
}
