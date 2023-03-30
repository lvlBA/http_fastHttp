package warehouseapp

import (
	"context"

	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

type Service interface {
	Receive(ctx context.Context, req *api.ReceiveGoodsRequest) (*api.ReceiveGoodsResponse, error)
}
