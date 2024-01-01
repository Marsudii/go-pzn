package main

import "fmt"

func main() {

	//FOR ADA INIT STATEMENT dan POST STATEMENT

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan", counter)
	}
	fmt.Println("selesai")

	//slice for loop use len
	nameSlice := []string{"marsudi", "udi", "eko"}
	for index := 0; index < len(nameSlice); index++ {
		fmt.Println("Perulangan SLice = ", nameSlice[index])
	}

	//array for loop use len
	nameArray := [...]string{"MARSUDI", "UDI", "EKO"}
	for i := 0; i < len(nameArray); i++ {
		fmt.Println("Perulangan Array", nameArray[i])
	}

	//map for loop use len
	nameMap := map[string]string{"NAMA": "MARSUDI", "ADDRES": "DEPOK"}
	for i := 0; i < len(nameMap); i++ {
		fmt.Println("Perulangan Map = ", nameMap["NAMA"])
	}

	//slice FOR Range
	numberslice := []int{100, 200, 300}
	for index, number := range numberslice {
		fmt.Println("index", index, "=", number)
	}

	//array For range
	numbersArray := [...]int{10, 20, 30}
	for index, number := range numbersArray {
		fmt.Println("Index", index, '=', number)
	}

	//map for range
	numbersMap := map[int]int{0: 1000, 1: 2000, 2: 3000}
	for i, number := range numbersMap {
		fmt.Println("Key", i, "=", number)
	}

	//misal tidak butuh index kasih _ underscore
	numbersMaps := map[int]int{0: 1000, 1: 2000, 2: 3000}
	for _, number := range numbersMaps {
		fmt.Println(number)
	}

}
