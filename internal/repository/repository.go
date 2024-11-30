package repository

import (
	"fmt"
)

// no uniqueness
var URLStorage = make(map[string]string) // ALIAS - orURL

func WriteURL(newAl string, newOrURL string) {
	// check if such url already exists if exists => change that
	URLStorage[newAl] = newOrURL
	fmt.Println("from write", URLStorage)
}

func GetShortURL(alias string) (string, error) {
	if curOrURL, ok := URLStorage[alias]; ok {
		return curOrURL, nil
	}
	return "", fmt.Errorf("%v does not exist", alias)
}

//url, isExist := UrlStorage[alias]
//fmt.Println("from GET", UrlStorage)
//if !(isExist) {
//	return "", errors.New(fmt.Sprintf("url %s not exist", alias))
//}
//fmt.Println("from GET", UrlStorage)
//return url, nil
