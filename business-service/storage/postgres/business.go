package postgres

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	// business "github.com/hdn-project/business-service/genproto"
	pb "github.com/hdn-project/business-service/genproto"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type businessRepo struct {
	db *sqlx.DB
}

//NewBusinessRepo ...
func NewBusinessRepo(db *sqlx.DB) *businessRepo {
	return &businessRepo{db: db}
}

func (r *businessRepo) CreateBusiness(business *pb.Business) (*pb.Business, error) {

	var (
		bus pb.Business
	)

	busId, err := uuid.NewV4()
	crtTime := time.Now()
	if err != nil {
		return nil, err
	}

	busQuery := `INSERT INTO business (id, salon_name, phone_numbers, status, location, created_at)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, salon_name, phone_numbers, status, location, created_at`

	err = r.db.QueryRow(busQuery,
		busId,
		business.SalonName,
		pq.Array(business.PhoneNumbers),
		business.Status,
		business.Location,
		crtTime,
	).Scan(
		&bus.Id,
		&bus.SalonName,
		pq.Array(&bus.PhoneNumbers),
		&bus.Status,
		&bus.Location,
		&bus.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	for _, val := range business.Staff {
		stID, err := uuid.NewV4()

		if err != nil {
			return nil, err
		}
		staffQuery := `INSERT INTO staff (id, first_name, last_name, phone_numbers, cost, status, 
		comment_id, business_id, calendar_id, client_id, men_service_id, women_service_id, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id`

		err = r.db.QueryRow(staffQuery,
			stID,
			val.FirstName,
			val.LastName,
			pq.Array(val.PhoneNumbers),
			val.Cost,
			val.Status,
			val.CommentId,
			bus.Id,
			val.CalendarId,
			val.ClientId,
			val.MenServicesId,
			val.WomenServicesId,
			crtTime,
		).Scan(&stID)
		if err != nil {
			return nil, err
		}
	}

	return &bus, nil
}

func (r *businessRepo) UpdateBusiness(business *pb.Business) (*pb.Empty, error) {

	updTime := time.Now()

	busQuery := `UPDATE business SET salon_name = $1, phone_numbers = $2, status = $3, location = $4, updated_at = $5 WHERE id = $6`

	_, err := r.db.Query(busQuery, business.SalonName, pq.Array(business.PhoneNumbers), business.Status, business.Location, updTime, business.Id)
	if err != nil {
		return nil, err
	}

	for _, val := range business.Staff {

		staffQuery := `UPDATE staff SET first_name = $1, last_name = $2, phone_number = $3, cost = 4, status = $5, comment_id = $6, 
		calendar_id = $7, client_id = $8, men_service_id = $9, women_service_id = $10, updated_at = $11 WHERE business_id = $12`

		_, err = r.db.Query(staffQuery,
			val.FirstName,
			val.LastName,
			pq.Array(val.PhoneNumbers),
			val.Cost,
			val.Status,
			val.CommentId,
			val.CalendarId,
			val.ClientId,
			val.MenServicesId,
			val.WomenServicesId,
			updTime,
			business.Id,
		)

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (r *businessRepo) DeleteBusiness(ID *pb.Id) (*pb.Empty, error) {

	delTime := time.Now()

	_, err := r.db.Query("UPDATE business SET deleted_at = $1 WHERE id = $2", delTime, ID)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Query("UPDATE business SET deleted_at = $1 WHERE business_id = $2", delTime, ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *businessRepo) GetByIdBusiness(ID *pb.Id) (*pb.Business, error) {
	var (
		bus *pb.Business
	)
	getByIdQuery := `SELECT salon_name, phone_numbers, status, location, created_at FROM business WHERE id =1$`
	err := r.db.QueryRow(getByIdQuery, ID).Scan(
		&bus.SalonName,
		pq.Array(&bus.PhoneNumbers),
		&bus.Status,
		&bus.Location,
		&bus.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	getStafftByIdQuery := `SELECT first_name, last_name, phone_number, cost, status, comment_id, calendar_id, 
		client_id, men_service_id, women_service_id, created_at FROM staff WHERE business_id = $1`

	rows, err := r.db.Query(getStafftByIdQuery, ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var staff pb.Staff
		err = rows.Scan(
			&staff.FirstName,
			&staff.LastName,
			pq.Array(&staff.PhoneNumbers),
			&staff.Cost,
			&staff.Status,
			&staff.CommentId,
			&staff.BusinessId,
			&staff.CalendarId,
			&staff.ClientId,
			&staff.MenServicesId,
			&staff.WomenServicesId,
			&staff.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		bus.Staff = append(bus.Staff, &staff)
	}
	return bus, nil
}

func (r *businessRepo) GetAllBusiness(*pb.Empty) (*pb.GetAllBusinessResponse, error) {
	var allbusiness []*pb.Business

	getByIdQuery := `SELECT id, salon_name, phone_numbers, status, location, created_at FROM business WHERE deleted_at = $2`
	rows, err := r.db.Query(getByIdQuery, "")
	for rows.Next() {
		var bus *pb.Business
		err = rows.Scan(
			&bus.Id,
			&bus.SalonName,
			pq.Array(&bus.PhoneNumbers),
			&bus.Status,
			&bus.Location,
			&bus.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		getStafftByIdQuery := `SELECT first_name, last_name, phone_number, cost, status, comment_id, calendar_id, 
			client_id, men_service_id, women_service_id, created_at FROM staff WHERE business_id = $1 AND deleted_at = $2`

		rows1, err := r.db.Query(getStafftByIdQuery, bus.Id, "")
		if err != nil {
			return nil, err
		}
		for rows1.Next() {
			var staff pb.Staff
			err = rows1.Scan(
				&staff.FirstName,
				&staff.LastName,
				pq.Array(&staff.PhoneNumbers),
				&staff.Cost,
				&staff.Status,
				&staff.CommentId,
				&staff.BusinessId,
				&staff.CalendarId,
				&staff.ClientId,
				&staff.MenServicesId,
				&staff.WomenServicesId,
				&staff.CreatedAt,
			)
			if err != nil {
				return nil, err
			}
			bus.Staff = append(bus.Staff, &staff)
		}
		allbusiness = append(allbusiness, bus)
	}
	return &pb.GetAllBusinessResponse{
		Business: allbusiness,
	}, nil
}

func (r *businessRepo) GetListBusiness(limit, page int64) (*pb.GetAllBusinessResponse, error) {
	offset := (page - 1) * limit

	var allbusiness []*pb.Business

	getByIdQuery := `SELECT id, salon_name, phone_numbers, status, location, created_at FROM business WHERE deleted_at = $1 LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(getByIdQuery, "", limit, offset)
	for rows.Next() {
		var bus *pb.Business
		err = rows.Scan(
			&bus.Id,
			&bus.SalonName,
			pq.Array(&bus.PhoneNumbers),
			&bus.Status,
			&bus.Location,
			&bus.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		getStafftByIdQuery := `SELECT first_name, last_name, phone_number, cost, status, comment_id, calendar_id, 
			client_id, men_service_id, women_service_id, created_at FROM staff WHERE business_id = $1 AND deleted_at = $2`

		rows1, err := r.db.Query(getStafftByIdQuery, bus.Id, "")
		if err != nil {
			return nil, err
		}
		for rows1.Next() {
			var staff pb.Staff
			err = rows1.Scan(
				&staff.FirstName,
				&staff.LastName,
				pq.Array(&staff.PhoneNumbers),
				&staff.Cost,
				&staff.Status,
				&staff.CommentId,
				&staff.BusinessId,
				&staff.CalendarId,
				&staff.ClientId,
				&staff.MenServicesId,
				&staff.WomenServicesId,
				&staff.CreatedAt,
			)
			if err != nil {
				return nil, err
			}
			bus.Staff = append(bus.Staff, &staff)
		}
		allbusiness = append(allbusiness, bus)
	}
	return &pb.GetAllBusinessResponse{
		Business: allbusiness,
	}, nil
}

func (r *businessRepo) CreateService(client *pb.User) (*pb.ServiceTypeDef, error) {

	var (
		services pb.ServiceTypeDef
		menServise   pb.MenServices
		womenServise pb.WomenServices
	)

	serviceID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	if client.Gender == "man" {

		menSerQuery := `INSERT INTO men_services (id, hair_cut, beard_cut, hair_coloring,
			special_hair_cut, beard_coloring, beard_trim, beard_shave, face_shave, boy_hair_cut, client_id) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10, $11) RETURNING id, hair_cut, beard_cut, hair_coloring,
			special_hair_cut, beard_coloring, beard_trim, beard_shave, face_shave, boy_hair_cut, client_id`
		err = r.db.QueryRow(menSerQuery,
			serviceID,
			services.MenService.HairCut,
			services.MenService.BeardCut,
			services.MenService.HairColoring,
			services.MenService.SpecialHairCut,
			services.MenService.BeardColoring,
			services.MenService.BeardTrim,
			services.MenService.BeardShave,
			services.MenService.FaceShave,
			services.MenService.BoyHairCut,
			client.Id,
		).Scan(
			&menServise.Id,
			&menServise.HairCut,
			&menServise.BeardCut,
			&menServise.HairColoring,
			&menServise.SpecialHairCut,
			&menServise.BeardColoring,
			&menServise.BeardTrim,
			&menServise.BeardShave,
			&menServise.FaceShave,
			&menServise.BoyHairCut,
			&menServise.ClientId,
		)
		if err != nil {
			return nil, err
		}

		return &pb.ServiceTypeDef{
			MenService:   &menServise,
			WomenService: nil,
		}, nil

	} else if client.Gender == "woman" {

		womenServise.Id = serviceID.String()

		womenSerQuery := `INSERT INTO women_services (id, hair_cut, hair_coloring,
			eyebrow_coloring, special_hair_cut, client_id) 
			VALUES ($1,$2,$3,$4,$5, $6) RETURNING id, hair_cut, hair_coloring,
			eyebrow_coloring, special_hair_cut, client_id`
		err = r.db.QueryRow(womenSerQuery,
			serviceID,
			services.WomenService.HairCut,
			services.WomenService.HairColoring,
			services.WomenService.EyebrowArching,
			services.WomenService.SpecialHairCut,
			client.Id,
		).Scan(
			&womenServise.Id,
			&womenServise.HairCut,
			&womenServise.HairColoring,
			&womenServise.EyebrowArching,
			&womenServise.SpecialHairCut,
			&womenServise.ClientId,
		)
		if err != nil {
			return nil, err
		}

		return &pb.ServiceTypeDef{
			MenService:   nil,
			WomenService: &womenServise,
		}, nil
	}

	err = errors.New("Gender Type Error")

	return nil, err
}
