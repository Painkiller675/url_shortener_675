package handlers

import (
	"fmt"
	"github.com/Painkiller675/url_shortener_675/internal/repository"
	"github.com/Painkiller675/url_shortener_675/internal/service"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func CreateShortURLHandler(res http.ResponseWriter, req *http.Request) {
	// content checking
	if !(strings.Contains(req.Header.Get("Content-Type"), "text/plain")) {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
		return
	}
	// body checking
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// write alias
	// TODO mb GetRandURL should return error too?
	randURL := service.GetRandURL(8, string(body))
	repository.WriteURL(string(body), randURL)
	// response molding
	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Content-Length", strconv.Itoa(len(randURL)))
	// TODO add content length?
	res.WriteHeader(http.StatusCreated) // 201
	_, err = res.Write([]byte(randURL))
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func GetLongURLHandler(res http.ResponseWriter, req *http.Request) {
	// content checking
	if !(strings.Contains(req.Header.Get("Content-Type"), "text/plain")) {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
		return
	}
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
