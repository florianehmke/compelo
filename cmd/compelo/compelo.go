package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"compelo/internal"
	"compelo/internal/api/handler"
	"compelo/internal/api/router"
	"compelo/internal/api/security"
)

func main() {
	dev := false
	flag.BoolVar(&dev, "dev", false, "dev mode")
	flag.Parse()

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

	svc := compelo.NewService(dbPath)
	hdl := handler.New(svc)
	sec := security.NewJWT(svc, 60*120, secret)
	log.Fatal(http.ListenAndServe(":"+port, router.New(hdl, sec)))
}
