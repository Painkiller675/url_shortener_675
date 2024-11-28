package repository

import (
	"errors"
	"fmt"
	"strings"
)

var UrlStorage = make(map[string]string) // ALIAS - URL

func WriteURL(url string, alias string) {
	isChanged := false
	// check if such url already exists if exists => change that
	for al, ur := range UrlStorage {
		if ur == url {
			delete(UrlStorage, al)
			UrlStorage[alias] = url
			isChanged = true
			break
		}
	}
	// if not exists => add
	if !isChanged {
		UrlStorage[alias] = url
	}
	fmt.Println("from write", UrlStorage)
}

func GetShortURL(alias string) (string, error) {
	for al := range UrlStorage {
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
