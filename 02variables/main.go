package main

import "fmt"

const LoginToken string = "raksh" // Public

func main() {
	var username string = "Rakshit"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of t ype: %T \n", smallVal)

	var smallFloat float32 = 255.4585858585858499494
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	
	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Another variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "Blueprint.com" // automatically detects the type
	fmt.Println(website) 

	// no var declarations
	numberOfUser := 30000 // walrus operator: only allowed inside functions
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}