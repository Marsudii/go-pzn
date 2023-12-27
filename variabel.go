package main

import "fmt"

func variabel() {

	//TIPE DATA DI DEKLARASIKAN VAR
	var name string
	name = "MArsudi"
	fmt.Println(name)
	name = "Marsudi"
	fmt.Println(name)

	//TIPE DATA LANGSUNG VAR
	var angka = 1
	fmt.Println(angka)
	angka = 3
	fmt.Println(angka)

	//TIPE DATA VARIABEL dengan := (NOTE:tidak dapat dibuat ulang)
	fullName := "HELLO"
	fmt.Println(fullName)

	// TIPE DATA BANYAK VAR
	var (
		fullname   = "MARSUDI"
		middlename = "UDI"
	)
	fmt.Println(fullname, middlename)

}
