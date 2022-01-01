package main

import (
	"log"
	"net/http"
	"os"

	"compelo/api/handler"
	"compelo/api/router"
	"compelo/api/security"
	"compelo/command"
	"compelo/event"
	"compelo/query"
)

func main() {
	secret, ok := os.LookupEnv("COMPELO_SECRET")
	if !ok {
		secret = "insecure_dev_secret"
		log.Println("COMPELO_SECRET environment variable is missing.")
		log.Println("Using 'insecure_dev_secret' instead.")
	}
	dbPath, ok := os.LookupEnv("COMPELO_DB_PATH")
	if !ok {
		dbPath = "db.sql"
		log.Println("COMPELO_DB_PATH environment variable is missing.")
		log.Println("Using default value instead: 'db.sql'.")
	}
	port, ok := os.LookupEnv("COMPELO_PORT")
	if !ok {
		port = "8080"
		log.Println("COMPELO_PORT environment variable is missing.")
		log.Println("Using default value instead: '8080'.")
	}

	bus := event.NewBus()
	store := event.NewStore(bus, dbPath)
	query := query.NewService(bus)

	// Load events, panic if that does not work.
	events, err := store.LoadEvents()
	if err != nil {
		log.Panicln(err)
	}

	command := command.NewService(store, events)
	handler := handler.New(query, command)
	security := security.New(query, 60*120, secret)

	log.Fatal(http.ListenAndServe(":"+port, router.New(handler, security)))
}
