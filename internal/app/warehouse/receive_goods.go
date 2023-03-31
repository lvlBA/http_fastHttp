package warehouseapp

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errors "github.com/lvlBA/restApi/internal/controllers"
	"github.com/lvlBA/restApi/internal/models"
	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

func (s *ServiceImpl) ReceiveGoods(ctx context.Context, req *api.ReceiveGoodsRequest) (*api.ReceiveGoodsResponse, error) {
	if err := validateReceiveGoodsReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := s.ctrlGoods.ReceiveGoods(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("error receive goods: %w", errors.AdaptingErrorDB(err))
	}

	return &api.ReceiveGoodsResponse{
		Goods: adaptGoodsModel(resp),
	}, nil
}

func validateReceiveGoodsReq(req *api.ReceiveGoodsRequest) error {
	return validation.Errors{
		"ID": validation.Validate(req.ID, validation.Required),
	}.Filter()
}

func adaptGoodsModel(model *models.Goods) *api.Goods {
	return &api.Goods{
		ID:   model.ID,
		Name: model.Name,
	}
}
