package main

import (
	"context"
	"fmt"
	"time"
)

/// Context
/*
* Tabii ki, "context" (bağlam) kavramı, Go dilinde paralel işlemler
* ve ağ işlemleri gibi senaryolarda kullanılan ve değerleri taşımak,
* paylaşmak ve zamanlamak için kullanılan bir mekanizmadır.
* "context" paketi, bu tür senaryolarda veri iletimini ve yönetimini
* kolaylaştırmak için kullanılır.
*
* "Context" kavramı, genellikle HTTP sunucularında, veritabanı sorgularında,
* istek/yanıt işlemlerinde ve diğer asenkron işlemlerde kullanılır.
* İşlemler sırasında paylaşılan değerleri ve istekleri taşımak ve kontrol
* etmek için kullanışlıdır.
 */

func main() {
	// Ana bir "context" oluşturuyoruz
	ctx := context.Background()

	// Bu ana "context"ten türetilmiş bir alt "context" oluşturuyoruz
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	// Alt "context" üzerinde çalışan bir fonksiyon çağrısı
	doSomething(ctxWithTimeout)

}

func doSomething(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("İşlem zaman aşımına uğradı veya iptal edildi.")
	default:
		fmt.Println("İşlem tamamlandı.")
	}

}
