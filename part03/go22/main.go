package main

import "fmt"

/// Struct Fonksiyonlar (Methodlar)
/*
Bu bölümde bir struct'a özel nasıl fonksiyon oluşturacağımızı göreceğiz.
*/

type person struct {
	name, lastname string
}

func (p person) sayHello() {
	fmt.Printf("Hi! My name is %s, My lastname is %s ", p.name, p.lastname)
}

func main() {
	p1 := person{name: "Eldi", lastname: "Almazbek"}
	p1.sayHello()
}

/*
`person` isminde bir struct tipi oluşturduk. Bu yapımızın tıpkı
insanlarda olduğu gibi isim ve soyisim değişkenleri var.
Hemen aşağısında bir fonksiyon oluşturduk. Bu fonksiyonumuzun özelliği
ise fonksiyonun isminden önce parantez içerisinde hangi struct'ta
çalışacağını belirtmemizdir. `person` struct'ının içerindeki değişkenlere
ise `p` değişkeni ile eriştik.

Daha sonra main fonksiyonumuzda `p1` isminde `person` tipinde bir nesne
oluşturduk. `p1.sayHello()` yazarak `person` struct tipinde oluşturduğumuz
nesne için olan `sayHello` fonksiyonumuzu çalıştırdık.
*/
