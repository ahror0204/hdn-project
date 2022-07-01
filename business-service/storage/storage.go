package storage

import (
	"github.com/hdn-project/business-service/storage/postgres"
	"github.com/hdn-project/business-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
	Business() repo.BusinessStorageI
}

type storagePg struct {
	db           *sqlx.DB
	businessRepo repo.BusinessStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:           db,
		businessRepo: postgres.NewBusinessRepo(db),
	}
}

func (s storagePg) Business() repo.BusinessStorageI {
	return s.businessRepo
}
