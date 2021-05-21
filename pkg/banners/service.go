package banners

import (
	"context"
	"errors"
	"sync"
)

type Service struct {
	mu    sync.RWMutex
	items []*Banner
}

func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
	Image   string
}

func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items, nil
}

func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}
	return nil, errors.New("item not found")
}

var curID int64 = 1

func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	if item.ID == 0 {
		curID++
		item.ID = curID
		s.items = append(s.items, item)
		return item, nil
	}
	for i := 0; i < len(s.items); i++ {
		if s.items[i].ID == item.ID {
			s.items[i] = item
			return item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	banner, err := s.ByID(ctx, id)
	if err != nil {
		return nil, err
	}
	index := -1
	s.mu.RLock()
	for i := 0; i < len(s.items); i++ {
		if s.items[i] == banner {
			index = i
			break
		}
	}
	s.items = append(s.items[:index], s.items[index+1:]...)
	s.mu.RUnlock()
	return banner, nil
}
