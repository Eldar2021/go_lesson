package main

import (
	"fmt"
	"sync"
	"time"
)

/// "Mutex" ile asenkron işlem sırası
/*
* Go dilinde, "Mutex" (kilit mekanizması) kullanarak belirli kod bloklarının
* aynı anda sadece bir goroutine tarafından çalıştırılmasını sağlayabilirsiniz.
* Mutex, "Mutual Exclusion" (Karşılıklı Dışlama) kavramının kısaltmasıdır ve
* birden fazla goroutine arasındaki eşzamanlı erişimi kontrol etmek için kullanılır.
 */

// Mutex kullanımı aşağıdaki adımlarla gerçekleşir:
/*
* 1) `sync` paketini içe aktarın.
* 2) Bir `Mutex` nesnesi oluşturun.
* 3) Paylaşılan kaynağa erişim yapmadan önce `Lock` metodunu çağırarak kilidi alın.
* 4) İşlem tamamlandığında `Unlock` metodunu çağırarak kilidi serbest bırakın.
 */

func main() {
	var mt sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			// Kilit alınır
			mt.Lock()

			// İşlem tamamlandığında kilit serbest bırakılır
			defer mt.Unlock()

			fmt.Println("Goroutine", n, "kilit aldı")
			time.Sleep(time.Second)
			fmt.Println("Goroutine", n, "kilit serbest bıraktı")
		}(i)
	}

	wg.Wait()
}
