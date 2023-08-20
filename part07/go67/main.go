package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/// Web Scrapper (goquery)
/*
* `goquery`, Go programlama dilinde web sayfalarından veri çekmek ve içerik analizi
* yapmak için kullanılan bir kütüphanedir. HTML belgelerini işlemek, seçici
* (selector) tabanlı sorgular kullanarak belge içeriğini aramak ve çıkarmak için
* kullanışlı bir arabirim sunar. Bu kütüphane, web scraping (web kazıma) görevlerini
* kolaylaştırmak ve veri madenciliği işlemlerini gerçekleştirmek için yaygın olarak
* kullanılır.
 */
/*
* `goquery`'nin bazı avantajları şunlar olabilir:
*
* 1) Kolay Kullanım: goquery, jQuery benzeri bir seçici tabanlı yapı kullanır,
*    bu da HTML belgeleri içindeki öğelere kolayca erişmeyi sağlar.
*
* 2) Veri Kazıma: Web sayfalarından veri çekme, tablo içeriklerini veya listeleri
*    çıkarma gibi işlemleri hızlı ve basit bir şekilde yapmanızı sağlar.
*
* 3) Veri Analizi: Web sayfalarının yapısını analiz ederek belirli veri parçalarını
*    çıkarma veya dönüşüm yapma işlemlerini kolaylaştırır.
*
* 4) Seçici Tabanlı Sorgular: HTML belgesindeki öğelere CSS selektörleri (seçici)
*    kullanarak erişebilirsiniz, bu da işlemi daha esnek ve hedefe yönelik hale getirir.
 */

func main() {
	webUrl := "https://flutter.dev/"
	response, err := http.Get(webUrl)

	controlError(err)
	defer response.Body.Close()

	fmt.Println(response.Body)

	fmt.Println("<----------------------------------->")
	doc, err := goquery.NewDocumentFromReader(response.Body)
	controlError(err)
	fmt.Println(doc)

	fmt.Println("<----------------------------------->")
	h2 := doc.Find("h3")

	h2.Each(printItem)
	fmt.Println("<----------------------------------->")
}

func controlError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printItem(index int, item *goquery.Selection) {
	title := strings.TrimSpace(item.Text())
	fmt.Println("Başlık", index+1, ":", title)
}
