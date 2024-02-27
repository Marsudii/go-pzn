package main

import "fmt"

// INTERFACE
type CrudName interface {
	InsertName() string
	DeleteName() string
}

// FUNGSI/METHOD INTERFACE CREATE
func CreateName(value CrudName) {
	fmt.Println("Berhasil Insert Data", value.InsertName())
}

// FUNGSI/METHOD INTERFACE DELETE
func DeleteName(value CrudName) {
	fmt.Println("Berhasil Hapus Data", value.DeleteName())
}

// STRUCT DATA NAME
type Person struct {
	Name string
}

// IMPLEMENTASI INSERT INTERFACE
func (person Person) InsertName() string {
	return person.Name
}

// IMPLEMENTASI DELETE INTERFACE
func (person Person) DeleteName() string {
	return person.Name
}

func main() {
	personInsert := Person{Name: "UDI Insert"}
	CreateName(personInsert)

	personDelete := Person{Name: "UDI Delete"}
	DeleteName(personDelete)
}
