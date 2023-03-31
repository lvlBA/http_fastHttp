package warehouseapi

import "net/http"

const (
	UrnApiCreateGoods  = UrnApiPrefix + "/goods/create"
	UrnApiReceiveGoods = UrnApiPrefix + "/goods/receive/:id"
	UrnApiDeleteGoods  = UrnApiPrefix + "/goods/delete/:id"

	HttpMethodReceiveGoods = http.MethodGet
	HttpMethodCreateGoods  = http.MethodPost
	HttpMethodDeleteGoods  = http.MethodDelete
)

type Goods struct {
	ID   string `json:"id"   xml:"id"`
	Name string `json:"name" xml:"name"`
}

type CreateGoodsRequest struct {
	Name string `json:"name" xml:"name"`
}

type CreateGoodsResponse struct {
	Goods *Goods `json:"goods" xml:"goods"`
}

type ReceiveGoodsRequest struct {
	ID string `json:"ID" xml:"ID"`
}

type ReceiveGoodsResponse struct {
	Goods *Goods `json:"goods" xml:"goods"`
}

type DeleteGoodsRequest struct {
	ID string `json:"ID" xml:"ID"`
}

type DeleteGoodsResponse struct {
}
