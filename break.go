package main

import "fmt"

func main() {

	//BREAK
	//loping 0 sampai 10 jika sudah 5 berhenti ( break )
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan", i)
	}

}
