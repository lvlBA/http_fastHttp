package warehouse

import (
	"context"

	"github.com/lvlBA/restApi/internal/controllers"
)

func (s *ServiceImpl) DeleteGoods(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.Goods().DeleteGoods(ctx, id))
}
