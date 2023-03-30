package app

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/buaazp/fasthttprouter"
	"github.com/julienschmidt/httprouter"
	"github.com/valyala/fasthttp"

	app "github.com/lvlBA/restApi/internal/app/warehouse"
	warehouseHttp "github.com/lvlBA/restApi/internal/http/warehouse/v1"
	"github.com/lvlBA/restApi/internal/warehouse"
	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

func Run(cfg *Config) error {
	// create controller layers
	warehouseCtrl := warehouse.New()
	warehouseApp := app.New(warehouseCtrl)
	warehouseApi := warehouseHttp.NewGoodsService(warehouseApp)

	wg := new(sync.WaitGroup)

	// http
	conn80, err := net.Listen("tcp", cfg.ListenAddress2)
	if err != nil {
		return fmt.Errorf("failed to listen %s: %w", cfg.ListenAddress2, err)
	}

	routerHttp := httprouter.New()
	routerHttp.Handle(api.HttpMethodReceiveGoods, api.UrnApiReceiveGoods, warehouseApi.ServeHttp)
	serverHttp := http.Server{
		Addr:    cfg.ListenAddress2,
		Handler: routerHttp,
	}
	wg.Add(1)
	go serve(wg, conn80, serverHttp.Serve)

	// fasthttp
	routerFastHttp := fasthttprouter.New()
	routerFastHttp.Handle(api.HttpMethodReceiveGoods, api.UrnApiReceiveGoods, warehouseApi.ServeFastHttp)

	conn8080, err := net.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		return fmt.Errorf("failed to listen %s: %w", cfg.ListenAddress, err)
	}
	serverFastHttp := fasthttp.Server{
		Handler: routerFastHttp.Handler,
	}

	wg.Add(1)
	go serve(wg, conn8080, serverFastHttp.Serve)

	wg.Wait()

	return nil
}

func serve(wg *sync.WaitGroup, l net.Listener, handle func(listener net.Listener) error) {
	defer wg.Done()

	if err := handle(l); err != nil {
		fmt.Printf("failed to serve: %s", err)
		os.Exit(1)
	}
}
