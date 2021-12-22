package main

import (
	"fmt"
	"net/http"
)

// not a good solution because after each request you need to wait for a response 
// so for example, the facebook request will be made ONLY after the google one received
// an answer
func main1() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazing.com",
	}

	for _, link := range links {
		verifyIfDomainIsDown1(link)
	}
}

func verifyIfDomainIsDown1(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("%s is down. \n", link)
		return
	}

	fmt.Printf("%s is okay. \n", link)
}