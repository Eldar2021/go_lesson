package main

import (
	"fmt"
	"time"
)

/// Anonim Goroutine Fonksiyonlar
/*
* "Anonim goroutine fonksiyonları", Go dilinde goroutine olarak çalıştırılmak
* üzere tanımlanan ve isimleri olmayan fonksiyonlardır. Bu tür fonksiyonlar,
* genellikle kodun belirli bir bölümünü paralel olarak çalıştırmak veya
* eşzamanlı işlemler yapmak için kullanılır. Anonim goroutine fonksiyonları,
* genellikle `go` anahtar kelimesi ile birlikte kullanılır ve anında
* goroutine olarak çalıştırılır.
*
* Anonim goroutine fonksiyonları şu şekilde tanımlanır ve kullanılır:
 */

func main() {
	// Anonim goroutine fonksiyonunu go anahtar kelimesi ile çalıştırma
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Goroutine:", i)
			time.Sleep(time.Second)
		}
	}()

	// Ana iş parçacığı bekler
	time.Sleep(10 * time.Second)

	fmt.Println("Program sonlandı.")
}

/*
* Yukarıdaki örnekte, `go func() { ... }()` şeklinde bir anonim goroutine
* fonksiyonu tanımlanmış ve çalıştırılmıştır. Bu anonim fonksiyon,
* bir goroutine içinde çalıştırılır ve ekrana birkaç kez "Goroutine: x"
* mesajını basar.
*
* Anonim goroutine fonksiyonları genellikle bir kod bloğunu paralel olarak
* çalıştırmak veya eşzamanlı işlem yapmak için kullanılır. Bu, kodun daha
* verimli çalışmasını ve paralel işlem gücünden yararlanmasını sağlar.
*
* Anonim goroutine fonksiyonlarına ek olarak, bu fonksiyonlar içinde
* değişkenlerin değerlerini yakalamak için dikkat etmeniz gereken bazı
* noktalar vardır. Çünkü goroutine'lar, ana iş parçacığından bağımsız olarak
* çalıştıkları için bu değişkenlerin senkronize ve güvenli bir şekilde
* erişilmesi gerekebilir.
*
* Anonim goroutine fonksiyonları, Go dilinde paralel ve eşzamanlı programlamada
* güçlü bir araçtır ve kodunuzun daha verimli çalışmasını sağlayabilir.
 */
