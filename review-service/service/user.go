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

func (s *ReviewService) CreateComment(ctx context.Context, req *pb.Comment) (*pb.Empty, error) {
    _, err := s.storage.Review().CreateComment(req)
    if err != nil {
        s.logger.Error("Error while creating review", l.Error(err))
        return nil, err
    }
    return &pb.Empty{}, nil
}


func (s *ReviewService) CreateReply(ctx context.Context, req *pb.ReplyComments) (*pb.Empty, error) {
    _, err := s.storage.Review().CreateReply(req)
    if err != nil {
        s.logger.Error("Error while creating review", l.Error(err))
        return nil, err
    }
    return &pb.Empty{}, nil
}
