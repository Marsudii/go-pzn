package main

import "fmt"

// OPERATOR PERBANDINGAN
func main() {

	var name1 = "Udi"
	var name2 = "Udi"
	var tidakSamaDengan bool = name1 != name2
	var samaDengan bool = name1 == name2

	fmt.Println("Tidak Sama Dengan =", tidakSamaDengan)
	fmt.Println("Sama Dengan =", samaDengan)

	var huruf1 = "a"
	var huruf2 = "b"
	var lebihDari bool = huruf1 > huruf2
	var kurangDari bool = huruf1 < huruf2

	fmt.Println("Lebih Dari =", lebihDari)
	fmt.Println("Kurang Dari =", kurangDari)

}
