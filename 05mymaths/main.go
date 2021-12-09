package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in Golang!")

	// var mynumberOne int = 3
	// var mynumberTwo float64 = 4

	// fmt.Println("The sum is:", mynumberOne+int(mynumberTwo))

	// random number generator
	// rand.Seed(time.Now().UnixNano()) // rand on basis on random
	// fmt.Println(rand.Intn(5) + 1) // + 1 as we may get 0

	// random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
