package entity

import (
	"go_trancation_center/domain/vobject"
	"image"
)

type Goods struct {
	ID        int         `json:"goods_id"`
	Name      string      `json:"good_name"`
	Price     float32     `json:"good_price"`
	SellPrice float32     `json:"sellPrice"`
	Image     image.Image `json:"image"`
	// 值对象
	SPU *vobject.SPU `json:"SPU"`
}
