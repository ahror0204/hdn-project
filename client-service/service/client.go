package service

import (
	"context"

	pb "github.com/hdn-project/client-service/genproto"
	l "github.com/hdn-project/client-service/pkg/logger"
	"github.com/hdn-project/client-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//ClientService ...
type ClientService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewClientService ...
func NewClientService(db *sqlx.DB, log l.Logger) *ClientService {
	return &ClientService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *ClientService) CreateUser(ctx context.Context, req *pb.Client) (*pb.Empty, error) {
	_, err := s.storage.Client().CreateUser(req)
	if err != nil {
		s.logger.Error("error while creating user", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating user")
	}
	return &pb.Empty{}, nil
}

func (s *ClientService) GetClientById(ctx context.Context, req *pb.ClientId) (*pb.Client, error) {
	client, err := s.storage.Client().GetClientById(req.Id)
	if err != nil {
		s.logger.Error("error while getting client", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting client")
	}
	return client, nil
}

func (s *ClientService) DeleteById(ctx context.Context, req *pb.ClientId) (*pb.Empty, error) {
	_, err := s.storage.Client().DeleteById(req.Id)
	if err != nil {
		s.logger.Error("error while deleting client", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while deleting client")
	}
	return &pb.Empty{}, nil
}

func (s *ClientService) UpdateClient(ctx context.Context, req *pb.Client) (*pb.Empty, error) {
	_, err := s.storage.Client().UpdateClient(req)
	if err != nil {
		s.logger.Error("error while updating client", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while updating client")
	}
	return &pb.Empty{}, nil
}

