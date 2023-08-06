package main

import "fmt"

/// Struct
/*
Go programlama dilinde sınıflar yoktur.
Sınıflar yerine struct'lar (yapılar) vardır.
Yapılar sayesinde bir nesne oluşturabilir ve bu
nesneye ait özellikler oluşturabiliriz. Örnek bir struct oluşturalım.
*/

// 1) `type` terimi ile yeni bir tür oluşturabiliyoruz.
// 2) `Person` bizbergen at
// 3) türünün de `struct` olacağını söyledik
type person struct {
	name     string
	lastname string
	age      int
}

func main() {
	person01 := person{"Eldi", "Almazbek", 23}
	fmt.Println(person01) // {Eldi Almazbek 23}

	person01.name = "Ali"
	person01.lastname = "Asan"
	person01.age = 22

	fmt.Println(person01) // {Ali Asan 22}

	// -------------------------------------------------
	person02 := person{}
	fmt.Println(person02) // {  0}

	person02.name = "Chris"
	person02.lastname = "Lang"
	person02.age = 32

	fmt.Println(person02) // {Chris Lang 32}

	person03 := person{age: 43, name: "Alan", lastname: "Brezi"}
	fmt.Println(person03) // {Alan Brezi 43}
}
