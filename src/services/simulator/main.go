package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kova98/spiza/services/simulator/data"
)

func main() {
	l := log.New(os.Stdout, "simulator-", log.LstdFlags)
	l.Println("Starting simulator")

	connStr := os.Getenv("SPIZA_DB_CONN_STR")
	if connStr == "" {
		l.Fatal("SPIZA_DB_CONN_STR environment variable empty")
	}
	db := data.InitDb(connStr)
	repo := data.NewOrderRepo(db)

	b := NewBus(l)
	b.Client.Subscribe("order/+/courier-assigned", 0, NewCourierAssignedHandler(l, repo))
	b.Client.Publish("test", 0, false, "Hello World!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	sig := <-c
	log.Println("Got signal:", sig)
}
