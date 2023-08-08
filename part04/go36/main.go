package main

import (
	"fmt"
	"time"
)

/// Select
// `Select` ile çoklu goroutine işlemlerinin iletişimini bekleyebiliriz.

func main() {
	k1 := make(chan string)
	k2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		k1 <- "Ses"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		k2 <- "Video"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-k1:
			fmt.Println("k1 is :", m1)
		case m2 := <-k2:
			fmt.Println("k2 is :", m2)
		}
	}
}

/*
* Yukarıdaki kodların bize ses ve video verisi sağlayacak bir programdan
* parça olduğu senaryosunu kuralım. Bu programda işlem yapabilmemiz için
* bize bu 2 verinin gelmesini beklememiz lazım. Verileri bekleme işlemini
* select ile yapıyoruz. Burada dikkat etmemiz gereken nokta 2 tane veri
* beklediğimiz için  for atamalarında i < 2 olarak girmeliyiz. Çünkü i := 0
* olduğu için i 2 olana kadar arada 2 sayı var. Bu sayı boşluğu da 2 veri
* almayı beklememizi sağlıyor. Örnek olarak i < 1 girip 2 veri almaya kalksak
* k2‘den gelen veriyi beklemeyecek bile. Tam tersi olarak 2 veri alacağımız
* halde i < 4 girsek program deadlock‘a girecektir. Yani başarısız bir program
* olacaktır.
 */
