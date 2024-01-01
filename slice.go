package main

import "fmt"

func main() {

	iniArray := [...]string{"Hello", "Array"}
	iniSLice := []string{"Hello", "Slice"}

	fmt.Println(iniArray)
	fmt.Println(iniSLice)

	//tipe data slice 3 data POINTER, LENGHT, CAPACITY

	hari := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	slice1 := hari[4:6]
	fmt.Println(slice1)

	slice2 := hari[2:] //index 2 sampai terakhir
	fmt.Println(slice2)

	slice3 := hari[:5]
	fmt.Println(slice3) //index pertama sampai ke 5

	slice4 := hari[:]
	fmt.Println(slice4) //ambil semuanya

	//funtion slice
	//panjang slice
	fmt.Println(len(slice4))
	//capasitas slice
	fmt.Println(cap(slice4))
	//append slice (menambahkan data di posisi terkahir)
	fmt.Println(append(slice4, "Libur"))
	fmt.Println(append(slice4[7:], "ga libur"))

	var newSlice []string = make([]string, 2, 5)

	newSlice[0] = "marsudi"
	newSlice[1] = "marsudi"
	//newSlice[2] = "eko" //harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) //2
	fmt.Println(cap(newSlice)) // 5

	newSlice2 := append(newSlice, "EKO")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	//copy data slice
	fromSLice := hari[:]
	toSlice := make([]string, len(fromSLice), cap(fromSLice))
	copy(toSlice, fromSLice)
	fmt.Println(fromSLice)
	fmt.Println(toSlice)

}
