package main

import (
	"belajar-golang-pzn/helper" //IMPORT HELPER METHOD HELLO
	"fmt"
)

func main() {

	callHelper := helper.Hello("budi")
	fmt.Println(callHelper)

	fmt.Println(helper.Apk) //bisa diakses
	//fmt.Println(helper.version)//tidak bisa
}
