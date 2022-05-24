package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://cisco.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, site := range sites {
		go checkStatus(site, c)
	}
	for site := range c {
		go func(s string) {
			time.Sleep(time.Second)
			checkStatus(s, c)
		}(site)
	}
}

func checkStatus(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "GET request has failed.")
		c <- site
		return
	}
	fmt.Println(site, "GET request has succeeded.")
	c <- site
}
