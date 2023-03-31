package warehouse

import (
	"context"

	"github.com/lvlBA/restApi/internal/controllers"
	db "github.com/lvlBA/restApi/internal/db/warehouse"
	"github.com/lvlBA/restApi/internal/models"
)

type CreateGoodsParams struct {
	Name string
}

func (s *ServiceImpl) CreateGoods(ctx context.Context, params *CreateGoodsParams) (*models.Goods, error) {
	resp, err := s.db.Goods().CreateGoods(ctx, &db.CreateGoodsParams{
		Name: params.Name,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
