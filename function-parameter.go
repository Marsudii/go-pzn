package main

import "fmt"

// paramemter boleh 1 bahkan lebih tetapi kita wajib mengirim paramater
func sayHelloTo(username string) {
	fmt.Println("Hello", username)
}

func main() {
	//memanggil function dan memberikan paramater
	sayHelloTo("Marsudi")
}
