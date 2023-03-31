package warehouse

import (
	"context"
	"github.com/lvlBA/restApi/internal/controllers"
	"github.com/lvlBA/restApi/internal/models"
)

func (s *ServiceImpl) ReceiveGoods(ctx context.Context, id string) (*models.Goods, error) {
	resp, err := s.db.Goods().ReceiveGoods(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
