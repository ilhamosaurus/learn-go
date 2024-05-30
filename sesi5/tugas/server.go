package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var PORT = 5002

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []*Product{
	{ID: 1, Name: "Keyboard", Price: 10000},
	{ID: 2, Name: "Mouse", Price: 5000},
	{ID: 3, Name: "Monitor", Price: 20000},
	{ID: 4, Name: "CPU", Price: 50000},
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()

	http.HandleFunc("/products", getProducts)

	http.HandleFunc("/products/create", createProduct)

	http.HandleFunc("/products/update", updateProduct)

	http.HandleFunc("/products/delete", deleteProduct)

	log.Println("Server started on port", PORT)

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if p.Name == "" || p.Price == 0 {
		http.Error(w, "name and price are required", http.StatusBadRequest)
		return
	}

	product := &Product{
		ID:    len(products) + 1,
		Name:  p.Name,
		Price: p.Price,
	}
	products = append(products, product)
	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == id {
			var p Product
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if p.Name == "" && p.Price == 0 {
				http.Error(w, "name and price are required", http.StatusBadRequest)
				return
			}

			products[i] = &Product{
				ID:    product.ID,
				Name:  p.Name,
				Price: p.Price,
			}

			json.NewEncoder(w).Encode(products)
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			json.NewEncoder(w).Encode(products)
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}
