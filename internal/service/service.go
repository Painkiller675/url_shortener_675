package service

import (
	"log"
	"math/rand"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandURL(n int, baseURL string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}

	resultURL, err := url.JoinPath(baseURL, string(b))
	if err != nil {
		log.Fatal(err)
	}

	return resultURL
}
