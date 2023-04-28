package entity

type Goods struct {
	ID    int     `json:"goods_id"`
	Name  string  `json:"good_name"`
	Price float32 `json:"good_price"`
	Type  string  `json:"good_type"`
}
