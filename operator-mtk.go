package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = 20
	var d = 5

	var resultPertambahan = a + b + c + d
	fmt.Println("Penjumlahan=", resultPertambahan)

	var resultPengurangan = a - c
	fmt.Println("Pengurangan=", resultPengurangan)

	var resultPerkalian = a * b
	fmt.Println("Perkalian=", resultPerkalian)

	var resultPembagian = a / b
	fmt.Println("Pembagian=", resultPembagian)

	//AUGMENT ASSIGMENT
	var i = 10
	i += 1 //i = i + 1
	fmt.Println("AUGMENT ASSIGMENT=", i)

	i += 9 //i = i + 9
	fmt.Println("AUGMENT ASSIGMENT=", i)

	// UNARY OPERATOR
	var j = 1
	j++ // j = 1 + 1
	fmt.Println("UNARY OPERATOR", j)
	j++ // j = 2 + 1
	fmt.Println("UNARY OPERATOR", j)
	j-- // j = 2 - 1
	fmt.Println("UNARY OPERATOR", j)

}
