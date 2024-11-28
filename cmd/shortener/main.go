package main

import (
	"github.com/Painkiller675/url_shortener_675/internal/config"
	"github.com/Painkiller675/url_shortener_675/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// init config
	cfg := config.MustLoad()

	// init router
	r := chi.NewRouter()

	// routing
	r.Route("/", func(r chi.Router) {
		r.Post("/", handlers.CreateShortURLHandler)
		r.Get("/{id}", handlers.GetLongURLHandler)
	})

	//start server
	err := http.ListenAndServe(cfg.Address, r)
	if err != nil {
		panic(err) // or log.Fatal()???

	}
}
