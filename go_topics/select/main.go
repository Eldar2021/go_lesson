package main

import (
	"fmt"
	"time"
)

/// Select
/*
* "Select", Go dilinde birden fazla iletişim kanalı üzerinden iletişim
* kurma yeteneğini sağlayan bir kontrol yapısıdır.
* Genellikle "goroutine"'lar arasındaki iletişimi koordine etmek ve
* senkronizasyon sağlamak için kullanılır. "Select" ifadesi, bir veya
* daha fazla iletişim operasyonunu beklemek ve ilgili operasyonun
* gerçekleşmesini beklemek için kullanılır.
*
* Go dilinde "select" ifadesi, genellikle "channel"'larla kullanılır.
* Birden fazla "channel" üzerinde beklemek ve hangi "channel" üzerinden
* veri gelirse o operasyonu gerçekleştirmek için kullanılır.
 */

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Eldiiar"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Almazbek"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch3 <- 22
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case msg3 := <-ch3:
			fmt.Println(msg3)
		}
	}
}
