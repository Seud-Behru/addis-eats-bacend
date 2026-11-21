package menu

// memoryStore holds dishes in a slice
type memoryStore struct {
	dishes []Dish
}

// NewMemoryStore builds the store with the eight dishes
func NewMemoryStore() *memoryStore {
	return &memoryStore{
		dishes: []Dish{
			{ID: "d1", Name: "Shiro", PriceSantim: 12000, IsFasting: true},
			{ID: "d2", Name: "Doro Wat", PriceSantim: 25000, IsFasting: false},
			{ID: "d3", Name: "Kitfo", PriceSantim: 30000, IsFasting: false},
			{ID: "d4", Name: "Misir Wat", PriceSantim: 10000, IsFasting: true},
			{ID: "d5", Name: "Gomen", PriceSantim: 8000, IsFasting: true},
			{ID: "d6", Name: "Tibs", PriceSantim: 22000, IsFasting: false},
			{ID: "d7", Name: "Injera", PriceSantim: 500, IsFasting: true},
			{ID: "d8", Name: "Tej", PriceSantim: 15000, IsFasting: false},
		},
	}
}

// All returns every dish
func (s *memoryStore) All() ([]Dish, error) {
	return s.dishes, nil
}

// ByID finds one dish by ID
func (s *memoryStore) ByID(id string) (Dish, error) {
	for _, d := range s.dishes {
		if d.ID == id {
			return d, nil
		}
	}
	return Dish{}, ErrNotFound
}