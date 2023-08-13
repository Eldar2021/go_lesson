package main

import (
	"fmt"
	"strings"
)

/// Strings
/*
* `Strings` paketi ile `string` türünde değerler üzerinde işlemler yapabiliriz.
* Kısaca kullanımlarından bahsedelim.
 */

func main() {
	var email string
	fmt.Println("Please enter email-address : ")
	fmt.Scan(&email)

	if strings.Contains(email, "@") {
		fmt.Println("E-posta Adresiniz Onaylandı!")

		emailSymbolCount := strings.Count(email, "@")
		fmt.Printf("@ işaretinden %d kullandınız\n", emailSymbolCount)

		emailSymbolIndex := strings.Index(email, "@")
		fmt.Printf("@ işareti %d indexsindedir\n", emailSymbolIndex)

		fmt.Println(strings.ToUpper("merhaba dünya"))

		fmt.Println(strings.ToLower("Merhaba Dünya"))
	} else {
		fmt.Println("Geçerli Bir E-posta Adresi Giriniz!")
	}

}
