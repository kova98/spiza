package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/kova98/spiza/services/api/data"
	"github.com/kova98/spiza/services/api/handlers"
)

func main() {
	l := log.New(os.Stdout, "services-api", log.LstdFlags)
	// get the connection string from the environment variable
	connStr := os.Getenv("SPIZA_DB_CONN_STR")
	if connStr == "" {
		l.Fatal("SPIZA_DB_CONN_STR environment variable empty")
	}
	db := data.InitDb(connStr)
	restaurantRepo := data.NewRestaurantRepo(db)
	itemRepo := data.NewItemRepo(db)
	menuCategoryRepo := data.NewMenuCategoryRepo(db)
	orderRepo := data.NewOrderRepo(db)

	rh := handlers.NewRestaurantsHandler(l, restaurantRepo)
	ih := handlers.NewItemHandler(l, itemRepo)
	mch := handlers.NewMenuCategoryHandler(l, menuCategoryRepo)
	oh := handlers.NewOrderHandler(l, orderRepo)

	router := mux.NewRouter()
	router.Use(handlers.CorsMiddleware)
	router.Use(handlers.CommonMiddleware)

	getRouter := router.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	getRouter.HandleFunc("/api/restaurant", rh.GetRestaurants)
	getRouter.HandleFunc("/api/restaurant/{id}", rh.GetRestaurant)
	getRouter.HandleFunc("/api/restaurant/{id}/order", oh.GetOrders)
	getRouter.HandleFunc("/api/restaurant/{id}/order-ws", oh.HandleOrderWebSocket)

	postRouter := router.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	postRouter.HandleFunc("/api/restaurant", rh.CreateRestaurant)
	postRouter.HandleFunc("/api/menu-category/{id}/item", ih.CreateItem)
	postRouter.HandleFunc("/api/menu-category", mch.CreateMenuCategory)
	postRouter.HandleFunc("/api/order", oh.CreateOrder)

	deleteRouter := router.Methods(http.MethodDelete, http.MethodOptions).Subrouter()
	deleteRouter.HandleFunc("/api/menu-category/{id}", mch.DeleteMenuCategory)
	deleteRouter.HandleFunc("/api/restaurant/{id}", rh.DeleteRestaurant)
	deleteRouter.HandleFunc("/api/item/{id}", ih.DeleteItem)

	putRouter := router.Methods(http.MethodPut, http.MethodOptions).Subrouter()
	putRouter.HandleFunc("/api/restaurant", rh.UpdateRestaurant)

	addr := "127.0.0.1:5002"

	s := http.Server{
		Addr:         addr,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting server on address " + addr)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
	cancel()
}
