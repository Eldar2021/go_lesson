package main

import (
	"fmt"
	"os"
)

/// Dosya Varlığı Kontrolü
/*
* Go programımızda kullancağımız bir bir dosyanın varlığını
* `os` paketi ile kontrol edebiliriz.
 */

func fileExist(fileName string) bool {
	value, err := os.Stat(fileName)
	fmt.Println("Bir value :) ", value)
	if err != nil {
		fmt.Println("Bir hata oluştu :) ", err)
		return false
	}

	return !value.IsDir()
}

func main() {
	filename := "main.go" // "example_file.txt"
	if fileExist(filename) {
		fmt.Println(filename, "bulunuyor")
	} else {
		fmt.Println(filename, "bulunmuyor!")
	}
}
