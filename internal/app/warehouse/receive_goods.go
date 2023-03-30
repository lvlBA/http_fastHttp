package warehouseapp

import (
	"context"
	"fmt"

	"github.com/lvlBA/restApi/internal/models"
	"github.com/lvlBA/restApi/internal/warehouse"
	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

func (s *ServiceImpl) Receive(ctx context.Context, req *api.ReceiveGoodsRequest) (*api.ReceiveGoodsResponse, error) {
	if err := validateReceiveGoodsReq(req); err != nil {
		return nil, err
	}
	// validateReq...

	resp, err := s.svc.ReceiveGoods(ctx, &warehouse.ReceiveGoodsParams{
		Name: req.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("error receive goods: %w", adaptError(err))
	}

	return &api.ReceiveGoodsResponse{
		Goods: adaptGoodsModel(resp),
	}, nil
}

func validateReceiveGoodsReq(req *api.ReceiveGoodsRequest) error {
	// ..
	return nil
}

func adaptGoodsModel(model *models.Goods) *api.Goods {
	return &api.Goods{
		ID:   model.ID,
		Name: model.Name,
	}
}
