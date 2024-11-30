package handlers

import (
	"fmt"
	"github.com/Painkiller675/url_shortener_675/internal/repository"
	"github.com/Painkiller675/url_shortener_675/internal/service"
	"io"
	"log"
	"net/http"
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
	// write alias
	// TODO mb GetRandURL should return error too?
	randAl := service.GetRandString(8)
	repository.WriteURL(randAl, string(body))
	// response molding
	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Content-Length", strconv.Itoa(len([]byte(randAl))))
	// TODO add content length?
	res.WriteHeader(http.StatusCreated) // 201
	_, err = res.Write([]byte(randAl))
	if err != nil { // bbrowser can't show 2 statuses
		log.Printf("Error writing to response: %v", err)
		//http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func GetLongURLHandler(res http.ResponseWriter, req *http.Request) {
	// content checking
	//if !(strings.Contains(req.Header.Get("Content-Type"), "text/plain")) {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
	//	return
	//}
	id := req.PathValue("id")
	//if id == "" {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//	return
	//}
	// response molding
	url, err := repository.GetShortURL(id)
	fmt.Println("REQUEST: ", id)
	fmt.Println("REQUEST: ", url)
	if err != nil { //TODO: mb change that to error 400>?
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect) // 307
}
