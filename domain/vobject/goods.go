package vobject

type Goods struct {
	Id    int     `json:"id"`
	Name  string  `json:"goods_name"`
	Type  string  `json:"goods_type"`
	Price float32 `json:"goods_price"`
}
