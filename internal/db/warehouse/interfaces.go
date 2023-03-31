package warehouse

import (
	"context"
	"database/sql"

	"github.com/lvlBA/restApi/internal/models"
)

type ReceiveGoodsParams struct {
	ID   string
	Name string
}

type SendGoodsParams struct {
	ID string
}

type sqlClient interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type services interface {
	Goods() Goods
}

type Service interface {
	services
	Begin(ctx context.Context) (Transaction, error)
}

type Transaction interface {
	services
	Commit() error
	Rollback() error
}

type sqlService interface {
	services
	sqlClient
	create(ctx context.Context, table string, req any) (string, error)
	update(ctx context.Context, table, id string, req any) error
	delete(ctx context.Context, table, id string) error
}

type Goods interface {
	CreateGoods(ctx context.Context, params *CreateGoodsParams) (*models.Goods, error)
	ReceiveGoods(ctx context.Context, id string) (*models.Goods, error)
	SendGoods(ctx context.Context, params *SendGoodsParams) error
	DeleteGoods(ctx context.Context, id string) error
}
