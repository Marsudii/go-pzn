package main

import "fmt"

//menggunakan type declaration

type Filter func(string) string

// FUNCTION AS PARAMETER mengunnakan type declaration
// BUAT FUNGSI CHEKCK KATA dengan parameter ("string kata", function filter), result "string")
func checkWords(words string, filter Filter) {
	fmt.Println("hello ", filter(words))
}

// FUNCTION AS PARAMETER tidak menggunakan type declaration
// BUAT FUNGSI CHEKCK KATA dengan parameter ("string kata", function filter), result "string")
//func checkWords(words string, filter func(words string) string) {
//	fmt.Println("hello ", filter(words))
//}

// function filter words
func filterSpam(name string) string {
	if name == "anjing" {
		return "Words Is Sensored"
	} else {
		return name
	}
}

func main() {
	//call function check kata dengan paramter katanya = string dan function filternya
	checkWords("Asik", filterSpam)

}
