package main

import (
	"fmt"
	"strconv"
)

/// Strconv (String Çeviri)
/*
* `strconv` paketi Golang ile bütünleşik gelen string tipi ve diğer tipler
* arasında çevirme işlemi yapabileceğimiz bir pakettir.
 */
func main() {
	//basit string-int arası çevirme
	sayi, _ := strconv.Atoi("-12") // string > int
	yazi := strconv.Itoa(-21)      // int > string

	//string'ten diğerlerine çevirme
	b, _ := strconv.ParseBool("true")      // string > bool
	i, _ := strconv.ParseInt("32", 10, 64) // string > int
	f, _ := strconv.ParseFloat("32.2", 64) // string > float
	u, _ := strconv.ParseUint("4", 10, 64) // string > uint

	//diğerlerinden string'e çevirme
	s1 := strconv.FormatBool(true)               // bool > string
	s2 := strconv.FormatInt(-32, 10)             // int > string
	s3 := strconv.FormatFloat(23.3, 'E', -1, 64) // float > string
	s4 := strconv.FormatUint(32, 16)             // uint > string

	fmt.Printf("sayi: %d tip: %T\n", sayi, sayi)
	fmt.Printf("yazi: %s tip: %T\n", yazi, yazi)
	fmt.Printf("b: %t tip: %T\n", b, b)
	fmt.Printf("f: %f tip: %T\n", f, f)
	fmt.Printf("i: %d tip: %T\n", i, i)
	fmt.Printf("u: %d tip: %T\n", u, u)
	fmt.Printf("%T %T %T %T", s1, s2, s3, s4)
}
