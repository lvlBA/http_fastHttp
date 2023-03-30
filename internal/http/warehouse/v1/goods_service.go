package warehousehttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/valyala/fasthttp"

	app "github.com/lvlBA/restApi/internal/app/warehouse"
	api "github.com/lvlBA/restApi/pkg/warehouse/v1"
)

type GoodsServiceImpl struct {
	app app.Service
}

func NewGoodsService(app app.Service) *GoodsServiceImpl {
	return &GoodsServiceImpl{
		app: app,
	}
}

func (s *GoodsServiceImpl) ServeFastHttp(ctx *fasthttp.RequestCtx) {
	req, err := s.decodeFastHttpRequest(&ctx.Request)
	if err != nil {
		createErrorFastHttpResponse(ctx, fasthttp.StatusInternalServerError, "failed to decode request %s", err)
		return
	}
	resp, err := s.app.Receive(ctx, req)
	if err != nil {
		fastHttpErrorAdapter(ctx, err)
		return
	}
	if err := s.encodeFastHttpResponse(&ctx.Response, resp); err != nil {
		createErrorFastHttpResponse(ctx, fasthttp.StatusInternalServerError, "failed to encode response: %s", err)
		return
	}

}

func (s *GoodsServiceImpl) decodeFastHttpRequest(req *fasthttp.Request) (*api.ReceiveGoodsRequest, error) {
	var result api.ReceiveGoodsRequest
	err := json.Unmarshal(req.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}
	return &result, nil
}

func (s *GoodsServiceImpl) encodeFastHttpResponse(r *fasthttp.Response, resp *api.ReceiveGoodsResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	r.SetBody(b)
	return nil
}

func (s *GoodsServiceImpl) ServeHttp(respHttp http.ResponseWriter, reqHttp *http.Request, params httprouter.Params) {
	ctx := context.Background()

	req, err := s.decodeHttpRequest(reqHttp)
	if err != nil {
		createErrorHttpResponse(respHttp, fasthttp.StatusInternalServerError, "failed to decode request: %s", err)
		return
	}

	resp, err := s.app.Receive(ctx, req)
	if err != nil {
		httpErrorAdapter(respHttp, err)
		return
	}

	if err := s.encodeHttpResponse(respHttp, resp); err != nil {
		createErrorHttpResponse(respHttp, fasthttp.StatusInternalServerError, "failed to encode response: %s", err)
		return
	}
}

func (s *GoodsServiceImpl) decodeHttpRequest(req *http.Request) (*api.ReceiveGoodsRequest, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	var result api.ReceiveGoodsRequest
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}
	return &result, nil
}

func (s *GoodsServiceImpl) encodeHttpResponse(r http.ResponseWriter, resp *api.ReceiveGoodsResponse) (err error) {
	r.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	_, _ = r.Write(b)
	return nil
}
