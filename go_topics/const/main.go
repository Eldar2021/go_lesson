package main

import "fmt"

/// Golang Constants
/*
* Sabit, sabit bir değer için bir ad veya tanımlayıcıdır.
* Bir değişkenin değeri değişebilir, ancak bir sabitin değeri sabit kalmalıdır.
 */

const PRODUCT string = "Canada"
const PRICE = 500

const (
	QUANTITY = 50
	STOCK    = true
)

func main() {
	fmt.Println(PRODUCT)
	fmt.Println(PRICE)

	fmt.Println(QUANTITY)
	fmt.Println(STOCK)
}
