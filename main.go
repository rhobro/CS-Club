package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/rhobro/csclub/handler"
)

func main() {
	log.Println("starting server")
	err := http.ListenAndServe(":8080", handler.Root())
	if err != nil {
		log.Fatalf("http: %s", err)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
