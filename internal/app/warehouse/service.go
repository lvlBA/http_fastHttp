package warehouseapp

import "github.com/lvlBA/restApi/internal/warehouse"

type ServiceImpl struct {
	svc warehouse.Service
}

func New(svc warehouse.Service) *ServiceImpl {
	return &ServiceImpl{
		svc: svc,
	}
}
