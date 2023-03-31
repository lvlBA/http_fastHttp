package warehouseapp

import (
	"context"
	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

type Service interface {
	CreateGoods(ctx context.Context, req *api.CreateGoodsRequest) (*api.CreateGoodsResponse, error)
	DeleteGoods(ctx context.Context, req *api.DeleteGoodsRequest) (*api.DeleteGoodsResponse, error)
	ReceiveGoods(ctx context.Context, req *api.ReceiveGoodsRequest) (*api.ReceiveGoodsResponse, error)
}
