package service

import (
	"errors"
	makves "github.com/cucumberjaye/makves_testovoe"
)

func (s *Service) GetItems(startId, endId int) ([]makves.Item, error) {
	if startId > endId {
		return nil, errors.New("end_id < start_id")
	}
	if startId < 0 || endId < 0 {
		return nil, errors.New("id can not be negative")
	}

	return s.repo.GetItems(startId, endId)
}

func (s *Service) GetItem(id int) (makves.Item, error) {
	var item makves.Item

	if id < 0 {
		return item, errors.New("id can not be negative")
	}

	return s.repo.GetItem(id)
}
