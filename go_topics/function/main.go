package main

import "fmt"

/// Function
/*
* Bir ad bir harfle başlamalıdır ve herhangi bir sayıda ek harf ve rakam içerebilir.
* Bir fonksiyon adı bir sayı ile başlayamaz.
* Bir fonksiyon adı boşluk içeremez.
* İsmi büyük harfle başlayan fonksiyonlar diğer paketlere aktarılır. İşlev adı küçük harfle başlıyorsa, diğer paketlere aktarılmaz, ancak bu işlevi aynı paket içinde çağırabilirsiniz.
* Bir isim birden fazla kelimeden oluşuyorsa, ilk kelimeden sonraki her kelime şu şekilde büyük harfle yazılmalıdır: empName, EmpAddress, vb.
* fonksiyon isimleri büyük/küçük harf duyarlıdır (car, Car ve CAR üç farklı değişkendir).
 */

// Creating a Function in Golang
func SimpleFunction() {
	fmt.Println("Hello World")
}

// function with parameters
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

// Simple function with return
func kosh(x int, y int) int {
	total := 0
	total = x + y
	return total
}

// return values of a function can be named
func rectangle(l int, b int) (area int) {
	var parameter int = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)

	area = l * b
	return // Return statement without specify variable name
}

// Functions Returning Multiple Values
func rectangle02(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

// Passing Address to a Function
func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
}

func main() {

	SimpleFunction()

	add(2, 3)

	sum := kosh(20, 30)

	fmt.Println(sum)
	fmt.Println("Area: ", rectangle(20, 30))

	a, p := rectangle02(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)

	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)
	update(&age, &text)
	fmt.Println("After :", text, age)

	/// Anonymous Functions

	// Assigning function to the variable.
	var (
		area = func(l int, b int) int {
			return l * b
		}
	)

	fmt.Println(area(20, 30))

	// Passing arguments to anonymous functions.
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	// Function defined to accept a parameter and return value.
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	/// Closures Functions
	/*
	* Closure'lar anonim fonksiyonların özel bir durumudur. Closure'lar,
	* fonksiyonun gövdesi dışında tanımlanan değişkenlere erişen anonim
	* fonksiyonlardır.
	 */

	// Anonymous function accessing the variable defined outside body.
	l := 20
	b := 30

	func() {
		var area int = l * b
		fmt.Println(area)
	}()

	// Anonymous function accessing variable on each iteration of loop inside function body.
	for i := 10.0; i < 100; i += 10.0 {
		result := func() float64 {
			return i * 39.370
		}
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, result())
	}

}
