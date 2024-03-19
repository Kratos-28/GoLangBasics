package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)

	}

	for l := range c {

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Link down", err)
		c <- link
		return
	}

	fmt.Println("Link is up", link)
	c <- link
}
