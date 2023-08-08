package main

import (
	"fmt"
	"sync"
	"time"
)

/// WaitGroup ile Asenkron İşlemleri Beklemek
/*
* Go dilinde, `sync` paketinde bulunan `WaitGroup` tipi, asenkron olarak
* çalışan goroutine'ların tamamlanmasını beklemek ve senkronize etmek
* için kullanılır. `WaitGroup`, goroutine'lar arasında senkronizasyon
* sağlamak için oldukça kullanışlı bir araçtır ve ana iş parçacığının
* belirli goroutine'ların tamamlanmasını beklemesini sağlar.
*
* WaitGroup kullanımı aşağıdaki adımlarla gerçekleşir:
 */

/*
* 1) `sync` paketini içe aktarın.
* 2) Bir `WaitGroup` nesnesi oluşturun.
* 3) Her goroutine, işlemi başlatırken `Add` metodunu çağırarak bekleme
*    grubuna bir goroutine ekler.
* 4) Her goroutine, işlemi tamamladığında `Done` metodunu çağırarak bekleme
*    grubundan çıkar.
* 5) Ana iş parçacığı, Wait metodunu çağırarak tüm goroutine'ların
*    tamamlanmasını bekler.
 */

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 7; i++ {
		// Her bir goroutine için bekleme grubuna bir goroutine ekliyoruz
		wg.Add(1)

		go func(n int) {
			// İşlem tamamlandığında bekleme grubundan çıkar
			defer wg.Done()

			fmt.Println("Goroutine", n, "başladı")
			time.Sleep(time.Second)
			fmt.Println("Goroutine", n, "tamamlandı")
		}(i)
	}

	// Tüm goroutine'ların tamamlanmasını bekliyoruz
	wg.Wait()

	fmt.Println("Tüm goroutine'lar tamamlandı.")
}

/*
* Yukarıdaki örnekte, beş farklı goroutine oluşturuluyor ve her biri bir
* saniye boyunca uyutuluyor. WaitGroup kullanılarak ana iş parçacığı,
* tüm goroutine'ların tamamlanmasını bekliyor.
*
* WaitGroup, goroutine'ları senkronize etmek ve tüm işlemler tamamlandığında
* devam etmek gibi senkronizasyon senaryolarında oldukça yararlıdır.
* Bu sayede tüm goroutine'lar işlerini tamamladığında ana iş parçacığı
* devam eder ve asenkron işlemleri beklemek daha kolay hale gelir.
 */
