package main

import (
	"belajar-golang-pzn/database"
	//blank identifire jika tidak digunakan kasih _
	_ "belajar-golang-pzn/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
