package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/// HTML Şablonlar (Templates)
/*
* HTML Şablonlar, Golang üzerinde web sayfalarının dinamikliği için kullanılır.
* Yani şablonlar kullanarak web sayfalarımızın belirlediğimiz bölümlerini
* Go üzerinden değişikliğe uğratabiliriz.
*
* Bu yazımızda HTML şablonların nasıl oluşturulacağına bakacağız. Çalışma mantığı
* çok basit. Şablon olarak kullanacağımız html dosyasınında sadece tasarımı
* yapıyoruz ve dinamik olacak kısımlara ise bir nevi işaretler koyuyoruz.
* Daha sonra Go bu şablon dosyasını işliyor ve işaret koyduğumuz yerlere gelecek
* değerleri yerleştiriyor. Düz mantık olarak bu işi yapıyor.
 */

// yakalayıcı fonksiyonumuz
func homePage(w http.ResponseWriter, r *http.Request) {
	// isim değişkenimiz
	p1 := Person{Name: "Eldiiar", LastName: "Almazbek", Age: 22}

	// burada şablon oluşturuyoruz
	homeTemplate, _ := template.ParseFiles("home.html")

	// Burada da şablonu çalıştırmasını ve isim
	// değişkenini kullanmasını istiyoruz.
	homeTemplate.Execute(w, p1)
}

func main() {
	fmt.Println("Start Server")

	// ana dizini anasayfa fonksiyonu ile yakalayalım
	http.HandleFunc("/", homePage)

	// portu 8000 yapalım ve sunucuyu başlatalım
	http.ListenAndServe(":8000", nil)
}

type Person struct {
	Name     string
	LastName string
	Age      int
}
