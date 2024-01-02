package main

import "fmt"

// variable argument (...)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	resultSum := sumAll(10, 10, 10, 10, 10)
	fmt.Println(resultSum)

	//misal kita punya slice
	angka := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	resultAngka := sumAll(angka...)
	fmt.Println(resultAngka)

}
