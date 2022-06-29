package postgres

import (
	pb "github.com/hdn-project/client-service/genproto"
	"github.com/jmoiron/sqlx"
)

type clientRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewClientRepo(db *sqlx.DB) *clientRepo {
	return &clientRepo{db: db}
}

func (r *clientRepo) CreateUser(user *pb.Client) (*pb.Empty, error) {
	query := `INSERT INTO clients (id,calendar_id, first_name, last_name,phone_numbers,email,status, payment_card,created_at,updated_at,deleted_at,) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := r.db.Exec(query, user.Id, user.CalendarId, user.FirstName, user.LastName, user.PhoneNumbers, user.Email, user.Status, user.PaymentCard, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (r *clientRepo) GetClientById(id string) (*pb.Client,error) {
	var client pb.Client
	query := `SELECT * FROM clients WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&client.Id,
		&client.CalendarId,
		&client.FirstName,
		&client.LastName,
		&client.PhoneNumbers,
		&client.Email,
		&client.Status,
		&client.PaymentCard,
		&client.CreatedAt,
		&client.UpdatedAt,
		&client.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *clientRepo) DeleteById(id string) (*pb.Empty, error) {
	query := `DELETE FROM clients WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (r *clientRepo) UpdateClient(user *pb.Client) (*pb.Empty, error) {
	query := `UPDATE clients SET calendar_id = $1, first_name = $2, last_name = $3, 
	phone_numbers = $4, email = $5, status = $6, payment_card = $7, created_at = $8, 
	updated_at = $9, deleted_at = $10 WHERE id = $11`
	_, err := r.db.Exec(query, user.CalendarId, user.FirstName, user.LastName, 
		user.PhoneNumbers, user.Email, user.Status, user.PaymentCard, user.CreatedAt, 
		user.UpdatedAt, user.DeletedAt, user.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}