package main

import (
	"fmt"
	"reflect"
)

/// Variadic Functions
/*
* Variadic fonksiyon, değişken sayıda argüman kabul eden bir fonksiyondur.
* Golang'da, fonksiyon imzasında belirtilenle aynı türden değişen sayıda
* argüman geçmek mümkündür. Varyadik bir fonksiyon bildirmek için, son
* parametrenin türünden önce, fonksiyonun bu türden herhangi bir sayıda
* argümanla çağrılabileceğini gösteren bir üç nokta, "..." konur. Bu tür
* fonksiyonlar, fonksiyona aktardığınız argümanların sayısını bilmediğiniz
* durumlarda kullanışlıdır; buna en iyi örnek, değişken bir fonksiyon olan
* fmt paketinin yerleşik Println fonksiyonudur.
 */

func variadicExample01(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

// Passing multiple string arguments to a variadic function
func variadicExample02(s ...string) {
	fmt.Println(s)
}

// Normal function parameter with variadic function parameter
func variadicExample03(str string, y ...int) int {

	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

// Pass different types of arguments in variadic function
func variadicExample04(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func main() {
	variadicExample01("red", "blue", "green", "yellow")

	variadicExample02()                                 // []
	variadicExample02("red", "blue")                    // [red blue]
	variadicExample02("red", "blue", "green")           // [red blue green]
	variadicExample02("red", "blue", "green", "yellow") // [red blue green yellow]

	fmt.Println(variadicExample03("Rectangle", 20, 30))
	fmt.Println(variadicExample03("Square", 20))

	variadicExample04(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})
}
