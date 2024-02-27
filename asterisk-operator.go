package main

import "fmt"

type Human struct {
	Name string
	Age  int
	City string
}

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	// Create a Human struct instance
	humanSatu := Human{"Marsudi", 25, "Depok"}
	// Create a pointer to the humanSatu instance
	humanDua := &humanSatu
	// Change the name of the humanDua instance
	humanDua.Name = "Budi"
	// Assign a new Human struct instance to the memory location pointed by humanDua
	humanDua = &Human{"Tono", 25, "Jakarta"}
	// Print the humanSatu instance
	fmt.Println(humanSatu)

	fmt.Println(humanDua)

	//MENGUBAH VALUE SEMUANYA PADA ADDRESS1 &
	address1 := Address{"Jakarta Timur", "DKI Jakarta", "Indonesia"}
	address2 := &address1
	address2.City = "Jakarta Selatan"
	address2 = &Address{"Jakarta Barat", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) //TIDAK BERUBAH
	fmt.Println(address2) //BERUBAH MENJADI JAKARTA BARAT

	address3 := Address{"Depok", "Jawa Barat", "IND"}
	address4 := &address3
	*address4 = Address{"Bogor", "Jabar", "IND"}

	fmt.Println(address3) //BERUBAH JADI BOGOR
	fmt.Println(address4) //BERUBAH JADI BOGOR
}
