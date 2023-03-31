package warehouse

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lvlBA/restApi/internal/db"

	"github.com/doug-martin/goqu/v9"
	"github.com/lvlBA/restApi/internal/models"
)

const tableNameGoods = "goods"

type GoodsImpl struct {
	svc sqlService
}

type CreateGoodsParams struct {
	Name string
}

func (g *GoodsImpl) CreateGoods(ctx context.Context, params *CreateGoodsParams) (*models.Goods, error) {
	model := &models.Goods{
		Meta: models.Meta{},
		Name: params.Name,
	}
	model.UpdateMeta()

	id, err := g.svc.create(ctx, tableNameGoods, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

func (g *GoodsImpl) ReceiveGoods(ctx context.Context, id string) (*models.Goods, error) {
	result := &models.Goods{}

	query, _, err := goqu.From(tableNameGoods).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrorAlreadyExists
		}
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	if err = g.svc.GetContext(ctx, result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GoodsImpl) SendGoods(ctx context.Context, params *SendGoodsParams) error {
	return nil
}

func (g *GoodsImpl) DeleteGoods(ctx context.Context, id string) error {
	return g.svc.delete(ctx, tableNameGoods, id)
}
