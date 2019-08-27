package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

var (
	listenAddr string
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
	flag.Parse()

	db := new(Memory)

	srv := &api{
		db: db,
	}

	srv = newAPI()

	server := http.Server{
		Addr:         listenAddr,
		Handler:      srv,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("server couldn't start %v", err)
	}

}
