package main

import (
	"fmt"
	"time"
)

/// Zamanlayıcılar (Tickers)
/*
* Go dilinde "time" paketi içerisinde bulunan `Ticker` tipi, belirli
* bir süre boyunca periyodik olarak işlemleri tetiklemek için kullanılır.
* `Ticker`, belirtilen aralıklarla belirli bir görevi otomatik olarak
* çalıştırmanızı sağlar.
 */

/*
* `Ticker` kullanımı aşağıdaki adımlarla gerçekleşir:
* 1) `time` paketini içe aktarın.
* 2) Bir `Ticker` nesnesi oluşturun ve belirlediğiniz aralığı belirtin.
* 3) `Ticker` nesnesinin `C` özelliğini kullanarak zamanlayıcıdan okuma yapın.
* 4) `Ticker` kullanımı tamamlandığında `Ticker` nesnesini durdurun.
 */

func main() {
	// 1 saniye aralıklarla zamanlayıcı oluşturuldu
	ticcKer := time.NewTicker(1 * time.Second)

	// Ticker kapatıldığında belleği temizle
	defer ticcKer.Stop()

	for i := 0; i < 10; i++ {
		select {
		case <-ticcKer.C:
			fmt.Println("Zamanlayıcı tetiklendi :)", i)
		case <-ticcKer.C:
			fmt.Println("Zamanlayıcı tetiklendi :(", i)
		}
	}
}
