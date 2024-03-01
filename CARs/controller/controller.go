package controller

import (
	"rent_car/storage"
)

type Controller struct {
	Store storage.Store
}

func NewController(store storage.Store) Controller {
	return Controller{
		Store: store,
	}
}
