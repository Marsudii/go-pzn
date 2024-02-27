package main

import "fmt"

type Addresss struct {
	City, Province, Country string
}

func main() {

	//TANPA NEW
	var alamat1 *Addresss = &Addresss{} //tanpa NEW
	var alamat2 *Addresss = alamat1
	alamat2.Country = "Indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)

	//DENGAN NEW OPERATOR
	var alamat3 *Addresss = new(Addresss)
	var alamat4 *Addresss = alamat3
	alamat4.Country = "Malaysia"
	fmt.Println(alamat3)
	fmt.Println(alamat4)
}
