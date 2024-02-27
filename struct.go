package main

import "fmt"

type User struct {
	Name, Address string
	Age           int
	Status        bool
}

func main() {

	var budi User
	budi.Name = "Budi"
	budi.Address = "Bandung"
	budi.Age = 10
	budi.Status = false

	//get all struct data
	fmt.Println(budi)
	//get spesific struct data
	fmt.Println(budi.Name)
	fmt.Println(budi.Age)

	var marsudi = User{
		Name:    "Marsudi",
		Address: "Jakarta",
		Age:     20,
		Status:  true,
	}
	//get all struct data
	fmt.Println(marsudi)
	//get spesific struct data
	fmt.Println(marsudi.Name)
	fmt.Println(marsudi.Age)

	var eko = User{"Eko", "Depok", 10, true}
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Age)
}
