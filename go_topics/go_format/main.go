package main

import "fmt"

/// Format ve Kaçış Karakterleri
/*
* Format Karakterleri ve Kullanım Alanları
* Format karakterleri metinsel bir ifadede (string), dizgiyi formatlandırmak
* için kullanılır. Yani bir metinde değişken yerleri biçimlendirmeye yarar.
 */

/*
* %T --> Değişkenin tipini verir
* %t --> Boolean değeri verir
* %d --> Int (tamsayı) değeri verir
* %b --> Sayının binary (ikili) karşılığını verir
* %c --> Karakter değerini verir
* %x --> Sayının hexadecimal (onaltılı) karşılığını verir
* %f --> Float (ondalıklı) değeri verir
* %s --> String (dizgi-metin) değeri verir
* %v --> Değeri otomatik belirler
 */

/*
* \a --> Komut satırında zil sesi çıkartır
* \b --> Silme tuşu görevini görür
* \f --> Merdiven metin yazar
* \n --> Yeni satıra geçer
* \r --> Return eder
* \t --> Tab tuşu gibi boşluk bırakır (4 boşluk)
* \v --> Dikey boşluk bırakır
* \  --> Ters-taksim yazar
* \' --> Tek tırnak yazar
* \" --> Çift tırnak yazar
 */

func main() {
	isim := "Kaan"
	yaş := 23
	kilo := 71.3
	evli := false

	fmt.Printf("İsim: %s, Yaş: %d, Kilo: %f, Evli: %t", isim, yaş, kilo, evli)

	fmt.Print("Bir\nİki\tÜç\\Dört")
}
