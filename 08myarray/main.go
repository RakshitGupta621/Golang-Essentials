package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays on Golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Candy"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList)) // give initialised size

	var vegList = [3]string{"potato", "beans", "mushrooom"}
	fmt.Println("Vege list is: ", vegList)
	fmt.Println("Vege list is: ", len(vegList))
}


// ! We do not use arrays as much as we do in other programming languages
// Instead we use Slices