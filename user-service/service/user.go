package service

import (
	"context"

	pb "github.com/hdn-project/user-service/genproto"
	l "github.com/hdn-project/user-service/pkg/logger"
	"github.com/hdn-project/user-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.Empty, error) {
	_, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.logger.Error("error while creating user", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating user")
	}
	return &pb.Empty{}, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	User, err := s.storage.User().GetUserById(req.Id)
	if err != nil {
		s.logger.Error("error while getting User", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while getting User")
	}
	return User, nil
}

func (s *UserService) DeleteById(ctx context.Context, req *pb.UserId) (*pb.Empty, error) {
	_, err := s.storage.User().DeleteById(req.Id)
	if err != nil {
		s.logger.Error("error while deleting User", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while deleting User")
	}
	return &pb.Empty{}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.Empty, error) {
	_, err := s.storage.User().UpdateUser(req)
	if err != nil {
		s.logger.Error("error while updating User", l.Error(err))
		return nil, status.Error(codes.Internal, "Error while updating User")
	}
	return &pb.Empty{}, nil
}
