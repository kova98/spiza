package main

import (
	"github.com/kova98/spiza/services/simulator/domain"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kova98/spiza/services/simulator/adapters"
	"github.com/kova98/spiza/services/simulator/data"
	"github.com/kova98/spiza/services/simulator/handlers"
)

func main() {
	l := log.New(os.Stdout, "simulator-", log.LstdFlags)
	l.Println("Starting simulator")

	connStr := os.Getenv("SPIZA_DB_CONN_STR")
	if connStr == "" {
		l.Fatal("SPIZA_DB_CONN_STR environment variable empty")
	}
	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	if googleApiKey == "" {
		l.Fatal("GOOGLE_API_KEY environment variable empty")
	}

	db := data.InitDb(connStr)
	repo := data.NewRepo(db)
	b := NewBus(l)
	traveler := adapters.NewTraveler(l, b.Client, googleApiKey)
	// TODO: load starting loc from db
	startingLoc := domain.LatLngToLocation("45.801125358549015,15.952160085480502")
	courier := &domain.Courier{Id: "1", Name: "Test Courier", Loc: startingLoc}

	cah := handlers.NewCourierAssignedHandler(l, repo, courier, traveler)
	b.Client.Subscribe("order/+/courier-assigned", 0, cah.Handle)
	ouh := handlers.NewOrderUpdatedHandler(l, repo, courier, traveler)
	b.Client.Subscribe("order/+", 0, ouh.Handle)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	sig := <-c
	log.Println("Got signal:", sig)
}
