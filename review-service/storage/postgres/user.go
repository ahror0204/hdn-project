package postgres

import (
	pb "github.com/hdn-project/review-service/genproto"
	"github.com/jmoiron/sqlx"
)

type reviewRepo struct {
    db *sqlx.DB
}

//NewUserRepo ...
func NewReviewRepo(db *sqlx.DB) *reviewRepo{
    return &reviewRepo{db: db}
}

func (r *reviewRepo) Create(user *pb.Review) (*pb.Review, error) {
    return nil, nil
}
