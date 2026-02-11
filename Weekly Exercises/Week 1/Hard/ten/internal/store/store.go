package store

type Item struct {
	ID    int
	Name  string
	Count int
}

type Store interface {
	Get(id int) (Item, bool)
	Put(Item) error
}

type InMemoryStore struct {
	items map[int]Item
}

func (s InMemoryStore) Get(id int) (Item, bool) {
	item, ok := s.items[id]
	return item, ok
}

func (s *InMemoryStore) Put(item Item) error {
	if s.items == nil {
		s.items = make(map[int]Item)
	}
	s.items[item.ID] = item
	return nil
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{items: make(map[int]Item)}
}

type FakeStore struct {
	items map[int]Item
}

func NewFakeStore() *FakeStore {
	return &FakeStore{items: make(map[int]Item)}
}

func (s FakeStore) Get(id int) (Item, bool) {
	item, ok := s.items[id]
	return item, ok
}

func (s *FakeStore) Put(item Item) error {
	if s.items == nil {
		s.items = make(map[int]Item)
	}
	s.items[item.ID] = item
	return nil
}
