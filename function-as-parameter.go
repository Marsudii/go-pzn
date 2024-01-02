package main

import "fmt"

func checkWords(words string, filter func(words string) string) {
	fmt.Println("hello ", filter(words))
}

func filterSpam(name string) string {
	if name == "anjing" {
		return "****"
	} else {
		return name
	}
}

func main() {
	fmt.Println(checkWords("wkwkw", filterSpam))
}
