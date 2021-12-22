package main

import (
	"fmt"
)

func isEven(number int) string {
	if number % 2 == 0 {
		return "even"
	}

	return "odd"
}

func main() {
	numbers := []int{}
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}
	

	for _, number := range numbers {
		fmt.Printf("%v is %v \n", number, isEven(number))		
	}
}