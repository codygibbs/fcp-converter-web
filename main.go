package main

import (
	"log"
	"net/http"
	"time"

	"github.com/codygibbs/converter-web/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Home)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:9002",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
