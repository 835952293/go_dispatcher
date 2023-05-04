package entity

import (
	"go_trancation_center/vobject"
	"time"
)

type TermBuying struct {
	ID         int       `json:"termBuyingId"`
	GoodsId    int       `json:"goodId"`
	Goodsname  string    `json:"goodsName"`
	CreateTime time.Time `json:"createTime"`
	// 值对象
	UserList []*vobject.UserInfo `json:"UserList"`
}
