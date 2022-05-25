package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sites := []string{
		"http://cisco.com",
		"http://golang.org",
		"http://nyt.com",
		"http://wsj.com",
	}

	// using locks instead of channels
	for _, site := range sites {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			checkStatus(s)
		}(site)
	}
	wg.Wait()
	fmt.Println("Done with all go routines.")

	// channels stuff
	// c := make(chan string)
	// for site := range c {
	//	go func(s string) {
	//		time.Sleep(time.Second)
	//		checkStatus(s, c)
	//	}(site)
	// }
}

func checkStatus(site string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "GET request has failed.")
		return
	}
	fmt.Println(site, "GET request has succeeded.")
}
