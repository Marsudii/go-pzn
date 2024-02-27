package main

import "fmt"

//panic akan menghentikan aplikasi sedangkan defer akan diekseskusi terakhir

func endApp() {
	fmt.Println("End App")
}

func runApps(error bool) {
	defer endApp()
	if error {
		panic("Upss erorr")
	}
}

func main() {
	runApps(true)
}
