package main

import (
	"html/template"
	"log"
	"net/http"
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

	state, err := db.GetCurrentState()
	if err != nil {
		l.Fatal("Unable to initialize state: ", err)
	}

	type Display struct {
		State
		GoogleApiKey string
	}

	fs := http.FileServer(http.Dir("./static"))
	// Serve static files for any request not matching the root path
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl, err := template.ParseFiles("./static/index.html")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, Display{state, googleApiKey})
		if err != nil {
			http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":3000", nil)

	l.Println(state)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	sig := <-c
	log.Println("Got signal:", sig)
}
