package main

import (
	"io"
	"net/http"
	"strconv"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	// method checking
	if req.Method != http.MethodPost {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
		return
	}
	// content checking
	if req.Header.Get("Content-Type") != "text/plain" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
		return
	}
	// body checking
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// response molding
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Header().Set("Content-Length", strconv.Itoa(len("http://localhost:8080/EwHXdJfB")))
	// TODO add content length?
	res.WriteHeader(http.StatusCreated) // 201
	_, err = res.Write([]byte("http://localhost:8080/EwHXdJfB"))
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func getURLHandler(res http.ResponseWriter, req *http.Request) {
	// method checking
	if req.Method != http.MethodGet {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// content checking
	if req.Header.Get("Content-Type") != "text/plain" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// id checking
	//id := req.URL.Query().Get("id")
	id := req.PathValue("id")
	if id == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// response molding
	res.Header().Set("Location", "https://practicum.yandex.ru/; charset=utf-8")
	res.WriteHeader(http.StatusTemporaryRedirect) // 307
	_, err := res.Write([]byte("123"))            // TODO is it a CORRECT response???
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/{id}", getURLHandler)
	//mux.HandleFunc("GET /{id}", func(res http.ResponseWriter, req *http.Request) {
	//	id := req.PathValue("id")
	//	fmt.Fprintf(res, "using URL with id=%v\n", id)
	//})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err) // or log.Fatal()???

	}
}
