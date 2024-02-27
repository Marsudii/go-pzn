package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	test := NewMap("")

	//check jika data map nill
	if test == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(test)
	}

}
