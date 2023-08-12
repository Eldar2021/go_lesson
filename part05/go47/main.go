package main

import "fmt"

/// Panic & Recover
/*
* `Panic` ve `Recover`, Golang’de hata ayıklama için kullanılan anahtar kelimelerdir.
* Size bunu daha iyi ve akılda kalıcı anlatmak için teorik anlatım yerine uygulamalı
* öğretim yapmak istiyorum. Böylece daha akılda kalıcı olur.
*
* Aşağıda `panic` durumu oluşturan bir örnek göreceğiz:
 */

func fullname(name *string, lastname *string) {
	if name == nil {
		panic("name can not be nil")
	}

	if lastname == nil {
		panic("lastname can not be nil")
	}

	fmt.Printf("%s %s\n", *name, *lastname)
	fmt.Println("TamIsim fonksiyonu bitti")
}

func main() {
	name := "Eldiiar"
	// lastname := nil

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("No error. :) Just  name or lastname con not be nil")
		}
	}()

	fullname(&name, nil)
	fmt.Println("Ana fonksiyon da bitti")
}
