package main

import (
	"fmt"
	"time"
)

/// Kanallar (Channels)
/*
Kanallar, Go dilinde goroutine'ler arasında veri aktarmak için kullanılan
bir yapıdır. Kanallar, tek yönlü veya çift yönlü olabilir. Tek yönlü kanallar,
yalnızca veri göndermek veya yalnızca veri almak için kullanılabilir.
Çift yönlü kanallar, hem veri göndermek hem de almak için kullanılabilir.

Kanallar, `make()` işlevi kullanılarak oluşturulur. `make()` işlevine,
kanalın adı, kanalın tipi ve kanalın boyutu (isteğe bağlı) parametre
olarak verilir. Örneğin, aşağıdaki kod, `chan int` tipinde bir tek yönlü
kanal oluşturur: `ch := make(chan int)`

Kanallar, `<-` operatörü kullanılarak kullanılır.
`<-` --> operatörü, kanaldan veri alır yada verir.  `ch <- 10`

Kanallar, goroutine'ler arasında iletişim kurmak için çok etkili bir yoldur.
Kanallar, goroutine'ler arasında senkronizasyonu sağlamak için de kullanılabilir.
*/

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Salam duyno" // Kanala veri gönderiliyor
	}()

	m := <-ch // Kanaldan veri alınıyor

	fmt.Println(m)

	// Birkaç saniye bekleyerek programın sonlanmasını sağlıyoruz
	time.Sleep(1 * time.Second)
}

/*
* İşte kanallar hakkında bazı ek bilgiler:
* 1) Kanallar, goroutine'ler arasında veri aktarmak için kullanılır.
* 2) Kanallar, tek yönlü veya çift yönlü olabilir.
* 3) Kanallar, `make()` işlevi kullanılarak oluşturulur.
* 4) Kanallar, `<-` operatörü kullanılarak kullanılır.
* 5) Kanallar, goroutine'ler arasında senkronizasyon sağlamak için kullanılabilir.
 */
