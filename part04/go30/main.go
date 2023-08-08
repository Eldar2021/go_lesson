package main

import (
	"fmt"
	"time"
)

/// Goroutine
/*
Goroutine, Go dilinde kullanılan hafif iş parçacığıdır.
Goroutine'ler, diğer iş parçacıkları ile aynı belleği paylaşır ve
aynı belleğe aynı anda erişebilirler. Bu, goroutine'leri paralel
programlama için ideal hale getirir.

Goroutine'ler, `sayHello` anahtar kelimesi kullanılarak oluşturulur.
Örneğin, aşağıdaki kod, `sayHello()` işlevini çalıştıran bir goroutine oluşturur:
*/

func main() {
	fmt.Println("Program başlandı.")

	go sayHello() // Goroutine olarak sayHello fonksiyonunu çalıştırır

	time.Sleep(1 * time.Second)

	fmt.Println("Program sonlandı.")
}

func sayHello() {
	fmt.Println("Hello!")
}

/*
* İşte goroutine'ler hakkında bazı ek bilgiler:
* 1) Goroutine'ler, çok hafiftir ve çok az bellek kullanır.
* 2) Goroutine'ler, diğer goroutine'lerle aynı belleği paylaşır.
* 3) Goroutine'ler, kanallar kullanılarak iletişim kurabilirler.
* 4) Goroutine'ler, Go dilinde programlama yapmanın çok etkili bir yolunu sağlar.
 */
