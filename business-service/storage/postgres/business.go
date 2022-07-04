package postgres

import (
	"errors"

	"github.com/gofrs/uuid"
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

	if err != nil {
		return nil, err
	}

	busQuery := `INSERT INTO business (id, salon_name, phone_numbers, status, location)
	VALUES ($1, $2, $3, $4, $5) RETURNIN id, salon_name, phone_numbers, status, location`

	err = r.db.QueryRow(busQuery,
		busId,
		business.SalonName,
		pq.Array(business.PhoneNumbers),
		business.Status,
		business.Location,
	).Scan(
		&bus.Id,
		&bus.SalonName,
		pq.Array(&bus.PhoneNumbers),
		&bus.Status,
		&bus.Location,
	)
	if err != nil {
		return nil, err
	}

	for _, val := range business.Staff {
		stID, err := uuid.NewV4()

		if err != nil {
			return nil, err
		}
		staffQuery := `INSERT INTO staff (id, first_name, last_name, phone_number, cost, status, 
		comment_id, business_id, calendar_id, client_id, men_service_id, women_service_id)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id`

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
		).Scan(&stID)
		if err != nil {
			return nil, err
		}
	}

	return &bus, nil
}

func (r *businessRepo) DefServiceType(services *pb.ServiceTypeDef) (*pb.ServiceTypeDef, error) {
	
	var (
		menServise pb.MenServices
		womenServise pb.WomenServices
	)

	serviceID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	
	if services.MenService != nil {

		menSerQuery := `INSERT INTO men_services (id, hair_cut, beard_cut, hair_coloring,
			special_hair_cut, beard_coloring, beard_trim, beard_shave, face_shave, boy_hair_cut) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id, hair_cut, beard_cut, hair_coloring,
			special_hair_cut, beard_coloring, beard_trim, beard_shave, face_shave, boy_hair_cut`
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
			services.MenService.ClientId,
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
		
		return &pb.ServiceTypeDef {
			MenService: &menServise, 
			WomenService: nil,
		}, nil

	}else if services.WomenService != nil {
		
		womenServise.Id = serviceID.String()
		

		womenSerQuery := `INSERT INTO women_services (id, hair_cut, hair_coloring,
			eyebrow_coloring, special_hair_cut) 
			VALUES ($1,$2,$3,$4,$5) RETURNING id, hair_cut, hair_coloring,
			eyebrow_coloring, special_hair_cut`
		err = r.db.QueryRow(womenSerQuery,
			serviceID,
			services.WomenService.HairCut,
			services.WomenService.HairColoring,
			services.WomenService.EyebrowArching,
			services.WomenService.SpecialHairCut,
			services.WomenService.ClientId,
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

		return &pb.ServiceTypeDef {
			MenService: nil, 
			WomenService: &womenServise,
		}, nil
	}

	err = errors.New("Gender Type Error")

	return nil, err
}


