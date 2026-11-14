package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Ishwar2510/e_commerce_in_go/internal/catalog"
	"github.com/Ishwar2510/e_commerce_in_go/internal/order"
	"github.com/Ishwar2510/e_commerce_in_go/internal/user"
)

type App struct {
	users   *user.Service
	catalog *catalog.Service
	orders  *order.Service
	// Simple protection for our in-memory demo.
	mu sync.Mutex
}

// /******* HTTP HANDLER
// create user api
// consider these as api calls
func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" || req.Email == "" {
		fmt.Println("there is an error")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	newUser := a.users.Create(req.Name, req.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getProduct api
func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	productList := a.catalog.List()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(productList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// createOrder api
func (a *App) createOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserId    int `json:"user_id"`
		ProductId int `json:"product_id"`
		Quantity  int
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid order details", http.StatusBadRequest)
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	newOrder, err := a.orders.Buy(
		req.UserId,
		req.ProductId,
		req.Quantity,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// the above functions defines the api call

func main() {

	// Initialize business modules.
	users := user.NewService()
	products := catalog.NewService()
	orders := order.NewService(users, products)

	// Initialize application.
	app := &App{
		users:   users,
		catalog: products,
		orders:  orders,
	}
	// Register HTTP routes.

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", app.createUser)
	mux.HandleFunc("GET /products", app.getProducts)
	mux.HandleFunc("POST /orders", app.createOrder)
	PORT := 8082

	log.Printf("E-commerce server running on :%d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
