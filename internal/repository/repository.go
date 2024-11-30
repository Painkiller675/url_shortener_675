package repository

import (
	"fmt"
	"sync"
)

type safeStruct struct { // TODO mb create 2 structures with diff mutexes??
	aliasOrURL map[string]string // 1, mostly for GET
	orURLAlias map[string]string // 2, mostly for POST
	mx1        sync.RWMutex
	mx2        sync.RWMutex
}

func (s *safeStruct) storeValue1(alias string, orURL string) {
	//s.mx1.Lock() //TODO mb 1st I should check everything nd then use mx?
	//defer s.mx1.Unlock()

	s.aliasOrURL[alias] = orURL
	//s.mx1.Unlock()
}

func (s *safeStruct) getValue1(alias string) string {
	//s.mx1.RLock()
	//defer s.mx1.RUnlock()
	if orURL, ok := s.aliasOrURL[alias]; ok {
		s.mx1.RUnlock()
		return orURL
	}
	//s.mx1.RUnlock()
	return "" // TODO add error type?
}

func (s *safeStruct) del1(alias string) {
	//s.mx1.RLock()
	//defer s.mx1.RUnlock()

	delete(s.aliasOrURL, alias)
	//s.mx1.Unlock()
}
func (s *safeStruct) storeValue2(alias string, orURL string) {
	//s.mx2.Lock() //TODO mb 1st I should check everything nd then use mx?
	//defer s.mx2.Unlock()

	s.orURLAlias[alias] = orURL
	//s.mx2.Unlock()
}

func (s *safeStruct) getValue2(orURL string) string {
	//s.mx2.RLock()
	//defer s.mx2.RUnlock()
	if alias, ok := s.orURLAlias[orURL]; ok {
		s.mx2.RUnlock()
		return alias
	}
	//s.mx2.RUnlock()
	return ""
}

func (s *safeStruct) StoreValue1(alias string, orURL string) {
	s.mx1.Lock()
	defer s.mx1.Unlock()
	s.storeValue1(alias, orURL)
}

func (s *safeStruct) StoreValue2(orURL string, alias string) {
	s.mx1.Lock()
	defer s.mx1.Unlock()
	s.storeValue2(orURL, alias)
}

func (s *safeStruct) GetValue1(alias string) string {
	s.mx1.Lock()
	defer s.mx1.Unlock()
	return s.getValue1(alias)
}

func (s *safeStruct) GetValue2(orURL string) string {
	s.mx1.Lock()
	defer s.mx1.Unlock()
	return s.getValue2(orURL)
}

func (s *safeStruct) Del1(alias string) {
	s.mx1.Lock()
	defer s.mx1.Unlock()
	s.del1(alias)
}

var URLStorage safeStruct

func WriteURL(newAl string, newOrURL string) {
	// check if such url already exists if exists => change that
	curAl := URLStorage.GetValue2(newOrURL)
	if curAl != "" {
		URLStorage.Del1(curAl)
		URLStorage.StoreValue1(newAl, newOrURL)
		URLStorage.StoreValue2(newOrURL, newAl)
		fmt.Println("FROM WriteURL (exists):")
		fmt.Println(URLStorage.aliasOrURL)
		fmt.Println(URLStorage.orURLAlias)
	} else { // just store a new orURL cause that's smth new
		URLStorage.StoreValue1(newAl, newOrURL)
		URLStorage.StoreValue2(newOrURL, newAl)
		fmt.Println("FROM WriteURL (add):")
		fmt.Println(URLStorage.aliasOrURL)
		fmt.Println(URLStorage.orURLAlias)
	}
}

func GetShortURL(alias string) (string, error) {
	curOrURL := URLStorage.GetValue1(alias)
	fmt.Println("FROM GET, curOrURL =", curOrURL)
	if curOrURL == "" {
		return "", fmt.Errorf("%v does not exist", alias)
	}
	return curOrURL, nil
}

//url, isExist := UrlStorage[alias]
//fmt.Println("from GET", UrlStorage)
//if !(isExist) {
//	return "", errors.New(fmt.Sprintf("url %s not exist", alias))
//}
//fmt.Println("from GET", UrlStorage)
//return url, nil
