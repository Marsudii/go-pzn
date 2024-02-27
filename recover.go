package main

import "fmt"

//recover akan menghentikan aplikasi sedangkan defer akan diekseskusi terakhir

func endApps() {
	fmt.Println("End App")
	messages := recover()
	fmt.Println("Terjadi Erorr =", messages)
}

func run(error bool) {
	defer endApps()
	if error {
		panic("panic error")
	}
}

func main() {
	run(true)
	//aplikasi penanganan erorr sama dengan try catch
	fmt.Println("Marsudi App")
}
