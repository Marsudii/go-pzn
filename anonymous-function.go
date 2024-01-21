package main

import "fmt"

// Anonymous Function
type BlockedName func(string) bool

func registerUser(name string, blockedName BlockedName) {
	if blockedName(name) {
		fmt.Println("Name for You Blocked !!", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {

	//Anonymous Function di jadikan variabel
	blockedName := func(name string) bool {
		return name == "anjing"
	}
	//panggil variabel Anonymous Function
	registerUser("anjing", blockedName)

	// bisa juga di tulis langsung di function register user
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

}
