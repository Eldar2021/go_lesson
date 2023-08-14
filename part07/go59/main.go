package main

import (
	"fmt"
	"runtime"
)

/// İşletim Sistemini Görme
/*
* Go programının çalıştığı işletim sistemi görmek için
* aşağıdaki kodları yazabilirsiniz.
 */

func main() {
	operationSystem := runtime.GOOS

	fmt.Println(operationSystem)

	if operationSystem == "windows" {
		fmt.Println("Windows için yönetici olarak çalıştırın.")
	} else if operationSystem == "linux" {
		fmt.Println("Linux için sudo komutu ile çalıştırın.")
	} else if operationSystem == "darwin" {
		fmt.Println("Mac için brew komutu ile çalıştırın.")
	} else {
		fmt.Println("Geçersiz işletim sistemi!")
	}
}
