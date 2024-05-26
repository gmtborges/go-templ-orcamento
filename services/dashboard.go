package services

import "github.com/gustavomtborges/orcamento-auto/stores"

type DashService struct {
	userStore stores.UserStorer
}

func NewDashService(userStore stores.UserStorer) *DashService {
	return &DashService{userStore: userStore}
}
