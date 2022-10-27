package main

import (
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.qq.com/",
}

func main() {
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil {
				println(resp.Status)
			} else {
				println("eeeee...")
			}
		}(url)
	}
	wg.Wait()
}
