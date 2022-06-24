package storage

import (
	"github.com/hdn-project/calendar-service/storage/postgres"
	"github.com/hdn-project/calendar-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
    User() repo.UserStorageI
}

type storagePg struct {
    db         *sqlx.DB
    userRepo   repo.UserStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
    return &storagePg{
        db:         db,
        userRepo:   postgres.NewUserRepo(db),
    }
}

func (s storagePg) User() repo.UserStorageI {
    return s.userRepo
}
