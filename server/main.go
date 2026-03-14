package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"ShoppingCalculator/helper"
	"ShoppingCalculator/int/src"
)

type itemInput struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type checkoutRequest struct {
	Customer string      `json:"customer"`
	Items    []itemInput `json:"items"`
}

type checkoutResponse struct {
	TotalQuantity int                     `json:"totalQuantity"`
	TotalCost     float64                 `json:"totalCost"`
	Suggestions   []helper.ProductStorage `json:"suggestions"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/checkout", withCORS(handleCheckout))
	mux.HandleFunc("/api/suggestions", withCORS(handleSuggestions))
	mux.Handle("/", http.FileServer(http.Dir("web")))

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Println("ShoppingCalculator server running at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server failed: %v", err)
	}
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func handleCheckout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}

	var payload checkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid JSON body"})
		return
	}

	items, products, err := normalizeItems(payload.Items)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	result := src.Calc(items)
	if err := src.Jupdate(products); err != nil {
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "failed to update shopping log"})
		return
	}

	suggestions, _ := loadSuggestions(5)

	writeJSON(w, http.StatusOK, checkoutResponse{
		TotalQuantity: result.TotalQuantity,
		TotalCost:     result.TotalCost,
		Suggestions:   suggestions,
	})
}

func handleSuggestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}

	suggestions, err := loadSuggestions(5)
	if err != nil {
		writeJSON(w, http.StatusOK, []helper.ProductStorage{})
		return
	}

	writeJSON(w, http.StatusOK, suggestions)
}

func normalizeItems(items []itemInput) ([]helper.Input, []helper.ProductStorage, error) {
	if len(items) == 0 {
		return nil, nil, errors.New("add at least one item")
	}

	inputs := make([]helper.Input, 0, len(items))
	products := make([]helper.ProductStorage, 0, len(items))

	for _, item := range items {
		name := strings.TrimSpace(item.Name)
		if name == "" {
			return nil, nil, errors.New("item name is required")
		}
		if item.Quantity <= 0 {
			return nil, nil, errors.New("quantity must be greater than 0")
		}
		if item.Price <= 0 {
			return nil, nil, errors.New("price must be greater than 0")
		}

		inputs = append(inputs, helper.Input{
			ItemName:      name,
			NumberOfItems: item.Quantity,
			PriceOfItem:   item.Price,
			Cost:          float64(item.Quantity) * item.Price,
		})

		products = append(products, helper.ProductStorage{
			Name:  name,
			Price: item.Price,
			Freq:  0,
		})
	}

	return inputs, products, nil
}

func loadSuggestions(limit int) ([]helper.ProductStorage, error) {
	logFilePath := logPath()
	file, err := os.Open(logFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []helper.ProductStorage
	if err := json.NewDecoder(file).Decode(&products); err != nil {
		return nil, err
	}

	if limit > 0 && len(products) > limit {
		products = products[:limit]
	}
	return products, nil
}

func logPath() string {
	if override := os.Getenv("SHOPPING_LOG_PATH"); override != "" {
		return override
	}
	return filepath.Join("storage", "shopinglogs.json")
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
