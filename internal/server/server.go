package server

import (
	"WithGres/internal/configs"
	"WithGres/internal/services"
	"fmt"
	"net/http"
)

func Start() error {
	cfg := configs.NewConfig()

	configureRouter()

	s := &http.Server{
		Addr: cfg.ServerAddress,
	}

	fmt.Printf("Server is Working on %s port", cfg.ServerAddress)

	return s.ListenAndServe()
}

func configureRouter() {
	http.HandleFunc("/markets", services.FindAllMarkets)
	http.HandleFunc("/markets/one", services.FindMarketById)
	http.HandleFunc("/markets/add", services.CreateMarket)
	http.HandleFunc("/markets/update", services.UpdateMarket)
	http.HandleFunc("/markets/delete", services.DeleteMarket)

	http.HandleFunc("/items", services.FindAllItems)
	 http.HandleFunc("/items/one", services.FindItemById)
	http.HandleFunc("/items/add", services.CreateItem)
	http.HandleFunc("/items/update", services.UpdateItem)
	http.HandleFunc("/items/delete", services.DeleteItem)
}
