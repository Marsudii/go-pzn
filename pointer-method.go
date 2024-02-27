package main

import "fmt"

type Man struct {
	Name string
}
type Women struct {
	Name string
}

// tanpa pointer method
func (man Man) MariedMen() {
	man.Name = "Mr" + man.Name
}

// dengan method pointer *
func (women *Women) MariedWomen() {
	women.Name = "Mrs " + women.Name
}

func main() {
	//tanpa pointer method
	udi := Man{"Udi"}
	udi.MariedMen()
	fmt.Println(udi.Name) //tidak berubah karena data di duplikate saat dikirim di method

	//menggunakan pointer method
	zirly := Women{"Zirly"}
	zirly.MariedWomen()
	fmt.Println(zirly.Name) //berubbah menjadi Mrs Zirly

}
