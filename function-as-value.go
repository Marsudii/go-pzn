package main

import "fmt"

func functionValue() string {
	return "Hello World"
}
func main() {

	//function di simpan kedalam variabel getFunctionValue
	getFunctionValue := functionValue
	fmt.Println(getFunctionValue())
}
