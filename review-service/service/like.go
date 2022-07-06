package service

import (
	"context"

	pb "github.com/hdn-project/review-service/genproto"
	l "github.com/hdn-project/review-service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ReviewService) CreateLike(ctx context.Context, req *pb.Like) (*pb.Empty, error) {
    err := s.storage.Like().Create(req)
    if err != nil {
        s.logger.Error("failed to create like", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
    }
    return &pb.Empty{}, nil

}

func (s *ReviewService) DeleteUserLikes(ctx context.Context, req *pb.DeleteUserLikesRequest) (*pb.Empty, error) {
	err := s.storage.Like().DeleteUserLikes(req.UserId)
	if err != nil {
		s.logger.Error("failed to delete like", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.Empty{}, nil
}

func (s *ReviewService) DeleteLike(ctx context.Context, req *pb.Like) (*pb.Empty, error) {
	err := s.storage.Like().DeleteLike(req)
	if err != nil {
		s.logger.Error("failed to delete like", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.Empty{}, nil
}

func (s *ReviewService) CountLikes(ctx context.Context, req *pb.Like) (*pb.CountLikesResponse, error) {
	count, exists, err := s.storage.Like().CountLikes(req)
	if err != nil {
		s.logger.Error("failed to count like", l.Error(err), l.Any("req", req))
		return nil,status.Error(codes.Internal, "Internal server error")
	}
	return &pb.CountLikesResponse{
		Count: count,
		Liked: exists,
		}, nil
}

func (s *ReviewService) ListLikeUsers(ctx context.Context, req *pb.ListLikeUsersRequest) (*pb.ListLikeUsersResponse, error) {
	users, count, err := s.storage.Like().ListLikeUsers(req.StaffId, req.UserId, req.Limit, req.Page)
	if err != nil {
		s.logger.Error("failed to list like users", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.ListLikeUsersResponse{
		Users: users,
		Count: count,
		}, nil
}