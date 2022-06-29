package storage

import (
	"github.com/hdn-project/client-service/storage/postgres"
	"github.com/hdn-project/client-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
    Client() repo.ClientStorageI
}

type storagePg struct {
    db         *sqlx.DB
    clientRepo   repo.ClientStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
    return &storagePg{
        db:         db,
        clientRepo:   postgres.NewClientRepo(db),
    }
}

func (s storagePg) Client() repo.ClientStorageI {
    return s.clientRepo
}
