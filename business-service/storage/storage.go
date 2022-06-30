package storage

import (
	"github.com/hdn-project/business-service/storage/postgres"
	"github.com/hdn-project/business-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
	User() repo.BusinessStorageI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.BusinessStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewBusinessRepo(db),
	}
}

func (s storagePg) User() repo.BusinessStorageI {
	return s.userRepo
}
