package main

import "fmt"

type Alamats struct {
	City, Province, Country string
}

// calling function pointer *Alamats
func ChangeAlamatToIndo(alamat *Alamats) {
	alamat.Country = "Indonesia"
}

func main() {
	alamat := Alamats{"subang", "jawa barat", ""}
	ChangeAlamatToIndo(&alamat) // & refrence ke alamat
	fmt.Println(alamat)
}
