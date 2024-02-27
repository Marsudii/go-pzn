package main

import "fmt"

type Alamat struct {
	AddressOne   string
	AddressTwo   string
	AddressThree string
}

func main() {
	// PASSING VARIABLE BY VALUE POINTER
	alamat1 := Alamat{"Jakarta", "Depok", "Jawa Barat"}
	//DUPLIKATE DATA DARI ALAMAT1
	alamat2 := alamat1 //copy value dari alamat1
	//UBAH VALUE ALAMAT 2 dengan PONITER ADDEREONE menjadi CINERE
	alamat2.AddressOne = "Cinere"
	fmt.Println(alamat1)
	fmt.Println(alamat2)

	//PASSING BY REFERENCE POINTER mengunakan & untuk mereference
	alamatJKT := Alamat{"Jakarta Belum Diubah", "Jakarta Timur", "Jakarta Selatan"}
	alamatHome := &alamatJKT
	fmt.Println(alamatHome)
	alamatHome.AddressOne = "Jakarta Pusat"
	fmt.Println(alamatHome)

}
