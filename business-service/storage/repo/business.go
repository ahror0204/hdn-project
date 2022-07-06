package repo

import (
	pb "github.com/hdn-project/business-service/genproto"
)

//BusinessStorageI ...
type BusinessStorageI interface {
	CreateBusiness(*pb.Business) (*pb.Business, error)
	UpdateBusiness(*pb.Business) (*pb.Empty, error)
	DeleteBusiness(*pb.Id) (*pb.Empty, error)
	GetByIdBusiness(*pb.Id) (*pb.Business, error)
	GetAllBusiness(*pb.Empty) (*pb.GetAllBusinessResponse, error)
	GetListBusiness(limit int64, page int64) (*pb.GetAllBusinessResponse, error)
	//----------------------------Services-----------------------------------
	CreateService(*pb.User) (*pb.ServiceTypeDef, error)
}
