package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	pb "taxi_service/genproto/taxi"

	"github.com/jmoiron/sqlx"
)

type TaxiRepo struct {
	db *sqlx.DB
}

// NewTaxiRepo ...
func NewTaxiRepo(db *sqlx.DB) *TaxiRepo {
	return &TaxiRepo{db: db}
}

func (r *TaxiRepo) Create(taxi *pb.Taxi) (*pb.Taxi, error) {
	if taxi.Id == "" {
		Id, err := uuid.NewUUID()
		if err != nil {
			return nil, err
		}
		taxi.Id = Id.String()
	}
	query := `
		INSERT INTO taxis (
			id, first_name, last_name, email, password, birthday, car_id, phone_number,gender,profile_photo
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING
			id, first_name, last_name, email, password, birthday, car_id, phone_number, gender, profile_photo, created_at, updated_at`
	var respTaxi pb.Taxi
	err := r.db.QueryRow(
		query,
		taxi.Id, taxi.FirstName, taxi.LastName, taxi.Email, taxi.Password, taxi.Birthday, taxi.CarId, taxi.PhoneNumber, taxi.Gender, taxi.ProfilePhoto).
		Scan(
			&respTaxi.Id,
			&respTaxi.FirstName,
			&respTaxi.LastName,
			&respTaxi.Email,
			&respTaxi.Password,
			&respTaxi.Birthday,
			&respTaxi.CarId,
			&respTaxi.PhoneNumber,
			&respTaxi.Gender,
			&respTaxi.ProfilePhoto,
			&respTaxi.CreatedAt,
			&respTaxi.UpdatedAt)
	if err != nil {
		log.Println("Error creating taxi method in postgres")
		return nil, err
	}
	return &respTaxi, nil
}

func (r *TaxiRepo) Update(taxi *pb.Taxi) (*pb.Taxi, error) {
	query := `
	UPDATE
		taxis
	SET
		first_name = $1,
		last_name = $2,
		email = $3,
		password = $4,
		birthday = $5,
		car_id = $6,
		phone_number = $7,
		gender = $8,
		profile_photo = $9
	WHERE
		id = $10
	RETURNING
		id,
		first_name,
		last_name,
		email,
	    password,
	    birthday,
	    car_id,
	    phone_number,
	    gender,
	    profile_photo,
		created_at,
		updated_at
	`
	var respTaxi pb.Taxi
	err := r.db.QueryRow(
		query,
		taxi.FirstName,
		taxi.LastName,
		taxi.Email,
		taxi.Password,
		taxi.Birthday,
		taxi.CarId,
		taxi.PhoneNumber,
		taxi.Gender,
		taxi.ProfilePhoto,
		taxi.Id,
	).Scan(
		&respTaxi.Id,
		&respTaxi.FirstName,
		&respTaxi.LastName,
		&respTaxi.Email,
		&respTaxi.Password,
		&respTaxi.Birthday,
		&respTaxi.CarId,
		&respTaxi.PhoneNumber,
		&respTaxi.Gender,
		&respTaxi.ProfilePhoto,
		&respTaxi.CreatedAt,
		&respTaxi.UpdatedAt,
	)
	if err != nil {
		log.Println("Error updating taxi in postgres")
		return nil, err
	}
	return &respTaxi, nil
}

func (r *TaxiRepo) Delete(taxi *pb.TaxiRequest) (*pb.Taxi, error) {
	query := `
	UPDATE
		taxis
	SET
		deleted_at = CURRENT_TIMESTAMP
	WHERE
		id = $1
	AND
		deleted_at IS NULL
	RETURNING
		id,
		first_name,
		last_name,
		phone_number,
		gender,
		email,
		password,
		refresh_token,
		created_at,
		updated_at,
		deleted_at
	`
	var respTaxi pb.Taxi
	err := r.db.QueryRow(
		query,
		taxi.TaxiId,
	).Scan(
		&respTaxi.Id,
		&respTaxi.FirstName,
		&respTaxi.LastName,
		&respTaxi.PhoneNumber,
		&respTaxi.Gender,
		&respTaxi.Email,
		&respTaxi.Password,
		&respTaxi.RefreshToken,
		&respTaxi.CreatedAt,
		&respTaxi.UpdatedAt,
		&respTaxi.DeletedAt,
	)
	if err != nil {
		log.Println("Error deleting taxi in postgres")
		return nil, err
	}
	return &respTaxi, nil
}

func (r *TaxiRepo) Get(taxi *pb.TaxiRequest) (*pb.Taxi, error) {
	query := `
	SELECT
		id,
		first_name,
		last_name,
		phone_number,
		gender,
		email,
		password,
		refresh_token,
		created_at,
		updated_at
	FROM
		taxis
	WHERE
		id = $1
	AND
		deleted_at IS NULL
	`
	var respTaxi pb.Taxi
	err := r.db.QueryRow(
		query,
		taxi.TaxiId,
	).Scan(
		&respTaxi.Id,
		&respTaxi.FirstName,
		&respTaxi.LastName,
		&respTaxi.PhoneNumber,
		&respTaxi.Gender,
		&respTaxi.Email,
		&respTaxi.Password,
		&respTaxi.RefreshToken,
		&respTaxi.CreatedAt,
		&respTaxi.UpdatedAt,
	)
	if err != nil {
		log.Println("Error getting taxi in postgres")
		return nil, err
	}
	return &respTaxi, nil
}

func (r *TaxiRepo) GetAll(req *pb.GetAllTaxisRequest) (*pb.GetAllTaxisResponse, error) {
	offset := req.Limit * (req.Page - 1)
	query := `
	SELECT
		id,
		first_name,
		last_name,
		phone_number,
		gender,
		email,
		password,
		refresh_token,
		created_at,
		updated_at
	FROM
		taxis
	WHERE
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2
	`
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	var allTaxis pb.GetAllTaxisResponse
	for rows.Next() {
		var taxi pb.Taxi
		if err := rows.Scan(
			&taxi.Id,
			&taxi.FirstName,
			&taxi.LastName,
			&taxi.PhoneNumber,
			&taxi.Gender,
			&taxi.Email,
			&taxi.Password,
			&taxi.RefreshToken,
			&taxi.CreatedAt,
			&taxi.UpdatedAt,
		); err != nil {
			log.Println("Error adding taxi to list as get all logic")
			return nil, err
		}
		allTaxis.AllTaxis = append(allTaxis.AllTaxis, &taxi)
	}
	return &allTaxis, nil
}

func (r *TaxiRepo) CheckField(req *pb.CheckTaxi) (*pb.CheckRes, error) {
	var existsClient int
	query := fmt.Sprintf("SELECT count(1) FROM taxis WHERE '%s' = $1 AND deleted_at IS NULL", req.Field)
	if err := r.db.QueryRow(
		query,
		req.Value,
	).Scan(&existsClient); err != nil {
		return nil, err
	}
	if existsClient == 0 {
		return &pb.CheckRes{
			Exists: false,
		}, nil
	}
	return &pb.CheckRes{
		Exists: true,
	}, nil
}

func (r *TaxiRepo) GetTaxiByEmail(req *pb.EmailRequest) (*pb.Taxi, error) {
	query := `
	SELECT 
		id,
		first_name,
		last_name,
		phone_number,
		gender,
		email,
		password,
		refresh_token,
		created_at,
		updated_at
	FROM 
		taxis
	WHERE
		email = $1
	AND
		deleted_at IS NULL
	`
	var responseTaxi pb.Taxi
	if err := r.db.QueryRow(
		query,
		req.Email,
	).Scan(
		&responseTaxi.Id,
		&responseTaxi.FirstName,
		&responseTaxi.LastName,
		&responseTaxi.PhoneNumber,
		&responseTaxi.Gender,
		&responseTaxi.Email,
		&responseTaxi.Password,
		&responseTaxi.RefreshToken,
		&responseTaxi.CreatedAt,
		&responseTaxi.UpdatedAt,
	); err != nil {
		log.Println("Error getting taxi by email")
		return nil, err
	}
	return &responseTaxi, nil
}

func (r *TaxiRepo) GetTaxiByRefreshToken(req *pb.TaxiToken) (*pb.Taxi, error) {
	query := `
	SELECT 
		id,
		first_name,
		last_name,
		phone_number,
		gender,
		email,
		password,
		refresh_token,
		created_at,
		updated_at
	FROM 
		taxis
	WHERE
		refresh_token = $1
	AND
		deleted_at IS NULL
	`
	var responseTaxi pb.Taxi
	if err := r.db.QueryRow(
		query,
		req.RefreshToken,
	).Scan(
		&responseTaxi.Id,
		&responseTaxi.FirstName,
		&responseTaxi.LastName,
		&responseTaxi.PhoneNumber,
		&responseTaxi.Gender,
		&responseTaxi.Email,
		&responseTaxi.Password,
		&responseTaxi.RefreshToken,
		&responseTaxi.CreatedAt,
		&responseTaxi.UpdatedAt,
	); err != nil {
		log.Println("Error getting taxi by token")
		return nil, err
	}
	return &responseTaxi, nil
}
