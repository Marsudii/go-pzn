package main

import "fmt"

//closure
//closure adalah membuat function tetapi mengakses variabel dari luar

func main() {

	//variabel counter = 0
	counter := 0

	// buat function increment di dalam variabel increment dengan mengakes counter
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	//panggil variabel increment sebanyak 3 kali
	increment()
	increment()
	increment()

	fmt.Println(counter)
}
