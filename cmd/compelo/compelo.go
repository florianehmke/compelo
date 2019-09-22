package main

import (
	"flag"
	"log"
	"os"

	"compelo/api"
)

func main() {
	dev := false
	flag.BoolVar(&dev, "dev", false, "dev mode")
	flag.Parse()

	secret, ok := os.LookupEnv("COMPELO_SECRET")
	if !ok {
		if !dev {
			log.Fatal("COMPELO_SECRET environment variable is required.")
		} else {
			secret = "unsecure_dev_secret"
		}
	}

	dbPath, ok := os.LookupEnv("COMPELO_DB_PATH")
	if !ok {
		dbPath = "db.sql"
		log.Println("No COMPELO_DB_PATH environment variable present.")
		log.Println("Using default value instead: 'db.sql'.")
	}
	port, ok := os.LookupEnv("COMPELO_PORT")
	if !ok {
		port = "8080"
		log.Println("No COMPELO_PORT environment variable present.")
		log.Println("Using default value instead: '8080'.")
	}

	log.Fatal(api.Setup(dbPath, secret).Run(":" + port))
}
