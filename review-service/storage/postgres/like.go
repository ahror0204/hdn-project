package postgres

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"

	pb "github.com/hdn-project/review-service/genproto"
	"github.com/hdn-project/review-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type likeRepo struct {
	db *sqlx.DB
}

//NewLikeRepo ...
func NewLikeRepo(db *sqlx.DB) repo.LikeStorageI {
	return &likeRepo{db: db}
}

func (r *likeRepo) Create(l *pb.Like) error {
	now := time.Now()
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error while generating uuid")
		return err
	}
	_, err = r.db.Exec(`
        INSERT INTO likes (id, user_id,staff_id,created_at, deleted_at) 
        VALUES ($1, $2, $3, $4, $5)`, id, l.UserId,l.StaffId, now, nil)
	if err != nil {
		fmt.Println("Error while inserting into likes")
		return err
	}
	return nil
}

func (r *likeRepo) DeleteUserLikes(userId string) error {
	_, err := r.db.Exec(`
        UPDATE likes SET deleted_at = $1 WHERE user_id = $2`, time.Now(), userId)
	if err != nil {
		fmt.Println("Error while deleting likes")
		return err
	}
	return nil
}

func (r *likeRepo) DeleteLike(l *pb.Like) error {
	_, err := r.db.Exec(`DELETE FROM likes WHERE user_id = $1`, l.UserId)
	if err != nil {
		fmt.Println("Error while deleting likes")
		return err
	}
	return nil
}

func (r *likeRepo) CountLikes(l *pb.Like) (int64, bool, error) {

	var count int64
	err := r.db.Get(&count, `SELECT count(*) FROM likes WHERE user_id = $1`, l.UserId)
	if err != nil {
		fmt.Println("Error while counting likes")
		return 0, false, err
	}
	return count, true, nil
}


func (r *likeRepo) ListLikeUsers(staffId,userId string, limit, page int64) ([]*pb.LikedUser, int64, error) {
    var users []*pb.LikedUser
    var count int64

    offset :=(page - 1) * limit

    rows,err := r.db.Query(`
    SELECT user_id
    FROM likes
    WHERE staff_id = $1 LIMIT $2 OFFSET $3`, staffId, limit,offset)
    if err != nil {
        fmt.Println("Error while listing likes")
        return nil, 0, err
    }

	defer rows.Close()

	for rows.Next() {
		var user pb.LikedUser
		if err := rows.Scan(&user.Id); err != nil {
			fmt.Println("Error while scanning likes")
			return nil, 0, err
		}
		
		users = append(users, &user)
	}

	count, err = r.ListLikeUsersCount(staffId)
	if err != nil {
		fmt.Println("Error while listing likes")
		return nil, 0, err
	}
	
    return users, count, nil
}

func (r *likeRepo) ListLikeUsersCount(staffId string) (int64, error) {
	var count int64

	row := r.db.QueryRow(`
	SELECT count(*)
	FROM likes
	WHERE staff_id = $1`, staffId)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}