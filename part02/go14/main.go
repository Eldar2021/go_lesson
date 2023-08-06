package main

import "fmt"

/// Boş Tanımlayıcılar
/*
Golang kodlarımızda bazen 2 adet değer döndüren fonksiyonlar kullanırız.
Bu değerlerden hangisini kullanmak istemiyorsak, değişken adı
yerine _ (alt tire) kullanırız.
*/

func simpleFunc01(input int) (int, int) {
	a := input * 2
	b := input * 4
	return a, b
}

func main() {
	a, b := simpleFunc01(3)
	fmt.Println(a, b) // 6, 12

	c, _ := simpleFunc01(3) // _ koldonulbagandygybyzdy bildirebiz
	fmt.Println(c)          // 6
}
