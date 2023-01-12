package app

import (
	"fmt"
	"github.com/cucumberjaye/makves_testovoe/configs"
	"github.com/cucumberjaye/makves_testovoe/internal/app/handler"
	"github.com/cucumberjaye/makves_testovoe/internal/app/repository"
	"github.com/cucumberjaye/makves_testovoe/internal/app/service"
	"github.com/cucumberjaye/makves_testovoe/internal/pkg/utils"
	"github.com/cucumberjaye/makves_testovoe/pkg/csv"
	"net/http"
)

type App struct {
	h *handler.Handler
	s *service.Service
	r service.Repository
}

func New() (*App, error) {
	a := &App{}

	data, err := csv.CsvToStr()
	if err != nil {
		return nil, err
	}

	items, err := utils.DataToItems(data)
	if err != nil {
		return nil, err
	}

	a.r = repository.New()
	a.r.LoadData(items)

	a.s = service.New(a.r)

	a.h = handler.New(a.s)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	err := configs.LoadConfig()
	if err != nil {
		return err
	}

	host := configs.Domain + ":" + configs.Port

	err = http.ListenAndServe(host, a.h.InitRoutes())
	if err != nil {
		return err
	}

	return nil
}
