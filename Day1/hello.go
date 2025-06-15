package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Welcome to Go,", name)

	users := []User{
		{"Alice", 17},
		{"Bob", 22},
	}

	for _, user := range users {
		if user.Age >= 18 {
			fmt.Println(user.Name, "is an Adult")
		} else {
			fmt.Println(user.Name, "is a Minor")
		}
	}
}
