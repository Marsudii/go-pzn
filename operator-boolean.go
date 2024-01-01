package main

import "fmt"

func main() {

	// AND OPERATOR
	var nilaiUas = 90
	var nilaiUts = 90

	var nilaiLulusUas bool = nilaiUas > 80
	var nilaiLulusUts bool = nilaiUts > 80
	var lulus bool = nilaiLulusUts && nilaiLulusUas
	fmt.Println("Status Lulus = ", lulus)

}
