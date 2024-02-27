package main

import "fmt"

func Danger() interface{} {

	return "Danger Bro"
}

func Warning() any {

	return "Warning Bro"
}

func main() {

	dangerTest := Danger()
	warningTest := Warning()

	fmt.Println(dangerTest)
	fmt.Println(warningTest)
}
