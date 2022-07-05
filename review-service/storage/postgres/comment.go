package postgres

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	pb "github.com/hdn-project/review-service/genproto"
	"github.com/hdn-project/review-service/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type commentRepo struct {
    db *sqlx.DB
}

//NewUserRepo ...
func NewCommentRepo(db *sqlx.DB) repo.CommentStorageI{
    return &commentRepo{db: db}
}

func (m *commentRepo) CreateComment(c *pb.Comment) (*pb.Comment, error) {
	now := time.Now()
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error while generating uuid")
		return nil, err
	}
	_, err = m.db.Exec(`
		INSERT INTO comments (id, user_id, staff_id, content, created_at,updated_at deleted_at) 
		VALUES ($1, $2, $3, $4, $5, $6)`, id, c.UserId, c.StaffId, c.Comment, now,nil, nil)
	if err != nil {
		fmt.Println("Error while inserting into comments")
		return nil, err
	}
	return c, nil
}

func(m *commentRepo) UpdateComment(r *pb.Comment) (*pb.Comment, error) {
	now := time.Now()
	_, err := m.db.Exec(`
		UPDATE comments SET comment = $1, updated_at = $2 WHERE id = $3`, r.Comment, now, r.Id)
	if err != nil {
		fmt.Println("Error while updating comments")
		return nil, err
	}
	return r, nil
}

func (m *commentRepo) DeleteComment(id, staffId string) error {
	_, err := m.db.Exec(`
		UPDATE comments SET deleted_at = $1 WHERE id = $2 AND staff_id = $3`, time.Now(), id, staffId)
	if err != nil {
		fmt.Println("Error while deleting comments")
		return err
	}
	return nil
}

func (m *commentRepo) DeleteUserComments(userId string) error {
	_, err := m.db.Exec(`
		Delete FROM comments WHERE user_id = $2`, userId)
	if err != nil {
		fmt.Println("Error while deleting comments")
		return err
	}
	return nil
}

func (m *commentRepo) CountComments(staffId string) (int64,  error) {
	var count int64
	err := m.db.Get(&count, `SELECT count(*) FROM comments WHERE staff_id = $1`, staffId)
	if err != nil {
		fmt.Println("Error while counting comments")
		return 0,  err
	}
	return count,  nil
}

func(m *commentRepo) ListComments(staffId string, limit, page int64) ([]*pb.Comment, int64, error) {
	var (
		comments []*pb.Comment
		count int64
	)

	offset := (page - 1) * limit
	
	rows, err := m.db.Queryx(`
		SELECT * FROM comments WHERE staff_id = $1 LIMIT $2 OFFSET $3`, staffId, limit, offset)
	
	if err != nil {
		fmt.Println("Error while listing comments")
		return nil, 0, err
	}

	for rows.Next() {
		var com pb.Comment
		err := rows.Scan(&com.Id, &com.UserId, &com.StaffId, pq.Array(&com.Comment), &com.CreatedAt, &com.UpdatedAt, &com.DeletedAt)
		if err != nil {
			fmt.Println("Error while scanning comments")
			return nil, 0, err
		}
		comments = append(comments, &com)
	}


	err = m.db.Get(&count, `SELECT count(*) FROM comments WHERE staff_id = $1`, staffId)
	
	if err != nil {
		fmt.Println("Error while counting comments")
		return nil, 0, err
	}
	return comments, count, nil
}
