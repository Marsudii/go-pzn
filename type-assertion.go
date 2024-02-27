package main

import "fmt"

func random() interface{} {
	return 1000
}

func main() {

	var data any = random()
	//resultString := data.(string)
	//fmt.Println(resultString)

	//CODE PANIC karena beda tipe data variabel
	//resultInt := data.(int)
	//fmt.Println(resultInt)

	//solusinya menggunakan swicth dan check tipe data
	switch data.(type) {
	case string:
		fmt.Println("Type data string", data)
	case int:
		fmt.Println("Type data Integer", data)
	case bool:
		fmt.Println("Type data Boolean", data)
	default:
		fmt.Println("Unknow Type Data", data)
	}

}
