package main

import (
	"fmt"
	"sync"
)

/// Mutex
/*
* "Mutex" (Mutex, Mutual Exclusion) Go dilinde eşzamanlı çalışan
* goroutine'ların veriye güvenli ve düzenli bir şekilde erişmelerini
* sağlamak için kullanılan bir mekanizmadır. Mutex, verilere paralel
* erişimi kontrol etmek ve senkronizasyon sağlamak amacıyla kullanılır.
*
* Go dilinde "sync" paketi altında yer alan sync.Mutex türü, mutex
* mekanizmasını sağlar. Mutex'lar, belirli bir kod parçasının sadece
* tek bir goroutine tarafından aynı anda çalıştırılmasına izin vererek
* veri tutarlılığını korur.
 */

var count int
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	count++
	fmt.Println(count)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", count)
}
