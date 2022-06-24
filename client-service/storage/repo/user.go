package repo

import (
	pb "github.com/hdn-project/client-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
    Create(*pb.User) (*pb.User, error)
}
