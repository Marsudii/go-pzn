package main

import "fmt"

//defer adalah membaca log ketika erorr walaupun ada erorr di tengah atau di atas function

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Run App")
}

func main() {
	runApp()
}
