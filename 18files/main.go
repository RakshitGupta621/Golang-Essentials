package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "This needs to go in a file - Rakshit Gupta"

	file, err := os.Create("./myTestFile.text")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length,_ := io.WriteString(file, content)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myTestFile.text")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file in \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}