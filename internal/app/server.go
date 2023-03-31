package app

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"net"

	app "github.com/lvlBA/restApi/internal/app/warehouse"
	"github.com/lvlBA/restApi/internal/controllers/warehouse"
	db "github.com/lvlBA/restApi/internal/db/warehouse"
	warehouseHttp "github.com/lvlBA/restApi/internal/http/warehouse/v1"
	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

func Run(cfg *Config) error {
	// create controller layers
	log, err := cfg.getLogger()
	if err != nil {
		return fmt.Errorf("failed to get logger: %w", err)
	}

	connDB, err := cfg.getDatabaseConnection()
	if err != nil {
		return fmt.Errorf("failed to get db connection: %w", err)
	}
	if err != nil {
		return fmt.Errorf("failed to get db connection: %w", err)
	}

	// controllers
	dbSvc := db.New(connDB)

	warehouseCtrl := warehouse.New(dbSvc)
	warehouseApp := app.New(warehouseCtrl, log)
	warehouseApi := warehouseHttp.NewGoodsService(warehouseApp)

	// fastHttp
	routerFastHttp := fasthttprouter.New()
	routerFastHttp.Handle(api.HttpMethodReceiveGoods, api.UrnApiReceiveGoods, warehouseApi.HandleReceiveGoods)
	routerFastHttp.Handle(api.HttpMethodCreateGoods, api.UrnApiCreateGoods, warehouseApi.HandleCreateGoods)
	routerFastHttp.Handle(api.HttpMethodDeleteGoods, api.UrnApiDeleteGoods, warehouseApi.HandleDeleteGoods)

	conn, err := net.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		return fmt.Errorf("failed to listen %s: %w", cfg.ListenAddress, err)
	}
	serverFastHttp := fasthttp.Server{
		Handler: routerFastHttp.Handler,
	}

	if err := serverFastHttp.Serve(conn); err != nil {
		return fmt.Errorf("error serve connections: %w", err)
	}

	return nil
}
