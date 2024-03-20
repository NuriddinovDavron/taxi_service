package service

import (
	"context"
	pbu "taxi_service/genproto/taxi"
	l "taxi_service/pkg/logger"
	"taxi_service/storage"

	grpcClient "taxi_service/service/grpc_client"

	"github.com/jmoiron/sqlx"
)

// TaxiService ...
type TaxiService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewTaxiService ...
func NewTaxiService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *TaxiService {
	return &TaxiService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *TaxiService) Create(ctx context.Context, req *pbu.Taxi) (*pbu.Taxi, error) {
	taxi, err := s.storage.Taxi().Create(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return taxi, nil
}

func (s *TaxiService) Update(ctx context.Context, req *pbu.Taxi) (*pbu.Taxi, error) {
	taxi, err := s.storage.Taxi().Update(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return taxi, nil
}

func (s *TaxiService) Delete(ctx context.Context, req *pbu.TaxiRequest) (*pbu.Taxi, error) {
	taxi, err := s.storage.Taxi().Delete(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return taxi, nil
}

func (s *TaxiService) Get(ctx context.Context, req *pbu.TaxiRequest) (*pbu.Taxi, error) {
	taxi, err := s.storage.Taxi().Get(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return taxi, nil
}

func (s *TaxiService) GetAll(ctx context.Context, req *pbu.GetAllTaxisRequest) (*pbu.GetAllTaxisResponse, error) {
	taxis, err := s.storage.Taxi().GetAll(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return taxis, nil
}

func (s *TaxiService) CheckField(ctx context.Context, req *pbu.CheckTaxi) (*pbu.CheckRes, error) {
	check, err := s.storage.Taxi().CheckField(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return check, nil
}

func (s *TaxiService) GetTaxiByEmail(ctx context.Context, req *pbu.EmailRequest) (*pbu.Taxi, error) {
	res, err := s.storage.Taxi().GetTaxiByEmail(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *TaxiService) GetTaxiByRefreshToken(ctx context.Context, req *pbu.TaxiToken) (*pbu.Taxi, error) {
	res, err := s.storage.Taxi().GetTaxiByRefreshToken(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}
