package main

import "fmt"

/// Type Assertions
/*
* "Type Assertions" (Tür İddiaları), Go dilinde bir arayüze sahip olmayan
* (interface{} türü) bir değerin belirli bir türe dönüştürülmesini veya
* kontrol edilmesini sağlayan bir mekanizmadır. Bu, veri türünü belirlemek
* veya belirli bir türe dönüştürmek istediğiniz durumlar için kullanılır.
*
* Tür iddiaları, çoğunlukla dinamik tür bilgisine ihtiyaç duyduğunuz yerlerde
* ve dinamik olarak tür değiştirmeniz gereken durumlarda kullanılır.
* Ancak, tür iddiaları kullanırken dikkatli olunması gereken bazı önemli
* noktalar vardır.
 */

// Temel Tür İddiası:
func main1() {
	var val interface{} = "Eldi"
	str, ok := val.(string)

	if ok {
		fmt.Println("Değer bir string:", str) // Değer bir string: Eldi
	} else {
		fmt.Println("Değer bir string değil.")
	}
}

// Dinamik Tür Dönüşümü:

func printType(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("int türü:", t)
	case string:
		fmt.Println("string türü:", t)
	case float32:
		fmt.Println("float32 türü:", t)
	case float64:
		fmt.Println("float64 türü:", t)
	default:
		fmt.Println("Bilinmeyen tür:", t)
	}
}

func main2() {
	printType(42)
	printType("Merhaba, Go!")
	printType(3.14)
	printType(true)
}

// run all code
func main() {
	main1()
	main2()
}
