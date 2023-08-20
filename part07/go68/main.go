package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

/// `chromedp` (Web Driver)
/*
* `chromedp`, Go programlama dili için geliştirilmiş bir kütüphanedir ve Google Chrome
* tarayıcısını otomatikleştirmek ve kontrol etmek için kullanılır. Bu kütüphane,
* Chrome'un tarayıcı otomasyon protokolünü kullanarak tarayıcı işlemlerini programatik
* olarak yürütmenizi sağlar. chromedp, web sayfaları üzerinde gezinme, form doldurma,
* içerik çekme, tarayıcı ayarlarını değiştirme gibi birçok tarayıcı tabanlı otomasyon
* görevini gerçekleştirmenize olanak tanır.
*
* `chromedp`'nin bazı avantajları şunlar olabilir:
*
* 1) Tam Kontrol: Web tarayıcısını tamamen programatik olarak kontrol edebilirsiniz.
*    Sayfaları açabilir, form doldurabilir, tıklamalar yapabilir ve diğer etkileşimleri
*    gerçekleştirebilirsiniz.
*
* 2) Veri Kazıma: Sayfalardan veri çekebilirsiniz. Özellikle JavaScript tarafından
*    oluşturulan içeriği çekebilmek için kullanışlıdır.
*
* 3) Sahne Görüntüsü Alma: Web sayfalarının görüntülerini (screenshot) alabilirsiniz.
*
* 4) Sayfa Etkileşimi: Web sayfaları üzerinde senaryoları taklit ederek test
*    edebilirsiniz.
*
* 5) Dinamik İçerik Yönetimi: Dinamik olarak yüklenen veya oluşturulan içeriklere
*    erişebilirsiniz.
*
* go get -u github.com/chromedp/chromedp
 */

func contronError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()

	// Tarayıcıyı başlatma
	err := chromedp.Run(ctx, chromedp.Navigate("https://github.com/Eldar2021"),
		// Sayfanın yüklenmesini bekliyoruz
		chromedp.Sleep(2*time.Second))

	contronError(err)

	// Sayfa içeriğini alıyoruz
	var pageTitle string
	err = chromedp.Run(ctx, chromedp.Title(&pageTitle))

	contronError(err)

	fmt.Println("Sayfa Başlığı:", pageTitle)
}
