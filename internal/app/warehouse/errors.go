package warehouseapp

import (
	"errors"

	"github.com/lvlBA/restApi/internal/warehouse"
)

var (
	ErrorNotFound      = errors.New("not found")
	ErrorAlreadyExists = errors.New("already exists")
	ErrorInternal      = errors.New("internal")
)

func adaptError(err error) error {
	switch {
	case err == nil:
		return err
	case errors.Is(err, warehouse.ErrorNotFound):
		return ErrorNotFound
	case errors.Is(err, warehouse.ErrorAlreadyExists):
		return ErrorAlreadyExists
	default:
		return ErrorInternal
	}
}
