package warehouse

import db "github.com/lvlBA/restApi/internal/db/warehouse"

type ServiceImpl struct {
	db db.Service
}

func New(db db.Service) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}
