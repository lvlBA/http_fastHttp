package warehouseapp

import (
	"context"
	"errors"
	"github.com/lvlBA/restApi/internal/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/restApi/internal/controllers"
	controllersGoods "github.com/lvlBA/restApi/internal/controllers/warehouse"
	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

func (s *ServiceImpl) CreateGoods(ctx context.Context, req *api.CreateGoodsRequest) (*api.CreateGoodsResponse, error) {
	if err := validateCreateGoodsReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	goods, err := s.ctrlGoods.CreateGoods(ctx, &controllersGoods.CreateGoodsParams{
		Name: req.Name,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "goods already exists")
		}
		s.log.Error(ctx, "failed to create goods", err, "request", req)

		return nil, status.Error(codes.Internal, "error create goods")
	}

	return &api.CreateGoodsResponse{
		Goods: adaptGoodsToApi(goods),
	}, nil
}

func validateCreateGoodsReq(req *api.CreateGoodsRequest) error {
	return validation.Errors{
		"name": validation.Validate(req.Name, validation.Required),
	}.Filter()
}

func adaptGoodsToApi(model *models.Goods) *api.Goods {
	return &api.Goods{
		ID:   model.ID,
		Name: model.Name,
	}
}
