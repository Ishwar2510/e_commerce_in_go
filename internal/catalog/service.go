package catalog

// the service define how we intercat with the clatlg basically all the operation that are needed to be done
// in cltlaog servuece
type Service struct {
	products map[int]Product
}

func NewService() *Service {
	return &Service{
		products: map[int]Product{
			1: {Id: 1, Name: "iPhone", Price: 8000000, Stock: 10},
			2: {Id: 2, Name: "Laptop", Price: 6500000, Stock: 5},
			3: {Id: 3, Name: "Headphones", Price: 500000, Stock: 20},
		},
	}
}
func (s *Service) List() []Product {
	pl := make([]Product, 0, len(s.products))
	for _, v := range s.products {
		pl = append(pl, v)
	}
	return pl
}

func (s *Service) Get(id int) (Product, bool) {
	p, ok := s.products[id]
	return p, ok
}

func (s *Service) ReduceStock(id, quantity int) bool {
	// here i have a doubt cant w directly get the pointer
	p, ok := s.products[id]
	if !ok || quantity > p.Stock || quantity < 1 {
		return false
	}
	p.Stock -= quantity
	s.products[id] = p
	return true
}
