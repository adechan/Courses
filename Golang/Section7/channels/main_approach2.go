package main

import (
	"time"
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazing.com",
	}

	c := make(chan string) // communicate through the channel with values of string
	for _, link := range links {
		go verifyIfDomainIsDown(link, c) // this function needs to be executed inside a new Go Routine
	}

	// handles when a request gets completed
	// for {  // INFINITE LOOP
	// 	// fmt.Println(<- c) // it waits for a message to come before going to the next one
	// 	// whenever we receive a value, we should verifyIfDomainIsDown again

	// 	go verifyIfDomainIsDown(<- c, c)
	// }

 
	// watch the channel c
	// and whenever a value comes out of it, assign that value to l
	for l := range c {  
		go func(l string) {
			time.Sleep(time.Second)
			verifyIfDomainIsDown(l, c)
		}(l)
	}
}

func verifyIfDomainIsDown(link string, c chan string) {
	_, err := http.Get(link) // the Go Routine will see that this is some blocking code and will start a new Go Routine for the next for loop

	if err != nil {
		fmt.Printf("%s is down. \n", link)
		c <- link
		return
	}

	fmt.Printf("%s is okay. \n", link)
	c <- link
}