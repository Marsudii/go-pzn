package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"ID Required"}
	}
	if id != "UDI" {
		return &NotFoundError{"Data Tidak Ditemukan"}
	}
	return nil
}

func main() {

	err := SaveData("UDI", nil)
	if err != nil {
		if errValidationError, ok := err.(*ValidationError); ok {
			fmt.Println(errValidationError)
		} else if errNotFoundError, ok := err.(*NotFoundError); ok {
			fmt.Println(errNotFoundError)
		} else {
			fmt.Println("unknow erorr", err.Error())
		}
	} else {
		fmt.Println("Suksess Data Tersimpan")
	}
}
