package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"helloserver/app"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}

func run() error {
	port := flag.Int("port", 8888, "port to listen on")

	server, err := app.NewServer(*port)
	if err != nil {
		log.Fatalf("Error creating server: %s", err.Error())
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", *port), server.Router)
}
