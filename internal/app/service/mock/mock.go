package mock

import (
	"errors"
	makves "github.com/cucumberjaye/makves_testovoe"
)

type Mock struct {
}

func (s *Mock) GetItems(startId, endId int) ([]makves.Item, error) {
	if startId == 0 {
		return nil, errors.New("test")
	}

	return make([]makves.Item, 0), nil
}

func (s *Mock) GetItem(id int) (makves.Item, error) {
	if id == 0 {
		return makves.Item{}, errors.New("test")
	}

	return makves.Item{}, nil
}
