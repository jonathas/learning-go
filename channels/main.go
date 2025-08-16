package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://facebook.com",
		"https://google.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Watch the channel, then every time a value is received, spawn a new goroutine
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	res, err := http.Get(link)
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Println(link, "might be down:", err)
		c <- link
		return
	}

	fmt.Println(link, "is up:", res.Status)
	c <- link
}