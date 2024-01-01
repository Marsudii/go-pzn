package main

import "fmt"

func main() {
	nameOne := "Marsudid"
	//IF EXPRESIION
	if nameOne == "Marsudi" {
		fmt.Println("nama diterima")
	}

	//IF ELSE EXPRESSION
	nameTwo := "Marsudid"
	if nameTwo == "Marsudi" {
		fmt.Println("nama diterima")
	} else {
		fmt.Println("nama ditolak")
	}

	//IF, ELSE IF Expresiion
	nameThree := "Yolo"
	if nameThree == "yolo" {
		fmt.Println("nama yolo")
	} else if nameThree == "Yolo" {
		fmt.Println("Yolo")
	} else {
		fmt.Println("nothing")
	}

	//IF SHORT STATEMENT

	nameFour := "Marsudi" //nama diterima karena panjang lebih dari 5
	// nameFour := "udi" //nama ditolak karena panjang kurang dari 5
	if lenghtName := len(nameFour); lenghtName < 5 {
		fmt.Println("nama minimal 5", nameFour)
	} else {
		fmt.Println("nama diterima")
	}
}
