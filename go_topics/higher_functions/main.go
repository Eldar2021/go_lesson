package main

import "fmt"

/// Higher Order Functions
/*
* Yüksek dereceli bir fonksiyon, bir fonksiyonu argüman olarak alan
* veya fonksiyonu çıktı olarak döndüren bir fonksiyondur.
*
* Daha yüksek dereceli fonksiyonlar, argüman olarak alarak ya da geri
* döndürerek diğer fonksiyonlar üzerinde çalışan fonksiyonlardır.
 */

// Passing Functions as Arguments to other Functions
func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

// Returning Functions from other Functions
func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

// User Defined Function Types
type First func(int) int
type Second func(int) First

func squareSum02(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	partial := partialSum(3)
	fmt.Println(partial(7))

	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))

	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum02(5)(6)(7))
}
