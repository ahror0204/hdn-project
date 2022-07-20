package storage

import (
	"github.com/hdn-project/user-service/storage/postgres"
	"github.com/hdn-project/user-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
	User() repo.UserStorageI
}

type storagePg struct {
	db       *sqlx.DB
	UserRepo repo.UserStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		UserRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.UserRepo
}
