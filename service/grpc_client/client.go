package grpcClient

import (
	"taxi_service/config"
)

type IServiceManager interface {
}

type serviceManager struct {
	cfg config.Config
}

func New(cfg config.Config) (IServiceManager, error) {
	return &serviceManager{
		cfg: cfg,
	}, nil
}
