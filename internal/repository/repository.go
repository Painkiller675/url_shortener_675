package repository

import (
	"errors"
	"fmt"
	"strings"
)

var URLStorage = make(map[string]string) // ALIAS - URL

func WriteURL(url string, alias string) {
	isChanged := false
	// check if such url already exists if exists => change that
	for al, ur := range URLStorage {
		if ur == url {
			delete(URLStorage, al)
			URLStorage[alias] = url
			isChanged = true
			break
		}
	}
	// if not exists => add
	if !isChanged {
		URLStorage[alias] = url
	}
	fmt.Println("from write", URLStorage)
}

func GetShortURL(alias string) (string, error) {
	for al := range URLStorage {
		if strings.Contains(al, alias) {
			return al, nil
		}
	}
	return "", errors.New(fmt.Sprintf("alias %s not exists", alias))
}

//url, isExist := UrlStorage[alias]
//fmt.Println("from GET", UrlStorage)
//if !(isExist) {
//	return "", errors.New(fmt.Sprintf("url %s not exist", alias))
//}
//fmt.Println("from GET", UrlStorage)
//return url, nil
