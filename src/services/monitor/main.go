package main

import (
	"encoding/json"
	"github.com/kova98/spiza/services/monitor/adapters"
	"github.com/kova98/spiza/services/monitor/domain"
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

	db := adapters.NewPostgresDb(connStr)
	state, err := db.GetCurrentState()
	if err != nil {
		l.Fatal("Unable to initialize state: ", err)
	}

	ws := adapters.NewWebsocketAdapter(l, &state)
	bus := adapters.NewMqttBus(l)
	bus.SubscribeOrderUpdated(ws.HandleOrderUpdated)
	bus.SubscribeOrderCreated(ws.HandleOrderCreated)
	bus.SubscribeCourierAssigned(ws.HandleCourierAssigned)
	bus.SubscribeCourierLocationUpdated(ws.HandleCourierLocationUpdated)

	type Display struct {
		domain.State
		GoogleApiKey string
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl, err := template.ParseFiles("./static/index.gohtml")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, Display{state, googleApiKey})
		if err != nil {
			l.Println("Failed to execute template: ", err)
			http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/ws", ws.HandleWebsocketConnection)
	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		state, err = db.GetCurrentState()
		if err != nil {
			l.Println("Failed to get state: ", err)
			http.Error(w, "Failed to get state", http.StatusInternalServerError)
			return
		}
		err := json.NewEncoder(w).Encode(state)
		if err != nil {
			l.Println("Failed to encode state: ", err)
			http.Error(w, "Failed to encode state", http.StatusInternalServerError)
		}
	})
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		l.Fatal("Failed to start server: ", err)
	}

	l.Println(state)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	sig := <-c
	log.Println("Got signal:", sig)
}
