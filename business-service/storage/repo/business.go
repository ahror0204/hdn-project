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
	CreateService(*pb.User) (*pb.MenServices, *pb.WomenServices, error)
	UpdateMenServiceByUserId(*pb.MenServices) (pb.Empty, error)
	UpdateWomenServiceByUserId(*pb.WomenServices) (pb.Empty, error)
	GetMenServiceByUserId(*pb.Id) (*pb.MenServices, error)
	GetWomenServiceByUserId(*pb.Id) (*pb.WomenServices, error)
	DeleteMenServiceByUserId(*pb.Id) (*pb.Empty, error)
	DeleteWomenServiceByUserId(*pb.Id) (*pb.Empty, error)
	GetAllMenSetvices(*pb.Empty) ([]*pb.MenServices, error)
	GetAllWomenSetvices(*pb.Empty) ([]*pb.WomenServices, error)
}