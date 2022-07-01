package postgres

import (
	pb "github.com/hdn-project/business-service/genproto"
	"github.com/jmoiron/sqlx"
)

type businessRepo struct {
	db *sqlx.DB
}

//NewBusinessRepo ...
func NewBusinessRepo(db *sqlx.DB) *businessRepo {
	return &businessRepo{db: db}
}

func (r *businessRepo) CreateBusiness(business *pb.Business) (*pb.Business, error) {

	busQuery := `INSERT INT`

	return nil, nil
}
