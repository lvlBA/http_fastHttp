package warehouse

import (
	"context"
	"github.com/lvlBA/restApi/internal/models"
)

type Service interface {
	CreateGoods(ctx context.Context, params *CreateGoodsParams) (*models.Goods, error)
	ReceiveGoods(ctx context.Context, id string) (*models.Goods, error)
	DeleteGoods(ctx context.Context, id string) error
}
