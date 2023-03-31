package warehouseapp

import (
	controllersGoods "github.com/lvlBA/restApi/internal/controllers/warehouse"
	"github.com/lvlBA/restApi/pkg/logger"
)

type ServiceImpl struct {
	ctrlGoods controllersGoods.Service
	log       logger.Logger
}

func New(ctrlGoods controllersGoods.Service, l logger.Logger) *ServiceImpl {
	return &ServiceImpl{
		ctrlGoods: ctrlGoods,
		log:       l,
	}
}
