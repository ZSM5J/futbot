package main

import (
	"net/http"

	"./config"
	"./db"
	"./handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"flag"
)

var (
	option = flag.String("o", "", "Choose option: start/migrate")
)

func main() {
	flag.Parse()

	config.LoadConfiguration("./config.json")
	switch op := *option; op {
	case "start":
		startServer()
	case "migrate":
		migrateDB()
	default:
		startServer()
		
	}
}

func startServer() {
	r := mux.NewRouter()
	r.Handle("/", handler.StatusHandler).Methods("GET")
	r.Handle("/", handler.MessageHandler).Methods("POST")
	r.Handle("/api/openbet", handler.AddOpenBet).Methods("GET")
	r.Handle("/api/sendres", handler.CulculateBets).Methods("GET")
	r.Handle("/api/add", handler.AddPoints).Methods("GET")
	r.Handle("/api/setmin", handler.MinBetChanger).Methods("GET")
	r.Handle("/api/setmax", handler.MaxBetChanger).Methods("GET")
	r.Handle("/api/sort", handler.SortPlaces).Methods("GET")
	r.Handle("/api/addnews", handler.AddNews).Methods("GET")
	http.ListenAndServe(":80", handlers.CORS()(r))
}

func migrateDB() {
	db.Migrate()
}
