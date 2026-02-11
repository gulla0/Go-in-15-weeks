package service

import (
	"Hard/ten/internal/store"
	"fmt"
)

type Service struct {
	store store.Store
}

func NewService(s store.Store) *Service {
	return &Service{store: s}
}

func (s Service) AddStock(id int, delta int) error {
	if delta <= 0 {
		return fmt.Errorf("delta must be greater than 0")
	}

	if item, ok := s.store.Get(id); !ok {
		return fmt.Errorf("item not found")
	} else {
		item.Count += delta
		s.store.Put(item)
		return nil
	}
}
