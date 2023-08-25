package main

import (
	// kütüphanemizi içeri aktaralım
	"github.com/gin-gonic/gin"
)

/// Gin Web Kütüphanesi
/*
* Gin Go'da yazılmış bir web kütüphanesidir. Performans ve üretkenlik odaklıdır.
* Sizlere basitçe web sunucu ve api oluşturmanız için kolaylık sağlar.
*
* Kurulum için: `go get -u github.com/gin-gonic/gin`
*
* Basit bir web sunucu oluşturma örneği:
 */

// anasayfayı yakalayacak olan fonksiyonumuz
func indexResponse(ctx *gin.Context) {
	//c ile gin nesnemize bağlam oluşturduk.
	//c'yi kullanarak artık gin özelliklerine erişebiliriz.

	//sayfaya düz yazı gönderdik
	ctx.String(200, "Merhaba Go! Benim ilk projeme hosh geldiniz!\nstring\njson")
	//Buradaki 200 sunucudan bir cevap geldiğini anlamına gelir
}

func stringResponse(ctx *gin.Context) {
	ctx.String(200, "Eldi Salam")
}

func jsonResponse(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"name":     "Eldiiar",
		"lastName": "Almazbek",
		"age":      22,
	})
}

func htmlResponce(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"title": "Hi I'm Flutter Developer",
	})
}

func main() {
	//gin'in varsayılan ayarlarında bir yönlendirici oluşturalım.
	router := gin.Default()

	//Burada templates klasörünün içindeki tüm şablonları yüklemesini isteyelim.
	router.LoadHTMLGlob("templates/*")

	//anasayfayı inde fonksiyonumuz yakalayacak
	router.GET("/", indexResponse)
	router.GET("/string", stringResponse)
	router.GET("/json", jsonResponse)
	router.GET("/html", htmlResponce)

	//daha sonra sunucuyu başlatıyoruz
	router.Run()
}
