package handlers

import (
	"github.com/Painkiller675/url_shortener_675/internal/service"
	"io"
	"net/http"
	"strconv"
)

func CreateShortURLHandler(res http.ResponseWriter, req *http.Request) {
	// method checking TODO mb we should uncomment that to get error 400
	//if req.Method != http.MethodPost {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
	//	return
	//}
	// content checking
	//if req.Header.Get("Content-Type") != "text/plain" {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //400
	//	return
	//}
	// body checking
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// response molding
	res.Header().Set("Content-Type", "text/plain")
	// TODO mb GetRandURL should return error too?
	randURL := service.GetRandURL(8, "http://localhost:8080/")
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
	// method checking TODO mb we should uncomment that to get error 400
	//if req.Method != http.MethodGet {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//	return
	//}
	// content checking
	//if req.Header.Get("Content-Type") != "text/plain" {
	//	http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//	return
	//}
	// id checking
	//id := req.URL.Query().Get("id")
	id := req.PathValue("id")
	if id == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// response molding
	res.Header().Set("Location", "https://practicum.yandex.ru/")
	res.WriteHeader(http.StatusTemporaryRedirect) // 307

}
