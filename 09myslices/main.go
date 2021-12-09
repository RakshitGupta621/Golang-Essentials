package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on slices")
	var fruitList = []string{"Apple", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Apricot")
	fmt.Println(fruitList)

	fruitList = fruitList[1:] // slicing
	fmt.Println(fruitList)

	fruitList = fruitList[1:3] // (1, 3) open brackets
	fmt.Println(fruitList)

	highScores := make([]int, 4) 

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321)
	// We can append like this even beyond the size of highScores
	fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slices based on index

	var courses = []string{"reacjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	// from start to where index is (not including it) and from index+1 to the end 
	fmt.Println(courses)
} 
