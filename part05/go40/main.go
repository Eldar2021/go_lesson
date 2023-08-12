package main

import "fmt"

/// Print Fonksiyonu
/*
* `Print` fonksiyonu Go dilinde komut satırı üzerinde yazdırma işlemi yapmak
* için kullanılır. `Print` fonksiyonunun en çok kullanılan 3 çeşidine bakalım.
 */

func main() {
	// `Print()` Fonksiyonu
	// Bu fonksiyonun içine parametreler girerek ekrana yazdırma işlemi yapabiliriz.
	fmt.Print("Merhaba Dünya!")

	// `Println()` Fonksiyonu
	// Bu fonksiyon ile içine parametre girerek ekrana yazdırma işlemi
	// yapabiliriz. Yazdırma işlemi yaptıktan sonra bir alt satıra geçer.
	fmt.Println("satır1")
	fmt.Println("satır2")

	// `Printf()` Fonksiyonu
	// Gelelim işimizi görecek olan `Printf()` fonksiyonuna. Bu fonksiyon
	// sayesinde metinsel bölümlerin arasına değişken yerleştirebiliriz.
	dil := "Go"
	yil := 2007
	fmt.Printf("%s dili %d yılından beri geliştiriliyor.", dil, yil)
}
