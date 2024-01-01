package main

import "fmt"

func main() {

	name := "udi"

	switch name {
	case "udi":
		fmt.Println("Hello Udi")
	case "eko":
		fmt.Println("hello Eko")
	default:
		fmt.Println("Hello Boleh kenalan ?")
	}

	switch lengthName := len(name); lengthName < 5 {
	case true:
		fmt.Println("panjang nama minimal 5")
	default:
		fmt.Println("nama Sudah Benar")
	}

}
