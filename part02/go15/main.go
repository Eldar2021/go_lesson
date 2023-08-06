package main

import "fmt"

/// Döngüler
/*
Programlama ile uğraşan arkadaşlarımızın da bileceği üzere,
programlama dillerinde while, do while ve for döngüleri vardır.
Bu döngüler ile yapacağımız işlemin belirli koşullarda tekrarlanmasını
sağlayabiliriz. Golang’ta ise diğer dillerin aksine sadece for döngüsü vardır.
Ama bu while ve do while ile yapılanları yapamayacağımız anlamına gelmiyor.
Golang’taki for döngüsü ile hepsini yapabiliriz. Yani dilin yapımcıları
tek döngü komutu ile hepsini yapabilmemize olanak sağlamışlar.

Gelelim for döngüsünün kullanımına. Go’da for döngüsü parametreleri
parantez içine alınmaz.
*/

func main() {
	// For Loop
	/*
		For döngüsünden ayrı olarak deger adında 0 sayısal
		değerini alan bir değişken oluşturduk. For döngüsünde
		ise sadece koşul parametresini belirttlik. Yani döngü deger
		değişkeni 10 sayısından küçük olduğu zaman çalışacak.
		For kod bloğu içerisinde her döngü tekrarlandığında deger
		değişkeni ekrana basılacak ve deger değişkenine +1 eklenecek.
	*/
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}

	// While (SADECE KOŞUL BELİRTEREK KULLANMA)
	/*
		Bu for yazım şekli while mantığı gibi çalışır.
		Parametrelerde sadece koşul belirtilir.

		For döngüsünden ayrı olarak deger adında 0 sayısal değerini
		alan bir değişken oluşturduk. For döngüsünde ise sadece koşul
		parametresini belirttlik. Yani döngü deger değişkeni 10
		sayısından küçük olduğu zaman çalışacak. For kod bloğu içerisinde
		her döngü tekrarlandığında deger değişkeni ekrana basılacak ve
		deger değişkenine +1 eklenecek.
	*/

	b := 10
	for b < 20 {
		fmt.Println(b)
		b++
	}
}
