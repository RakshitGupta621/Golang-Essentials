package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	rakshit := User{"Rakshit", "rakshit_2019bite041@nitsri.net", true, 20}
	fmt.Println(rakshit)
	fmt.Printf("Details are %+v\n", rakshit)
	fmt.Printf("Name is %v and email is %v.", rakshit.Name, rakshit.Email)
	rakshit.GetStatus()
	rakshit.NewMail()
	fmt.Printf("Name is %v and email is %v.", rakshit.Name, rakshit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("\nIs user status: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user: ", u.Email)
}

// !There are no classes so we use methods by using functions in structures
