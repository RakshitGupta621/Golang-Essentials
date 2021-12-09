package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for pizza: ")
	// Regular syntax is input, err
	input, _ := reader.ReadString('\n') // Error is catch and input is try
	// now we need to understand how to convert this string to number

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is rating %T ", input)
}
