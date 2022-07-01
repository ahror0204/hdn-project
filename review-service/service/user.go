package service

import (
	"context"

	pb "github.com/hdn-project/review-service/genproto"
	l "github.com/hdn-project/review-service/pkg/logger"
	"github.com/hdn-project/review-service/storage"
	"github.com/jmoiron/sqlx"
)

//UserService ...
type ReviewService struct {
    storage storage.IStorage
    logger  l.Logger
}

//NewUserService ...
func NewReviewService(db *sqlx.DB, log l.Logger) *ReviewService {
    return &ReviewService{
        storage: storage.NewStoragePg(db),
        logger:  log,
    }
}

func (s *ReviewService) Create(ctx context.Context, req *pb.Review) (*pb.Review, error) {
    return nil, nil
}
