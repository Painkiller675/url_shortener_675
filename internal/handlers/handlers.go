package handlers

import (
	"github.com/Painkiller675/url_shortener_675/internal/service"
	"io"
	"log"
	"net/http"
	"strconv"
)

var origURL string

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
	defer req.Body.Close() // TODO here it might be completed by the lib? => mb del that
	origURL = string(body)
	// write aliass
	// TODO mb GetRandURL should return error too?
	randURL := service.GetRandURL(8, origURL)
	//repository.WriteURL(string(body), randURL)
	// response molding
	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Content-Length", strconv.Itoa(len([]byte(randURL))))
	// TODO add content length?
	res.WriteHeader(http.StatusCreated) // 201
	_, err = res.Write([]byte(randURL))
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
	u := req.URL
	_ = u
	//id := req.PathValue("id") // the cap
	//_ = id                    // the cap
	//if id == "" {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//	return
	//}
	// response molding
	//url, err := repository.GetShortURL(id)
	//fmt.Println("REQUEST: ", id)
	//fmt.Println("REQUEST: ", url)
	//if err != nil { //TODO: mb change that to error 400>?
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//	return
	//}
	res.Header().Set("Location", origURL)
	res.WriteHeader(http.StatusTemporaryRedirect) // 307
}
