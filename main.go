package main

import (
	"flag"
	"log"
	"net/http"

	"echo-server/handler"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", "localhost:8080", "http service address")

	flag.Parse()
	log.SetFlags(0)
}

func main() {
	h := handler.NewHandler()

	http.HandleFunc("/ws", h.Connect)

	log.Printf("Server started on: ws://%s/ws\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
