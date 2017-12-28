package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urlsToCheck := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://mysp1122.com",
	}

	c := make(chan string)

	for _, url := range urlsToCheck {
		go checkLink(url, c)
	}

	//
	//for { // infinite loop
	for l := range c { // wait for data to be received on the channel
		go func(link string) {
			time.Sleep(1 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	// fmt.Println("checking ", link)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down: ", err)
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
