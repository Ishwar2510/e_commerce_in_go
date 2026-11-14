package order

import (
	"errors"

	"github.com/Ishwar2510/e_commerce_in_go/internal/catalog"
	"github.com/Ishwar2510/e_commerce_in_go/internal/user"
)

type Service struct {
	user    *user.Service
	catalog *catalog.Service
	order   []Order
	nextId  int
}

func NewService(user *user.Service, catalog *catalog.Service) *Service {
	return &Service{
		user:    user,
		catalog: catalog,
		nextId:  1,
	}
}

func (s *Service) Buy(userId, productId, quantity int) (Order, error) {
	// check if user exist or not
	if _, exists := s.user.Get(userId); !exists {
		return Order{}, errors.New("user not found")
	}

	product, exist := s.catalog.Get(productId)
	if !exist {
		return Order{}, errors.New("product not found")
	}

	if product.Stock < quantity {
		return Order{}, errors.New("product is out of stock")
	}

	// reduce quantity before creating the order
	if ok := s.catalog.ReduceStock(productId, quantity); !ok {
		return Order{}, errors.New("unable to reduce stock")
	}

	newOrder := Order{
		Id:         s.nextId,
		UserId:     userId,
		ProductId:  product.Id,
		Quantity:   quantity,
		TotalValue: product.Price * float64(quantity),
		Status:     "Placed",
	}
	s.order = append(s.order, newOrder)
	s.nextId++
	return newOrder, nil

}

// This is our first business workflow: an order service uses the User and Catalog modules.

//**********************************
// Notice that Order does not directly access their internal maps. It communicates through their services.
