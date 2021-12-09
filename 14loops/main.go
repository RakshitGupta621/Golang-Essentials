package main

import "fmt"

func main() {
	// fmt.Println("Welcome to loops in Golang")
	// days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	// fmt.Println(days)

// 1.
	// for i:=0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }

// 2.
	// for i := range days{
	// 	fmt.Println(days[i])
	// }

// 3.
	// for index, day := range days{
	// 	fmt.Printf("index is %v and day is %v\n", index, day)
	// }

// 4.
	// for _, day := range days{
	// 	fmt.Printf("Day is %v\n", day)
	// }

// 5.
	// while loop, break, continue
	arbitrary := 1
 
	for arbitrary < 10 {
		if arbitrary == 2 {
			goto labelName
		}

		if arbitrary == 5 {
			arbitrary++
			break // can also use continue here to show its implementation
		}
		fmt.Println("Value is: ", arbitrary)
		arbitrary++
	}

// 6. go to 
	labelName: 
		fmt.Println("Good afternoon")

}
