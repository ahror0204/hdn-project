package storage

import (
	"github.com/hdn-project/review-service/storage/postgres"
	"github.com/hdn-project/review-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
    Review() repo.ReviewStorageI
}

type storagePg struct {
    db         *sqlx.DB
    reviewRepo   repo.ReviewStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
    return &storagePg{
        db:         db,
        reviewRepo:   postgres.NewReviewRepo(db),
    }
}

func (s storagePg) Review() repo.ReviewStorageI {
    return s.reviewRepo
}
