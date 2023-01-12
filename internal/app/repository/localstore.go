package repository

import (
	"errors"
	makves "github.com/cucumberjaye/makves_testovoe"
	"sync"
)

type LocalStore struct {
	Store map[int]makves.Item
	mx    sync.Mutex
}

func New() *LocalStore {
	return &LocalStore{
		Store: make(map[int]makves.Item),
		mx:    sync.Mutex{},
	}
}

func (l *LocalStore) GetItems(startId, endId int) ([]makves.Item, error) {
	var items = []makves.Item{}

	l.mx.Lock()
	defer l.mx.Unlock()

	for startId < endId {
		item, ok := l.Store[startId]
		if !ok {
			startId++
			continue
		}
		items = append(items, item)
		startId++
	}
	if len(items) == 0 {
		return nil, errors.New("empty result")
	}
	return items, nil
}

func (l *LocalStore) GetItem(id int) (makves.Item, error) {
	var item makves.Item

	l.mx.Lock()
	defer l.mx.Unlock()

	item, ok := l.Store[id]
	if !ok {
		return item, errors.New("id does not exist")
	}

	return item, nil
}

func (l *LocalStore) LoadData(items []makves.Item) {
	for _, item := range items {
		l.Store[item.Id] = item
	}
}
