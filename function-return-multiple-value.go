package main

import "fmt"

func sayHelloMultiple() (string, int) {
	return "Marsudi", 24
}

func checkErorr() (string, error) {
	return "Mengabaikan Value return", nil
}

func main() {
	name, age := sayHelloMultiple()
	fmt.Println("Nama saya = ", name, " Dan Umur saya =", age)

	abaikanReturnValue, _ := checkErorr()
	fmt.Println(abaikanReturnValue)
}
