package storage

import (
	"github.com/hdn-project/review-service/storage/postgres"
	"github.com/hdn-project/review-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
    Like() repo.LikeStorageI
    Comment() repo.CommentStorageI
}

type storagePg struct {
    db         *sqlx.DB
    likeRepo   repo.LikeStorageI
    commentRepo repo.CommentStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
    return &storagePg{
        db:         db,
        likeRepo:   postgres.NewLikeRepo(db),
        commentRepo: postgres.NewCommentRepo(db),
    }
}

func (s storagePg) Like() repo.LikeStorageI {
    return s.likeRepo
}

func (s storagePg) Comment() repo.CommentStorageI {
    return s.commentRepo
}
