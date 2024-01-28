package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kova98/spiza/services/simulator/data"
	"github.com/kova98/spiza/services/simulator/handlers"
	"github.com/kova98/spiza/services/simulator/util"
)

func main() {
	l := log.New(os.Stdout, "simulator-", log.LstdFlags)
	l.Println("Starting simulator")

	connStr := os.Getenv("SPIZA_DB_CONN_STR")
	if connStr == "" {
		l.Fatal("SPIZA_DB_CONN_STR environment variable empty")
	}
	db := data.InitDb(connStr)
	repo := data.NewRepo(db)
	b := NewBus(l)
	traveler := util.NewTraveler(l, b.Client)
	startingLoc := data.LatLngToLocation("45.800169905837784,15.943209331950337")
	courier := &data.Courier{Id: "1", Name: "Test Courier", Loc: startingLoc}

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
