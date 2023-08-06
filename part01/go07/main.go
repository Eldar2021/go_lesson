package main

import "fmt"

/// Değişkenler ve Atanması

/*
Değişkenler içerisinde değer barındırarak RAM’e kaydettiğimiz bilgilerdir.
Değişkenler programımızın işleyişinde önemli bir role sahiptir.
Değişkenleri şu şekillerde atayabiliriz. Değişkenler var ifadesi ile atanır.
Tabi ki zorunlu değildir.
*/

func main() {
	var name01 string = "Umar"
	var age01 int = 40
	var isStudent01 bool = true

	fmt.Println(name01)
	fmt.Println(age01)
	fmt.Println(isStudent01)

	// Değişken atamasında illaki değişkenin veri tipini belirtmemiz gerekmez.
	// Yazdığımız değere göre Golang otomatik olarak veri tipini algılar.

	var name02 = "Usman"
	var age02 = 30
	var isStudent02 = true

	fmt.Println(name02)
	fmt.Println(age02)
	fmt.Println(isStudent02)

	// En basit şekilde değişken ataması yapmak istersek;
	// Başına var eklemeden de değişken atamak mümkündür. Bu şekilde yapmak için := işaretlerini kullanırız. 
	// Aynı şekilde bu yöntemde de verinin tipi otomatik algılanır.

	name03 := "Ali"
	age03 := 20
	isStudent03 := true

	fmt.Println(name03)
	fmt.Println(age03)
	fmt.Println(isStudent03)

}
