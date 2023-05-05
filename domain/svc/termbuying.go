package svc

type TermBuyingSvc interface {
}

type termBuyingSvc struct {
}

func NewTermBuyingSvc() TermBuyingSvc {
	return &termBuyingSvc{}
}
