package handler

import (
	makves "github.com/cucumberjaye/makves_testovoe"
	"github.com/go-chi/chi/v5"
)

type Service interface {
	GetItems(startId, endId int) ([]makves.Item, error)
	GetItem(id int) (makves.Item, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/get-items/{id}", h.GetItem)
	r.Get("/get-items/{start_id}-{end_id}", h.GetItems)

	return r
}
