package main

import (
	"compelo/api"
	"log"
)

func main() {
	log.Fatal(api.Setup("file::memory:").Run())
}
