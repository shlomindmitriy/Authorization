package main

import (
	"github.com/Authorization/internal/database/inmemory"
	"github.com/Authorization/internal/service"
)

func main() {
	// Поднять http server
	// создать две эндпоинта
	// 1) v1/auth/login    GET
	// 2) v1/auth/register POST
	db := inmemory.NewInMemory(nil)
	svc := service.NewService(db)
	ha := handler{
		svc: svc,
	}

}

type handler struct {
	svc *service.Service
}

func (h *handler) login()    {}
func (h *handler) register() {}
