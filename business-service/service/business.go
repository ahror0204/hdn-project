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
	return &pb.Empty{}, nil
}

func (s *BusinessService) DeleteBusiness(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	_, err := s.storage.Business().DeleteBusiness(req)
	if err != nil {
		s.logger.Error("failed while deleting business", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting business")
	}
	return &pb.Empty{}, nil
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

func (s *BusinessService) CreateService(ctx context.Context, req *pb.User) (*pb.ServiceTypeDef, error) {

	menService, womenService, err := s.storage.Business().CreateService(req)
	if err != nil {
		s.logger.Error("failed while creating service type", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating service type")
	}

	return &pb.ServiceTypeDef{
		MenService:   menService,
		WomenService: womenService,
	}, nil
}

func (s *BusinessService) UpdateMenServiceByUserId (ctx context.Context, req *pb.MenServices) (*pb.Empty, error) {

	_, err := s.storage.Business().UpdateMenServiceByUserId(req)
	if err != nil {
		s.logger.Error("failed while updating men Service", l.Error(err))
		return nil, status.Error(codes.Internal, "failed whilev updating men Service")
	}

	return &pb.Empty{}, nil
}

func (s *BusinessService) UpdateWomenServiceByUserId (ctx context.Context, req *pb.WomenServices) (*pb.Empty, error) {

	_, err := s.storage.Business().UpdateWomenServiceByUserId(req)
	if err != nil {
		s.logger.Error("failed while updating women Service", l.Error(err))
		return nil, status.Error(codes.Internal, "failed whilev updating women Service")
	}

	return &pb.Empty{}, nil
}

func (s *BusinessService) GetMenServiceByUserId (ctx context.Context, req *pb.Id) (*pb.MenServices, error) {
	menservice, err := s.storage.Business().GetMenServiceByUserId(req)
	if err != nil {
		s.logger.Error("failed while getting men Service by user id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting men Service by user id")
	}

	return menservice, nil
}

func (s *BusinessService) GetWomenServiceByUserId (ctx context.Context, req *pb.Id) (*pb.WomenServices, error) {
	womenservice, err := s.storage.Business().GetWomenServiceByUserId(req)
	if err != nil {
		s.logger.Error("failed while getting women Service by user id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting women Service by user id")
	}

	return womenservice, nil
}

func (s *BusinessService) DeleteMenServiceByUserId (ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	_, err := s.storage.Business().DeleteMenServiceByUserId(req)
	if err != nil {
		s.logger.Error("failed while deleting men Service by user id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting men Service by user id")
	}
	return &pb.Empty{}, nil
}

func (s *BusinessService) DeleteWomenServiceByUserId (ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	_, err := s.storage.Business().DeleteWomenServiceByUserId(req)
	if err != nil {
		s.logger.Error("failed while deleting women Service by user id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting women Service by user id")
	}
	return &pb.Empty{}, nil
}

func (s *BusinessService) GetAllMenSetvices (ctx context.Context, req *pb.Empty) (*pb.AllMenSetvices, error) {
	allMenServices, err := s.storage.Business().GetAllMenSetvices(req)
	if err != nil {
		s.logger.Error("failed while getting all men services", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting all men services")
	}
	return &pb.AllMenSetvices{
		MenServices: allMenServices,
	}, nil
} 

func (s *BusinessService) GetAllWomenSetvices (ctx context.Context, req *pb.Empty) (*pb.AllWomenSetvices, error) {
	allWomenServices, err := s.storage.Business().GetAllWomenSetvices(req)
	if err != nil {
		s.logger.Error("failed while getting all women services", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting all women services")
	}
	return &pb.AllWomenSetvices{
		WomenServices: allWomenServices,
	}, nil
} 