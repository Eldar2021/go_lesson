package main

import (
	"fmt"
	"os"
)

/// `os` ile Dosya Okuma ve Yazma
/*
* `os` paketi standart Golang paketleri içerisinde gelir ve
* dosya işlemleri yapabilmemiz için bize fonksiyonlar sağlar.
 */

func errorControl(err error) {
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}
}

func readFile(fileName string) {
	file, err := os.ReadFile(fileName)

	errorControl(err)

	fmt.Println(string(file))
}

func writeFile(fileName string, data []byte) {
	fmt.Println(data)

	// Dosya yazma işlemini başlatıyoruz
	// 0644 dosya yazdırma izni oluyor
	err := os.WriteFile(fileName, data, 0644)
	errorControl(err)
}

func main() {
	fileName := "example.txt"
	data := []byte("Almazbek")

	readFile(fileName)

	writeFile(fileName, data)

	readFile(fileName)

}
