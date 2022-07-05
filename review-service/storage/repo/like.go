package repo

import (
	pb "github.com/hdn-project/review-service/genproto"
)

//LikeStorageI ...
type LikeStorageI interface {
	Create(*pb.Like) error
	DeleteUserLikes(userId string) error
	DeleteLike(*pb.Like) error
	CountLikes(*pb.Like) (int64, bool, error)
	ListLikeUsers(staffId,userId string, limit, page int64) ([]*pb.LikedUser,int64, error)
}
