package warehouse

import (
	"github.com/jmoiron/sqlx"
)

type txImpl struct {
	*sqlx.Tx
	*serviceImpl
}
