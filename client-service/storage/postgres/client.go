package postgres

import (
	"time"

	"github.com/gofrs/uuid"
	pb "github.com/hdn-project/client-service/genproto"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type clientRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewClientRepo(db *sqlx.DB) *clientRepo {
	return &clientRepo{db: db}
}

func (r *clientRepo) CreateUser(user *pb.Client) (*pb.Empty, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	user.Id = id.String()
	tim := time.Now()
	query := `INSERT INTO clients (id,calendar_id, first_name, last_name,phone_numbers,email,status, payment_card,created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = r.db.Exec(query, user.Id, user.CalendarId, user.FirstName, user.LastName, pq.Array(user.PhoneNumbers), user.Email, user.Status, user.PaymentCard,tim)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (r *clientRepo) GetClientById(id string) (*pb.Client,error) {
	var client pb.Client
	query := `SELECT id, calendar_id, first_name, last_name, phone_numbers, email,status,
	payment_card,created_at FROM clients WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&client.Id,
		&client.CalendarId,
		&client.FirstName,
		&client.LastName,
		pq.Array(&client.PhoneNumbers),
		&client.Email,
		&client.Status,
		&client.PaymentCard,
		&client.CreatedAt,
		// &client.UpdatedAt,
		// &client.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *clientRepo) DeleteById(id string) (*pb.Empty, error) {
	tim := time.Now()
	query := `UPDATE clients SET deleted_at = $1 WHERE id = $2`
	_, err := r.db.Exec(query, tim, id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (r *clientRepo) UpdateClient(user *pb.Client) (*pb.Empty, error) {
	tim := time.Now()
	query := `UPDATE clients SET calendar_id = $1, first_name = $2, last_name = $3, 
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