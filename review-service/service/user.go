package service

import (
	"context"

	pb "github.com/hdn-project/review-service/genproto"
	l "github.com/hdn-project/review-service/pkg/logger"
	"github.com/hdn-project/review-service/storage"
	"github.com/jmoiron/sqlx"
)

//UserService ...
type UserService struct {
    storage storage.IStorage
    logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
    return &UserService{
        storage: storage.NewStoragePg(db),
        logger:  log,
    }
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
    return nil, nil
}
