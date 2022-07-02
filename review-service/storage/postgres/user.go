package postgres

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"

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

func (r *reviewRepo) CreateComment(user *pb.Comment) (*pb.Empty, error) {
    id, err := uuid.NewV4()
    if err != nil {
        fmt.Println("Error while generating UUID")
        return nil, err
    }
    tim := time.Now()
    query := `INSERT INTO reviews (id,client_id,business_id, likes,dislikes,comment,creted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
    _, err = r.db.Exec(query, id, user.ClientId, user.BusinessId, user.Like, user.Dislike, pq.Array(user.Comment), tim)
    if err != nil {
        fmt.Println("Error while inserting data")
        return nil, err
    }
    return &pb.Empty{}, nil
}

func (r *reviewRepo) CreateReply(req *pb.ReplyComments) (*pb.Empty, error) {
    id, err := uuid.NewV4()
    if err != nil {
        fmt.Println("Error while generating UUID")
        return nil, err
    }
    tim := time.Now()
    query := `INSERT INTO reviews (id, review_id, reply_comment, created_at) VALUES ($1, $2, $3)`
    _, err = r.db.Exec(query, req.ReviewId, id, pq.Array(req.ReplyComment), tim) 
    if err != nil {
        fmt.Println("Error while inserting data")
        return nil, err
    }
    return &pb.Empty{}, nil  
}
