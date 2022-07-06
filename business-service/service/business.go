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

func (s *BusinessService) UpdateBusiness(ctx context.Context, req *pb.Business) (*pb.Empty, error) {

	_, err := s.storage.Business().UpdateBusiness(req)
	if err != nil {
		s.logger.Error("failed while updating business", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating business")
	}
	return nil, nil
}

func (s *BusinessService) DeleteBusiness(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	_, err := s.storage.Business().DeleteBusiness(req)
	if err != nil {
		s.logger.Error("failed while deleting business", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting business")
	}
	return nil, nil
}

func (s *BusinessService) GetByIdBusiness(ctx context.Context, req *pb.Id) (*pb.Business, error) {
	byIdBusiness, err := s.storage.Business().GetByIdBusiness(req)
	if err != nil {
		s.logger.Error("failed while getting business by id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting business by id")
	}

	return byIdBusiness, nil
}

func (s *BusinessService) GetAllBusiness(ctx context.Context, req *pb.Empty) (*pb.GetAllBusinessResponse, error) {
	allBusinesses, err := s.storage.Business().GetAllBusiness(req)
	if err != nil {
		s.logger.Error("failed while getting all business", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting all business")
	}
	return allBusinesses, nil
}

func (s *BusinessService) GetListBusiness(ctx context.Context, req *pb.GetListBusinessRequest) (*pb.GetAllBusinessResponse, error) {
	listBusinesses, err := s.storage.Business().GetListBusiness(req.Limit, req.Page)
	if err != nil {
		s.logger.Error("failed while getting list of business", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting list of business")
	}
	return listBusinesses, nil
}
//-------------------------------------------------Services--------------------------------------------------------


func (s *BusinessService) CreateService (ctx context.Context, req *pb.User) (*pb.ServiceTypeDef, error) {
	
	service, err := s.storage.Business().CreateService(req)
	if err != nil {
		s.logger.Error("failed while defining service type", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while defining service type")
	}

	return service, nil
}