package main

import "fmt"

/// `init()` Fonksiyonu (Ön Yükleme)
/*
* Golang’te bir uygulama çalışırken genelde çalışan ilk fonksiyon main()
* fonksiyonu oluyor. Bazen programın açılışında ayarlamamız gereken ön
* durumlar oluşuyor. İşte init() fonksiyonu bize bu imkanı sunuyor.
* Ufak bir örnekle yazdıklarıma anlam katalım.
 */
func init() {
	fmt.Println("Init fonksyonu çalıştı")
}

func main() {
	fmt.Println("Main fonksyonu çalıştı")
}

/*
Init fonksyonu çalıştı
Main fonksyonu çalıştı
*/
