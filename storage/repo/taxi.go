package repo

import (
	pb "taxi_service/genproto/taxi"
)

// TaxiStorageI ...
type TaxiStorageI interface {
	Create(user *pb.Taxi) (*pb.Taxi, error)
	Update(request *pb.Taxi) (*pb.Taxi, error)
	Delete(request *pb.TaxiRequest) (*pb.Taxi, error)
	Get(request *pb.TaxiRequest) (*pb.Taxi, error)
	GetAll(request *pb.GetAllTaxisRequest) (*pb.GetAllTaxisResponse, error)
	CheckField(req *pb.CheckTaxi) (*pb.CheckRes, error)
	GetTaxiByEmail(req *pb.EmailRequest) (*pb.Taxi, error)
	GetTaxiByRefreshToken(req *pb.TaxiToken) (*pb.Taxi, error)
}
