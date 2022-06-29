package repo

import (
	pb "github.com/hdn-project/client-service/genproto"
)

//UserStorageI ...
type ClientStorageI interface {
    CreateUser(*pb.Client) (*pb.Empty, error)
	GetClientById(id string) (*pb.Client, error)
	DeleteById(id string) (*pb.Empty, error)
	UpdateClient(*pb.Client) (*pb.Empty, error)
}
