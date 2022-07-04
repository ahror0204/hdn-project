package repo

import (
	pb "github.com/hdn-project/business-service/genproto"
)

//BusinessStorageI ...
type BusinessStorageI interface {
	CreateBusiness(*pb.Business) (*pb.Business, error)
	CreateService(*pb.ServiceTypeDef) (*pb.ServiceTypeDef, error)
}
