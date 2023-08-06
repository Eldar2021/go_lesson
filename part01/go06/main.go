package main

import "fmt"

/// Atama Operatörleri
/*
Atama operatörleri değişkenlere ve sabitlere değer atamak için kullanılır.
Aşağıdaki tabloda c’nin değeri 10’dur. (c=10)

=  -> Atama Operatörüdür (c = 2)
+= -> Kendiyle toplar (c += 2)
-= -> Kendinden çıkarır (c -= 2)
*= -> Kendiyle çarpar (c *= 2)
/= -> Kendine böler (c /= 2)
%= -> Kendine bölümünden kalanı atar (c %= 3)
*/

func main() {
	var c int = 2
	fmt.Println(c)
	c += 2
	fmt.Println(c)
	c -= 2
	fmt.Println(c)
	c *= 2
	fmt.Println(c)
	c /= 2
	fmt.Println(c)
	c %= 2
	fmt.Println(c)
}
