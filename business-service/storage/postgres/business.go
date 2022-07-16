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
			business_id, calendar_id, user_id, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

		err = r.db.QueryRow(staffQuery,
			stID,
			val.FirstName,
			val.LastName,
			pq.Array(val.PhoneNumbers),
			val.Cost,
			val.Status,
			bus.Id,
			val.CalendarId,
			val.UserId,
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
		//???????????? should we update clendar and user id
		staffQuery := `UPDATE staff SET first_name = $1, last_name = $2, phone_numbers = $3, cost = $4, status = $5, 
		calendar_id = $6, user_id = $7, updated_at = $8 WHERE business_id = $9`

		_, err = r.db.Query(staffQuery,
			val.FirstName,
			val.LastName,
			pq.Array(val.PhoneNumbers),
			val.Cost,
			val.Status,
			val.CalendarId,
			val.UserId,
			updTime,
			business.Id,
		)

		if err != nil {
			return nil, err
		}
	}

	return &pb.Empty{}, nil
}

func (r *businessRepo) DeleteBusiness(ID *pb.Id) (*pb.Empty, error) {

	delTime := time.Now()

	_, err := r.db.Query("UPDATE business SET deleted_at = $1 WHERE id = $2", delTime, ID.Id)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Query("UPDATE staff SET deleted_at = $1 WHERE business_id = $2", delTime, ID.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (r *businessRepo) GetByIdBusiness(ID *pb.Id) (*pb.Business, error) {
	// var (
	// 	bus pb.Business{}
	// )
	var bus = pb.Business{}
	getByIdQuery := `SELECT salon_name, phone_numbers, status, location, created_at FROM business WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(getByIdQuery, ID.Id).Scan(
		&bus.SalonName,
		pq.Array(&bus.PhoneNumbers),
		&bus.Status,
		&bus.Location,
		&bus.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	getStafftByIdQuery := `SELECT first_name, last_name, phone_numbers, cost, status, calendar_id, 
		user_id, created_at FROM staff WHERE business_id = $1 AND deleted_at IS NULL`

	rows, err := r.db.Query(getStafftByIdQuery, ID.Id)
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
			&staff.CalendarId,
			&staff.UserId,
			&staff.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		bus.Staff = append(bus.Staff, &staff)
	}
	return &bus, nil
}

func (r *businessRepo) GetAllBusiness(*pb.Empty) (*pb.GetAllBusinessResponse, error) {
	var allbusiness = pb.GetAllBusinessResponse{}

	getAllQuery := `SELECT id, salon_name, phone_numbers, status, location, created_at FROM business WHERE deleted_at IS NULL`

	rows, err := r.db.Query(getAllQuery)
	for rows.Next() {

		var bus pb.Business
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

		getStafftByIdQuery := `SELECT first_name, last_name, phone_numbers, cost, status, calendar_id, 
			user_id, created_at FROM staff WHERE business_id = $1 AND deleted_at IS NULL`
		rows1, err := r.db.Query(getStafftByIdQuery, bus.Id)
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
				&staff.CalendarId,
				&staff.UserId,
				&staff.CreatedAt,
			)
			if err != nil {
				return nil, err
			}
			bus.Staff = append(bus.Staff, &staff)
		}
		allbusiness.Business = append(allbusiness.Business, &bus)
	}
	return &allbusiness, nil
}

func (r *businessRepo) GetListBusiness(limit, page int64) (*pb.GetAllBusinessResponse, error) {
	offset := (page - 1) * limit

	var allbusiness = pb.GetAllBusinessResponse{}

	getListQuery := `SELECT id, salon_name, phone_numbers, status, location, created_at FROM business WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(getListQuery, limit, offset)
	
	for rows.Next() {
		var bus pb.Business

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

		getStafftByIdQuery := `SELECT first_name, last_name, phone_numbers, cost, status, calendar_id, 
			user_id, created_at FROM staff WHERE business_id = $1 AND deleted_at IS NULL`

		rows1, err := r.db.Query(getStafftByIdQuery, bus.Id)
		if err != nil {
			return nil, err
		}
		for rows1.Next() {
			var staff = pb.Staff{}
			err = rows1.Scan(
				&staff.FirstName,
				&staff.LastName,
				pq.Array(&staff.PhoneNumbers),
				&staff.Cost,
				&staff.Status,
				&staff.CalendarId,
				&staff.UserId,
				&staff.CreatedAt,
			)

			if err != nil {
				return nil, err
			}
			bus.Staff = append(bus.Staff, &staff)
		}
		allbusiness.Business = append(allbusiness.Business, &bus)
	}
	return &allbusiness, nil
}
//----------------------------------------------------------------------------------------------------

func (r *businessRepo) CreateService(user *pb.User) (*pb.MenServices, *pb.WomenServices, error) {

	var (
		menServise    pb.MenServices
		womenServise  pb.WomenServices
	)

	serviceID, err := uuid.NewV4()
	
	if err != nil {
		return nil, nil, err 
	}

	if user.Gender == "male" {
		menSerQuery := `INSERT INTO men_services (id, user_id) VALUES ($1, $2) RETURNING id, user_id`
		err = r.db.QueryRow(menSerQuery, serviceID,	user.Id).Scan(&menServise.Id, &menServise.UserId,)

		if err != nil {
			return nil, nil, err
		}

		return &menServise, nil, nil

	} else if user.Gender == "female" {

		womenServise.Id = serviceID.String()

		womenSerQuery := `INSERT INTO women_services (id, user_id) VALUES ($1,$2) RETURNING id, user_id`
		err = r.db.QueryRow(womenSerQuery, serviceID, user.Id).Scan(&womenServise.Id, &womenServise.UserId)
		if err != nil {
			return nil, nil,  err
		}

		return nil, &womenServise, nil
	}

	err = errors.New("Gender Type Error")

	return &pb.MenServices{}, &pb.WomenServices{}, err
}

func (r *businessRepo) UpdateMenServiceByUserId (menservice *pb.MenServices) (pb.Empty, error) {
	updateQuery := `UPDATE men_services SET hair_cut = $1, beard_cut = $2, hair_coloring = $3,
	special_hair_cut = $4, beard_coloring = $5, beard_trim = $6, beard_shave = $7, 
	face_shave = $8, boy_hair_cut = $9 WHERE user_id = $10`

	_, err := r.db.Query(updateQuery,
		menservice.HairCut,
		menservice.BeardCut,
		menservice.HairColoring,
		menservice.SpecialHairCut,
		menservice.BeardColoring,
		menservice.BeardTrim,
		menservice.BeardShave,
		menservice.FaceShave,
		menservice.BoyHairCut,
		menservice.UserId,
	)
	if err != nil {
		return pb.Empty{}, err 
	}
	return pb.Empty{}, nil
}

func (r *businessRepo) UpdateWomenServiceByUserId (womenservice *pb.WomenServices) (pb.Empty, error) {
	updateQuery := `UPDATE women_services SET hair_cut = $1, hair_coloring = $2,
	special_hair_cut = $3, eyebrow_arching = $4 WHERE user_id = $5`

	_, err := r.db.Query(updateQuery,
		womenservice.HairCut,
		womenservice.HairColoring,
		womenservice.SpecialHairCut,
		womenservice.EyebrowArching,
		womenservice.UserId,
	)
	if err != nil {
		return pb.Empty{}, err 
	}
	return pb.Empty{}, nil
}

func (r *businessRepo) GetMenServiceByUserId (ID *pb.Id) (*pb.MenServices, error) {
	var (
		menService pb.MenServices
	)
	query := `SELECT id, hair_cut, beard_cut, hair_coloring, special_hair_cut, beard_coloring, 
				beard_trim, beard_shave, face_shave, boy_hair_cut FROM men_services WHERE user_id = $1`
	err := r.db.QueryRow(query, ID.Id).Scan(
		&menService.Id,
		&menService.HairCut,
		&menService.BeardCut,
		&menService.HairColoring,
		&menService.SpecialHairCut,
		&menService.BeardColoring,
		&menService.BeardTrim,
		&menService.BeardShave,
		&menService.FaceShave,
		&menService.BoyHairCut,
	)
	if err != nil {
		return nil, err
	}
	return &menService, nil
}

func (r *businessRepo) GetWomenServiceByUserId (ID *pb.Id) (*pb.WomenServices, error) {
	var (
		womenService pb.WomenServices
	)
	query := `SELECT id, hair_cut, hair_coloring, eyebrow_arching, 
	special_hair_cut FROM women_services WHERE user_id = $1`
	err := r.db.QueryRow(query, ID.Id).Scan(
		&womenService.Id,
		&womenService.HairCut,
		&womenService.HairColoring,
		&womenService.EyebrowArching,
		&womenService.SpecialHairCut,
	)
	if err != nil {
		return nil, err
	}

	return &womenService, nil
}

func (r *businessRepo) DeleteMenServiceByUserId (ID *pb.Id) (*pb.Empty, error) {
	query := `DELETE FROM men_services WHERE user_id = $1`

	_, err := r.db.Query(query, ID.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (r *businessRepo) DeleteWomenServiceByUserId (ID *pb.Id) (*pb.Empty, error) {
	query := `DELETE FROM women_services WHERE user_id = $1`

	_, err := r.db.Query(query, ID.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (r *businessRepo) GetAllMenSetvices (*pb.Empty) ([]*pb.MenServices, error) {
	var (
		menService pb.MenServices
		menServices []*pb.MenServices
	)
	query := `SELECT id, hair_cut, beard_cut, hair_coloring, special_hair_cut, beard_coloring, 
				beard_trim, beard_shave, face_shave, boy_hair_cut FROM men_services`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() { 
		err = rows.Scan(
			&menService.Id,
			&menService.HairCut,
			&menService.BeardCut,
			&menService.HairColoring,
			&menService.SpecialHairCut,
			&menService.BeardColoring,
			&menService.BeardTrim,
			&menService.BeardShave,
			&menService.FaceShave,
			&menService.BoyHairCut,
		)
		if err != nil {
			return nil, err
		}
		menServices = append(menServices, &menService)
	}
	return menServices, nil
}

func (r *businessRepo) GetAllWomenSetvices (*pb.Empty) ([]*pb.WomenServices, error) {
	var (
		womenService pb.WomenServices
		womenServices []*pb.WomenServices
	)
	query := `SELECT id, hair_cut, hair_coloring, eyebrow_arching, 
	special_hair_cut FROM women_services`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() { 
		err = rows.Scan(
			&womenService.Id,
			&womenService.HairCut,
			&womenService.HairColoring,
			&womenService.EyebrowArching,
			&womenService.SpecialHairCut,
		)
		if err != nil {
			return nil, err
		}
		womenServices = append(womenServices, &womenService)
	}
	return womenServices, nil
}
