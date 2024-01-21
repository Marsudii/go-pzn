package main

import "fmt"

//recursive-function

// program factorial menggunakan loop tanpa recursive-function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//program factorial menggunakan recursive-function

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		//memanggil fungsi dirinya sendiri
		return value * factorialRecursive(value-1)
	}
}

func main() {
	//cetak factorial loop
	fmt.Println(factorialLoop(10))
	//cetak recursive function
	fmt.Println(factorialRecursive(10))
}
