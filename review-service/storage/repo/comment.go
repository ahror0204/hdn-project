package repo

import (
	pb "github.com/hdn-project/review-service/genproto"
)

//CommentStorageI ...
type CommentStorageI interface {
	CreateComment(*pb.Comment) (*pb.Comment, error)
	UpdateComment(*pb.Comment) (*pb.Comment, error)
	DeleteComment(id, staffId string) error
	DeleteUserComments(userId string) error
	CountComments(staffId string) (int64, error)
	ListComments(staffId string, limit, page int64) ([]*pb.Comment, int64, error)

}
