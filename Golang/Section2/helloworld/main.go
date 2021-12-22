package main

import "fmt";

func main() {
	fmt.Println("Hello World")
}

// 1. How do we run the code in our project? 
// -> go run main.go : compiles all the files and then instantly executes them
// -> go build main.go : 
	// - just compiles a program  (gies us "main.exe", an executable file build on our code)
	// - and then you can execute the file by typing main.exe

// 2. What does 'package main' mean?
// package == project == workspace
// - a package can have many .go files in it, and each of these files must have "package main"
// to declare the package they are from
// --> Why is it called main? 
// <-- 
// 2 types of packages in Go: executable + reusable
// executable: generates a file that we can run
// reusable: code used as helpers 
// main is used ONLY for executable files and it must have a function called main in it

// 3. What does "import "fmt" " mean?
// - import is used to give our package access to code written in another package
// "fmt" - a package - Standard lib

// 4. What is that "func" thing?
// just a normal function declaration

// 5. How is the main.go file organized? 
// 