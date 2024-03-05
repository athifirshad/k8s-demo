package main

import (
	"flag"
	"fmt"

	"github.com/go-chi/chi/v5"
)

type config struct {
	port string
}
type application struct {
	config
	router *chi.Mux
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", "0.0.0.0:4000", "API server port")
	flag.Parse()

	app := &application{
		config: cfg,
		router: chi.NewRouter(),
	}
	app.Routes()
	if err := app.serve(); err != nil {
		fmt.Println("Server failed to start")
	}
}
