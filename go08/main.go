package main

import "fmt"

/// Sabitler
/*
Sabitler de değişkenler gibi değer alır.
Fakat adından da anlaşılabileceği üzere verilen değer daha sonradan değiştirilemez.

Sabitler tanımlanırken başına const eklenir. Örnek olarak;
*/

func main() {
	const isim string = "Ali"
	// const isim = "Veli" {give error}
	fmt.Println(isim)

	// const ile := beraber kullanılamaz.

	/*
		Yanlış kullanım: const isim := "Ali"
		Doğru kullanım: const isim = "Ali"
	*/

	const name string = "Ali"
	// name = "Ali" {give error}
	const yas = 20

	fmt.Println(name)
	fmt.Println(yas)
}
