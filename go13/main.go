package main

import "fmt"

/// Anonim Fonksiyonlar
/*
Anonim fonksiyonların en büyük özelliği isimsiz olmalarıdır.
(Zaten adından da belli oluyor 🤔) Yazıldıkları yerde direkt olarak çalışırlar.
Çalışırken diğer fonksiyonlardaki gibi parametre verilemediği için
fonksiyonun sonuna parametre eklenerek çalışıtırılırlar. Örneğimizi görelim:
*/

func main() {
	value := "Salam Duyno"

	func(a string) {
		fmt.Println(a)
	}(value)
}
