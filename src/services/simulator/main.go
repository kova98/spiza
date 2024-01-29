package main

import (
	"github.com/kova98/spiza/services/simulator/domain"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kova98/spiza/services/simulator/adapters"
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

	db := adapters.NewPostgresDb(connStr)
	bus := adapters.NewMqttBus(l)
	maps := adapters.NewGoogleMaps(l, googleApiKey)

	// TODO: load starting loc from db
	startingLoc := domain.LatLngToLocation("45.801125358549015,15.952160085480502")
	courier := &domain.Courier{Id: "1", Name: "Test Courier", Loc: startingLoc}

	cah := handlers.NewCourierAssignedHandler(l, db, courier, maps)
	bus.SubscribeCourierAssigned(cah.Handle)
	ouh := handlers.NewOrderUpdatedHandler(l, db, courier, maps)
	bus.SubscribeOrderUpdated(ouh.Handle)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	sig := <-c
	log.Println("Got signal:", sig)
}
