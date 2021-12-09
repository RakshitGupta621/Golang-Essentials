package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["CPP"] = "C++"
	languages["PY"] = "Python"
	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "PY")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in Golang
	// for key, value := range languages {
	// 	fmt.Printf("For key %v, value is %v\n", key, value)
	// }

	// when don't care about the key
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
