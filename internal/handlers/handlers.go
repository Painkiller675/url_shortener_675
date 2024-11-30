package handlers

import (
	"github.com/Painkiller675/url_shortener_675/internal/repository"
	"github.com/Painkiller675/url_shortener_675/internal/service"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func CreateShortURLHandler(res http.ResponseWriter, req *http.Request) {
	// content checking
	//if !(strings.Contains(req.Header.Get("Content-Type"), "text/plain")) {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
	//	return
	//}
	// body checking
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// write aliass
	// TODO mb GetRandURL should return error too?
	randAl := service.GetRandString(8)
	repository.WriteURL(randAl, string(body))
	// response molding
	baseURL := "http://localhost:8080/"
	resultURL, err := url.JoinPath(baseURL, randAl)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Content-Length", strconv.Itoa(len([]byte(resultURL))))
	// TODO add content length?
	res.WriteHeader(http.StatusCreated) // 201
	_, err = res.Write([]byte(resultURL))
	if err != nil {
		log.Printf("Error writing to response: %v", err)
		return
	}
}

func GetLongURLHandler(res http.ResponseWriter, req *http.Request) {
	// content checking
	//if !(strings.Contains(req.Header.Get("Content-Type"), "text/plain")) {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
	//	return
	//}
	id := req.PathValue("id") // the cap
	//_ = id                    // the cap
	//if id == "" {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//	return
	//}
	// response molding ...
	orURL, err := repository.GetShortURL(id)
	//fmt.Println("REQUEST: ", id)
	//fmt.Println("REQUEST: ", url)
	if err != nil { //TODO: mb change that to error 400>?
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", orURL)
	res.WriteHeader(http.StatusTemporaryRedirect) // 307
}
