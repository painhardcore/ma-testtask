package httpsearcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

//HTTPSearcher is for searching some string count on web by http
type HTTPSearcher struct {
	total   int64
	limiter chan struct{}

	sync.RWMutex
	wg sync.WaitGroup
}

//NewWithLimit creates new HTTPSearcher limiting concurrent searches with input
func NewWithLimit(limit int64) *HTTPSearcher {
	s := new(HTTPSearcher)
	s.limiter = make(chan struct{}, limit)
	return s

}

//Search validate and launch search process in goroutine
func (s *HTTPSearcher) Search(url, word string) {
	//simple validation
	if url == "" || word == "" {
		fmt.Printf("[Error] Missing url or search word (url = %s , word = %s)\n", url, word)
		return
	}

	s.limiter <- struct{}{}
	s.wg.Add(1)

	go s.search(url, word)
	return
}
func (s *HTTPSearcher) search(url, word string) {

	defer s.done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Count for %s: 0 [error][request](%s)\n", url, err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Count for %s: 0 [error][resp read](%s)\n", url, err.Error())
		return
	}

	count := strings.Count(string(body), word)
	fmt.Printf("Count for %s: %d\n", url, count)
	s.addTotal(int64(count))
	return
}

//GetTotal for getting total number
func (s *HTTPSearcher) GetTotal() int64 {
	s.RLock()
	defer s.RUnlock()
	return s.total
}

//Wait for sync all goroutines
func (s *HTTPSearcher) Wait() {
	s.wg.Wait()
}

func (s *HTTPSearcher) addTotal(count int64) {
	s.RLock()
	s.total += count
	s.RUnlock()
}

func (s *HTTPSearcher) done() {
	<-s.limiter
	s.wg.Done()
}
