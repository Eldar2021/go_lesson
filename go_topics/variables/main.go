package main

import (
	"fmt"
	"reflect"
)

/// Variables
/*
Golang Değişkenleri için Adlandırma Kuralları
Bir Golang değişkenini adlandırmak için aşağıdaki kurallar geçerlidir:

Bir ad bir harfle başlamalıdır ve herhangi bir sayıda ek harf ve rakam içerebilir.
Bir değişken adı bir sayı ile başlayamaz.
Bir değişken adı boşluk içeremez.
Bir değişkenin adı küçük harfle başlıyorsa, bu değişkene yalnızca geçerli paket içinde erişilebilir ve bu, dışa aktarılmamış değişkenler olarak kabul edilir.
Bir değişkenin adı büyük harfle başlıyorsa, bu değişkene geçerli paketin dışındaki paketlerden erişilebilir, bu da dışa aktarılan değişkenler olarak kabul edilir.
Bir isim birden fazla kelimeden oluşuyorsa, ilk kelimeden sonraki her kelime şu şekilde büyük harfle yazılmalıdır: empName, EmpAddress, vb.
Değişken adları büyük/küçük harf duyarlıdır (car, Car ve CAR üç farklı değişkendir).
*/

func main() {
	// -----------------------------------------
	// Declaration with initialization
	var a int = 10
	var b string = "Canada"

	fmt.Println(a)
	fmt.Println(b)

	// -----------------------------------------
	// Declaration without initialization
	var c int
	var d string

	c = 10
	d = "Canada"

	fmt.Println(c)
	fmt.Println(d)

	// -----------------------------------------
	// Declaration with type inference
	var e = 10
	var f = "Canada"

	fmt.Println(reflect.TypeOf(e)) // int
	fmt.Println(reflect.TypeOf(f)) // string

	// -----------------------------------------
	// Declaration of multiple variables
	var fname, lname string = "John", "Doe"

	fmt.Println(fname + lname)

	// -----------------------------------------
	// Short Variable Declaration in Golang
	name := "John Doe"

	fmt.Println(reflect.TypeOf(name)) // string

	// -----------------------------------------
	// Zero Values
	var quantity float32
	fmt.Println(quantity) // 0

	var price int16
	fmt.Println(price) // 0

	var product string
	fmt.Println(product) // ""

	var inStock bool
	fmt.Println(inStock) // false

	// -----------------------------------------
	// Golang Değişken Bildirim Bloğu
	/* Değişken bildirimleri, daha fazla okunabilirlik
	* ve kod kalitesi için bloklar halinde gruplandırılabilir.
	 */

	var (
		p  = "Mobile"
		q  = 50
		pr = 50.50
		in = true
	)

	fmt.Println(q)
	fmt.Println(p)
	fmt.Println(pr)
	fmt.Println(in)
}
