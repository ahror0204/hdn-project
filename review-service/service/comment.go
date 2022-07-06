package service

import (
	"context"

	pb "github.com/hdn-project/review-service/genproto"
	l "github.com/hdn-project/review-service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ReviewService) CreateComment(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	com,err := s.storage.Comment().CreateComment(req)
	if err != nil {
		s.logger.Error("failed to create comment", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return com, nil
}

func (s *ReviewService) UpdateComment(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	com,err := s.storage.Comment().UpdateComment(req)
	if err != nil {
		s.logger.Error("failed to update comment", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return com, nil
}

func (s *ReviewService) DeleteComment(ctx context.Context, req *pb.DeleteUserCommentsRequest) (*pb.Empty, error) {
	err := s.storage.Comment().DeleteComment(req.Id, req.UserId)
	if err != nil {
		s.logger.Error("failed to delete comment", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.Empty{}, nil
}

func (s *ReviewService) DeleteUserComments(ctx context.Context, req *pb.DeleteUserCommentsRequest) (*pb.Empty, error) {
	err := s.storage.Comment().DeleteUserComments(req.UserId)
	if err != nil {
		s.logger.Error("failed to delete user comments", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.Empty{}, nil
}

func (s *ReviewService) CountComments(ctx context.Context, req *pb.Comment) (*pb.CountCommentsResponse, error) {
	count, err := s.storage.Comment().CountComments(req.StaffId)
	if err != nil {
		s.logger.Error("failed to count comments", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.CountCommentsResponse{
		Count: count,
		}, nil
}


func (s *ReviewService) ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	//default id or fake id
	var id string
	comments, count, err := s.storage.Comment().ListComments(id,  req.Limit, req.Page)
	if err != nil {
		s.logger.Error("failed to list comments", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.ListCommentsResponse{
		Results: comments,
		Count: count,
		}, nil
}