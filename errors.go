package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilaiSatu int, nilaiDua int) (int, error) {
	if nilaiDua == 0 {
		return 0, errors.New("Pembagian dengan 0")
	} else {
		return nilaiSatu / nilaiDua, nil
	}
}
func main() {
	result, err := Pembagian(100, 2)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())

	}
}
