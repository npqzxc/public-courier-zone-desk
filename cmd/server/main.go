package main

import (
	"log"
	"net/http"

	httpapi "public-courier-zone-desk/internal/http"
	"public-courier-zone-desk/internal/seed"
	"public-courier-zone-desk/internal/store"
)

func main() {
	s := store.New(seed.DefaultRecords())
	log.Fatal(http.ListenAndServe(":8080", httpapi.Router(s)))
}
