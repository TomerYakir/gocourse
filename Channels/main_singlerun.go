package main

import (
	"fmt"
	"net/http"
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

	// now wait for data to be received on the channel - this is blocking
	for range urlsToCheck { // can be for i := 0; i < len(urlsToCheck); i++
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	// fmt.Println("checking ", link)
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, "might be down: ", err)
		c <- link + " might be down " + err.Error()
		return
	}
	c <- link + " is up!"
}
