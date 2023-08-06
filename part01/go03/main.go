package main

import "fmt"

/// Aritmetik Operatörler
/*
Aritmetik operatörler programlamada matematiksel işlemler yapabilmemize olanak sağlar.
+ -> toplar
- -> Çıkarır
* -> Çarpar
/ -> Böler
% -> Bölümden kalanı verir (Mod)
++ -> 1 arttırır
-- -> 1 eksiltir
*/

func main() {
	var a = 12
	var b = 8
	var c = a + b
	fmt.Println(c)

	var d = 12
	var e = 8
	var f = d - e
	fmt.Println(f)

	var g = 12
	var h = 8
	var l = g * h
	fmt.Println(l)

	var m = 12
	var n = 8
	var o = m / n
	fmt.Println(o)

	var p = 12
	var r = 8
	var s = p % r
	fmt.Println(s)

}
