package svc

type GoodsSvc interface {
}

type goodsSvc struct {
}

func NewGoodsSvc() GoodsSvc {
	return &goodsSvc{}
}
