package main

import "fmt"

func main() {

	// TIPE DATA KONSTAN tidak digunakan tidak masalah tetapi jika diubah tidak bisa
	const angkaOne int = 1
	const angkaTwo = 2
	fmt.Println(angkaOne, angkaTwo)

	// TIPE DATA KONSTAT BANYAK
	const (
		angkaThree = 3
		angkaFour  = 4
	)
	fmt.Println(angkaFour)
}
