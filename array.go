package main

import "fmt"

func main() {

	var product [3]string
	product[0] = "Buku"
	product[1] = "Pulpen"
	product[2] = "Pensil"

	//AMBIL SEMUA DATA
	fmt.Println(product)
	//AMBIL SATU PER SATU BERDASARKAN INDEX
	fmt.Println(product[0])
	fmt.Println(product[1])
	fmt.Println(product[2])

	//buat langusung array beserta isinya
	var angka = [10]int{
		0,
		1,
	}

	fmt.Println(angka)

	var number = [...]int{
		90,
		80,
	}
	fmt.Println(number)

	//FUNCTION ARRAY
	//jumlah panjang array
	fmt.Println(len(number))
	//mendapat data di posisi index
	fmt.Println(number[1])
	//mengubah data di posisi index
	number[1] = 900
	fmt.Println(number[1])

}
