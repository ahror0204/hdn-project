package repo

import (
	pb "github.com/hdn-project/User-service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreateUser(*pb.User) (*pb.Empty, error)
	GetUserById(id string) (*pb.User, error)
	DeleteById(id string) (*pb.Empty, error)
	UpdateUser(*pb.User) (*pb.Empty, error)
}
