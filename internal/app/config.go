package app

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/lvlBA/restApi/internal/bootstrap"
	"github.com/lvlBA/restApi/pkg/logger"
)

type Config struct {
	ListenAddress  string `json:"listen_address"     yaml:"listenAddress"      env:"LISTEN_ADDRESS"     envDefault:":8080"`
	ListenAddress2 string `json:"listen_address2"    yaml:"listenAddress2"     env:"LISTEN_ADDRESS2"    envDefault:":80"`
	DbHost         string `json:"db_host"            yaml:"db_host"            env:"DB_HOST"            envDefault:"postgres://db:db@localhost:5488/db"`
	LogLevel       string `json:"log_level"          yaml:"log_level"          env:"LOG_LEVEL"          envDefault:"error"`
}

func (c *Config) getDatabaseConnection() (*sqlx.DB, error) {
	return sqlx.Connect("pgx", c.DbHost)
}

func (c *Config) getLogger() (logger.Logger, error) {
	return bootstrap.InitLogger(&bootstrap.ConfigLogger{
		Level:             c.LogLevel,
		DisableCaller:     true,
		DisableStacktrace: true,
	})
}
