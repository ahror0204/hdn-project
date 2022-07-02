package repo

import (
	pb "github.com/hdn-project/review-service/genproto"
)

//UserStorageI ...
type ReviewStorageI interface {
    CreateComment(*pb.Comment) (*pb.Empty, error)
	CreateReply(*pb.ReplyComments) (*pb.Empty, error)
}
