package warehouseapp

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/restApi/internal/controllers"
	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

func (s *ServiceImpl) DeleteGoods(ctx context.Context, req *api.DeleteGoodsRequest) (*api.DeleteGoodsResponse, error) {
	if err := validateDeleteLocationReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlGoods.DeleteGoods(ctx, req.ID); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to delete goods", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete goods")
	}

	return &api.DeleteGoodsResponse{}, nil
}

func validateDeleteLocationReq(req *api.DeleteGoodsRequest) error {
	return validation.Errors{
		"ID": validation.Validate(req.ID, validation.Required, is.UUIDv4),
	}.Filter()
}
