package main

import "fmt"

/// Tür Dönüşümü

/*
Tür dönüşümü şu şekilde gerçekleştirilir.
tür(değer)
Örnek olarak bakmak gerekir ise;
*/

func main() {
	i := 42
	f := float64(i)
	u := uint(f)

	// Yukarıdaki yapılan işlemleri açıklayacak olursak eğer,
	// i adında bir `int` değişken tanımladık. f adındaki değişkende
	// i değişkenini `float64` türüne dönüştürdük. u adındaki değişkende
	// ise f değişkenini `uint` türüne çevirdik.

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)

	// Tüm türler arasında bu şekilde dönüşüm gerçekleştiremezsiniz.
	// Bir sayıyı string tipine dönüştürmek istediğimizde ne olacağına bakalım.

	a := 8378
	b := string(rune(a))
	fmt.Println(b)
}
