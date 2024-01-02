package main

import "fmt"

func getHello() (firstName, middleName, lastName string) {
	firstName = "Udi"
	middleName = "Budi"
	lastName = "Bola"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getHello()
	fmt.Println(a, b, c)
}
