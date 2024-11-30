package main

import (
	"github.com/Painkiller675/url_shortener_675/internal/config"
	"github.com/Painkiller675/url_shortener_675/internal/handlers"
	"net/http"
)

func main() {
	// init config
	cfg := config.MustLoad()

	// init router
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handlers.CreateShortURLHandler)
	mux.HandleFunc("GET /{id}", handlers.GetLongURLHandler)

	//start server
	err := http.ListenAndServe(cfg.Address, mux)
	if err != nil {
		panic(err) // or log.Fatal()???

	}
}
