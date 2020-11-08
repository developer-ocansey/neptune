package main

import (
	"log"
	"net/http"
	"os"
	"sync"
)

const scheme = "http://"

func main() {
	var wg sync.WaitGroup
	var mutx sync.Mutex

	if len(os.Args) < 2 {
		panic("please add url to the arguments when running this programs")
	}

	for _, url := range os.Args[1:] {
		go send(url, &wg, &mutx)
		wg.Add(1)
	}
	wg.Wait()
}

func send(url string, wg *sync.WaitGroup, mutx *sync.Mutex) {
	defer wg.Done()
	defer mutx.Unlock()

	res, err := http.NewRequest(http.MethodGet, scheme+url, nil)
	if err != nil {
		log.Fatal(err)
	}
	mutx.Lock()
	log.Print(res.Response.Status)
	return
}
