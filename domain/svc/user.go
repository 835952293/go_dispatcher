package svc

type UserSvc interface {
}

type userSvc struct {
}

func NewUserSvc() UserSvc {
	return &userSvc{}
}
