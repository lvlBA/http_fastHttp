package warehouse

import (
	"context"

	"github.com/lvlBA/restApi/internal/models"
)

type ReceiveGoodsParams struct {
	Name string
}

type SendGoodsParams struct {
	ID string
}

type Service interface {
	ReceiveGoods(ctx context.Context, params *ReceiveGoodsParams) (*models.Goods, error)
	SendGoods(ctx context.Context, params *SendGoodsParams) error
}

type ServiceImpl struct{}

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) ReceiveGoods(ctx context.Context, params *ReceiveGoodsParams) (*models.Goods, error) {
	return &models.Goods{
		ID:   "Some Id",
		Name: params.Name,
	}, nil
}

func (s *ServiceImpl) SendGoods(ctx context.Context, params *SendGoodsParams) error {
	return nil
}
