package main

import (
	"fmt"
	"time"
)

/// Anonim Goroutine Fonksiyonlar
/*
Bu yazımız Goroutine ve Kanallar dersi için biraz alıştırma tadında olacak.

Daha önceki yazılarımızda belirli bir fonksiyonu Goroutine ile asenkron
(eş zamanlı) olarak çalıştırmayı gördük. Bu yazımızda da anonim bir
Goroutine fonksiyonunu göreceğiz. Bu fonksiyonun özelliği bir ismi
olmaması ve asenkron olarak çalışmasıdır.
*/

func main() {
	ch := make(chan string) // kanal oluşturuyoruz

	go func() {
		time.Sleep(1 * time.Second) // 2 saniye uyku
		ch <- "Eldiiar"             // İletişime geçiriyoruz
		fmt.Println("Anonim fonksiyon yazısı")
	}()

	fmt.Println("Öylesine bir yazı")
	fmt.Println(<-ch) // kanaldan gelen veri bekleniyor
}
