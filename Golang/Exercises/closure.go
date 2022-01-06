package main

import "fmt"

func intSeq() func() int {
    i := 0 // this assignment happens only when you use intSeq()
	fmt.Println(" >> i ", i)
    return func() int {
        i++
        return i
    }
}

func main() {

	// we call intSeq assigning the result (a function) to nextInt
	// the function value captures its own "i" value, which will be updated
	// every time we call nextInt
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
    fmt.Println(newInts())
}