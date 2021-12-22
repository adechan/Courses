package main

import (
	"io"
	"os"
	"fmt"
)

func readFromFile(fileName string) {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(string(bs))
}

func readFromFile2(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}

func main() {
	args := os.Args
	fileName := args[1]

	readFromFile(fileName)
	readFromFile2(fileName)
}
