package warehousehttp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/lvlBA/restApi/internal/app/warehouse"
)

func createErrorFastHttpResponse(ctx *fasthttp.RequestCtx, status int, msg string, values ...interface{}) {
	ctx.SetStatusCode(status)
	ctx.SetBodyString(fmt.Sprintf(msg, values...))
}

func fastHttpErrorAdapter(ctx *fasthttp.RequestCtx, err error) {
	if err == nil {
		return
	}
	code, msg := errorAdapter(err)
	ctx.SetStatusCode(code)
	ctx.SetBodyString(msg)
}

func createErrorHttpResponse(resp http.ResponseWriter, status int, msg string, values ...interface{}) {
	resp.WriteHeader(status)
	_, _ = resp.Write([]byte(fmt.Sprintf(msg, values...)))
}

func httpErrorAdapter(resp http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	code, msg := errorAdapter(err)
	resp.WriteHeader(code)
	_, _ = resp.Write([]byte(msg))
}

func errorAdapter(err error) (code int, msg string) {
	switch {
	case errors.Is(err, warehouseapp.ErrorNotFound):
		code = fasthttp.StatusNotFound
		msg = err.Error()
	case errors.Is(err, warehouseapp.ErrorAlreadyExists):
		code = fasthttp.StatusConflict
		msg = err.Error()
	case errors.Is(err, warehouseapp.ErrorInternal):
		code = fasthttp.StatusInternalServerError
		msg = err.Error()
	}

	return
}
