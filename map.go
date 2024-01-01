package main

import "fmt"

func main() {

	//var person = map[string]string{}
	//person["name"] = "udi"
	//person["age"] = "19"

	person := map[string]string{
		"name": "Udi",
		"age":  "24",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])

	//fmt.Println(person["salah"]) //string kosong

	//funtion map

	panjangMap := len(person)
	getMap := person["name"]
	person["name"] = "Marsudi"

	//menghapus map dengan key name
	delete(person, "name")

	fmt.Println(panjangMap) //panjang map
	fmt.Println(getMap)     // get data key map
	fmt.Println(person)     //ubah data map dengab key name = marsudi

}
