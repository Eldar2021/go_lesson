package main

import "fmt"

/// Anonim Struct'lar
/*
Golang’ta tıpkı anonim fonksiyonlar olduğu gibi anonim
`struct`` methodlar da oluşturabiliriz. Örneğimizi görelim:
*/

func main() {
	person01 := struct {
		name, lastname string
	}{"Eldi", "Almazbek"}

	fmt.Println(person01) // {Eldi Almazbek}
}
