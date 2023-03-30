package warehouseapi

import "net/http"

const (
	UrnApiReceiveGoods = urnApiPrefix + "/goods/receive"
	UrnApiSendGoods    = urnApiPrefix + "/goods/send/:id"

	HttpMethodReceiveGoods = http.MethodPost
	HttpMethodSendGoods    = http.MethodDelete
)

type Goods struct {
	ID   string `json:"id"   xml:"id"`
	Name string `json:"name" xml:"name"`
}

type ReceiveGoodsRequest struct {
	Name string `json:"name" xml:"name"`
}

type ReceiveGoodsResponse struct {
	Goods *Goods `json:"goods" xml:"goods"`
}
