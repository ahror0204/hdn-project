package repo

import (
	pb "github.com/hdn-project/review-service/genproto"
)

//UserStorageI ...
type ReviewStorageI interface {
    Create(*pb.Review) (*pb.Review, error)
}
