package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := log.New(os.Stdout, "monitor-", log.LstdFlags)
	l.Println("Starting monitor")
	connStr := os.Getenv("SPIZA_DB_CONN_STR")
	if connStr == "" {
		l.Fatal("SPIZA_DB_CONN_STR environment variable empty")
	}
	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	if googleApiKey == "" {
		l.Fatal("GOOGLE_API_KEY environment variable empty")
	}

	db := NewPostgresDb(connStr)

	s, err := db.GetCurrentState()
	if err != nil {
		l.Fatal("Unable to initialize state: ", err)
	}

	l.Println(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	sig := <-c
	log.Println("Got signal:", sig)
}
