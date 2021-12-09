package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	rakshit := User{"Rakshit", "rakshit_2019bite041@nitsri.net", true, 20}
	fmt.Println(rakshit)
	fmt.Printf("Details are %+v\n", rakshit)
	fmt.Printf("Name is %v and email is %v.",rakshit.Name,rakshit.Email )
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// there are no classes in golang
// there is no inheritance in golang
