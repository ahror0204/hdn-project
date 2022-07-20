package postgres

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	pb "github.com/hdn-project/user-service/genproto"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *pb.User) (*pb.Empty, error) {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error while generating uuid: ")
		return nil, err
	}
	tim := time.Now()
	query := `INSERT INTO Users (id,calendar_id, first_name, last_name,phone_numbers,email,status, payment_card,created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = r.db.Query(query, id, user.CalendarId, user.FirstName, user.LastName, pq.Array(user.PhoneNumbers), user.Email, user.Status, user.PaymentCard, tim)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (r *UserRepo) GetUserById(id string) (*pb.User, error) {
	var User pb.User
	query := `SELECT id, calendar_id, first_name, last_name, phone_numbers, email,status,
	payment_card,created_at FROM Users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&User.Id,
		&User.CalendarId,
		&User.FirstName,
		&User.LastName,
		pq.Array(&User.PhoneNumbers),
		&User.Email,
		&User.Status,
		&User.PaymentCard,
		&User.CreatedAt,
		// &User.UpdatedAt,
		// &User.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (r *UserRepo) DeleteById(id string) (*pb.Empty, error) {
	tim := time.Now()
	query := `UPDATE Users SET deleted_at = $1 WHERE id = $2`
	_, err := r.db.Exec(query, tim, id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (r *UserRepo) UpdateUser(user *pb.User) (*pb.Empty, error) {
	tim := time.Now()
	query := `UPDATE Users SET calendar_id = $1, first_name = $2, last_name = $3, 
	phone_numbers = $4, email = $5, status = $6, payment_card = $7,updated_at = $8 WHERE id = $9`
	_, err := r.db.Exec(query,
		user.CalendarId,
		user.FirstName,
		user.LastName,
		pq.Array(user.PhoneNumbers),
		user.Email,
		user.Status,
		user.PaymentCard,
		tim,
		user.Id,
	)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
