package main

import "fmt"

/// Import (Kütüphane Ekleme)

/*
Bu yazıda sizlere Golang’ta paket import etmenin tüm yöntemlerini göstereceğim.

1)fmt paketini import ettik.
`import "fmt"`

2) Birden fazla paket import ettik
`import (
    "fmt"
    "net/http"
)`

3) fmt paketini import edip f olarak kullanacağımızı belirttik. Örnek olarak
fmt.Println() yazmak yerine f.Println() yazacağız.
`import f "fmt"`

4) Dikkat ederseniz, import kelimesinden sonra nokta koyduk. Bu işlem sayesinde
fmt.Println() yazmak yerine sadece Println() yazarak aynı işi yapmış oluruz.
`import . "fmt"`

5) Bazen Golang yazarken kütüphaneyi ekleyip kullanmadığımız zamanlar olur.
Böyle durumlarda program çalıştırılırken veya derlenirken kullandığınız
editör veya ide bu bölümü silebilir. import ederken _ (alt tire) koyarak
bunun üstesinden gelebiliriz.
`import _ "fmt"`
*/

func main() {
	fmt.Println("Go Import lesson")
}
