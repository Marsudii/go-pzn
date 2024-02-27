package helper

// TIDAK BISA DIAKSES DI BEDA PACKEGS KARENA HURUF KECIL
var version = "10000"

// BISA DIAKSES DILUAR PACKEGES
var Apk = "v1.0.0"

func Hello(name string) string {
	return "hello " + name
}

func helloWord(name string) string {
	return "hello " + name
}
