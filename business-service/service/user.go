package service

import (
	"context"

	pb "github.com/hdn-project/business-service/genproto"
	l "github.com/hdn-project/business-service/pkg/logger"
	"github.com/hdn-project/business-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//BusinessService ...
type BusinessService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewBusinessService ...
func NewBusinessService(db *sqlx.DB, log l.Logger) *BusinessService {
	return &BusinessService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *BusinessService) CreateBusiness(ctx context.Context, req *pb.Business) (*pb.Business, error) {
	
	business, err := s.storage.Business().CreateBusiness(req)
	if err != nil {
		s.logger.Error("failed while creating business", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating business")
	}

	return business, nil
}
