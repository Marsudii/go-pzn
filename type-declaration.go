package main

import "fmt"

func typeDeclaration() {

	type noKTP string
	var ktpUdi noKTP = "1212121"
	fmt.Println(ktpUdi)

	var contoh = "909090"
	var contohKtp noKTP = noKTP(contoh)
	fmt.Println(contohKtp)

}
