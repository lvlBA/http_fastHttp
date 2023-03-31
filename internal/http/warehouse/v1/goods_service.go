package warehousehttp

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (s *GoodsServiceImpl) HandleReceiveGoods(ctx *fasthttp.RequestCtx) {
	req, err := s.decodeReceiveGoodsRequest(ctx)
	if err != nil {
		createErrorFastHttpResponse(ctx, fasthttp.StatusInternalServerError, "failed to decode request: %s", err)
		return
	}

	resp, err := s.app.ReceiveGoods(ctx, req)
	if err != nil {
		fastHttpErrorAdapter(ctx, err)
		return
	}

	if err := s.encodeReceiveGoodsResponse(&ctx.Response, resp); err != nil {
		createErrorFastHttpResponse(ctx, fasthttp.StatusInternalServerError, "failed to encode response: %s", err)
		return
	}
}

func (s *GoodsServiceImpl) decodeReceiveGoodsRequest(ctx *fasthttp.RequestCtx) (*api.ReceiveGoodsRequest, error) {
	id, ok := ctx.UserValue("id").(string)
	if !ok {
		return nil, errors.New("id is not set")
	}

	result := api.ReceiveGoodsRequest{
		ID: id,
	}

	return &result, nil
}

func (s *GoodsServiceImpl) encodeReceiveGoodsResponse(r *fasthttp.Response, resp *api.ReceiveGoodsResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	r.SetBody(b)

	return nil
}

func (s *GoodsServiceImpl) HandleCreateGoods(ctx *fasthttp.RequestCtx) {
	req, err := s.decodeCreateGoodsRequest(ctx)
	if err != nil {
		createErrorFastHttpResponse(ctx, fasthttp.StatusInternalServerError, "failed to decode request: %s", err)
		return
	}

	resp, err := s.app.CreateGoods(ctx, req)
	if err != nil {
		fastHttpErrorAdapter(ctx, err)
		return
	}

	if err := s.encodeCreateGoodsResponse(&ctx.Response, resp); err != nil {
		createErrorFastHttpResponse(ctx, fasthttp.StatusInternalServerError, "failed to encode response: %s", err)
		return
	}
}

func (s *GoodsServiceImpl) decodeCreateGoodsRequest(ctx *fasthttp.RequestCtx) (*api.CreateGoodsRequest, error) {
	var result api.CreateGoodsRequest
	err := json.Unmarshal(ctx.Request.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &result, nil
}

func (s *GoodsServiceImpl) encodeCreateGoodsResponse(r *fasthttp.Response, resp *api.CreateGoodsResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	r.SetBody(b)

	return nil
}

func (s *GoodsServiceImpl) HandleDeleteGoods(ctx *fasthttp.RequestCtx) {
	req, err := s.decodeDeleteGoodsRequest(ctx)
	if err != nil {
		createErrorFastHttpResponse(ctx, fasthttp.StatusInternalServerError, "failed to decode request: %s", err)
		return
	}

	resp, err := s.app.DeleteGoods(ctx, req)
	if err != nil {
		fastHttpErrorAdapter(ctx, err)
		return
	}

	if err := s.encodeDeleteGoodsResponse(&ctx.Response, resp); err != nil {
		createErrorFastHttpResponse(ctx, fasthttp.StatusInternalServerError, "failed to encode response: %s", err)
		return
	}
}

func (s *GoodsServiceImpl) decodeDeleteGoodsRequest(ctx *fasthttp.RequestCtx) (*api.DeleteGoodsRequest, error) {
	id, ok := ctx.UserValue("id").(string)
	if !ok {
		return nil, errors.New("id is not set")
	}

	result := api.DeleteGoodsRequest{
		ID: id,
	}

	return &result, nil
}

func (s *GoodsServiceImpl) encodeDeleteGoodsResponse(r *fasthttp.Response, resp *api.DeleteGoodsResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	r.SetBody(b)

	return nil
}
