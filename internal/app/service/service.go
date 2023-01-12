package service

import (
	makves "github.com/cucumberjaye/makves_testovoe"
)

type Repository interface {
	GetItems(startId, endId int) ([]makves.Item, error)
	GetItem(id int) (makves.Item, error)
	LoadData(items []makves.Item)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}
