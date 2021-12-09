package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	// greeter()
	// greeterTwo()
	result := adder(3, 5)
	fmt.Println("Result is: ", result)
	// proRes := proAdder(2, 5, 8, 7)
	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)
}

// at the end of func we have to define the return type also
func adder(valOne int, vaTwo int) int {
	return valOne + vaTwo
}

// ...int means all values are of type integers
// func proAdder(values ...int) int {
// 	total := 0
// 	for _, val := range values {
// 		total += val
// 	}
// 	return total
// }

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}

// func greeterTwo() {
// 	fmt.Println("Another method")
// }

// func greeter() {
// 	fmt.Println("Namastey from golang")
// }
