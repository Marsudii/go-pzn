package main

import "fmt"

func sayHelloToReturnValue(name string) string {

	return "hello " + name
}

func main() {

	result := sayHelloToReturnValue("Marsudi")

	fmt.Println(result)
}
