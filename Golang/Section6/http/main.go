package main

import (
	"io"
	"os"
	"fmt"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// METHOD 1 to get body
	// bs := make([]byte, 99999)	
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// METHOD 2
	// os.Stdout implements Writer Interface
	// resp.Body implements Reader Interface
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// logWriter is now implementing the Writer interface
// because logWriter has a function called Write with []byte argument and 
// (int, error) return type
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))

	return len(bs), nil
}