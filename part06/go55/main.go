package main

import (
	"io"
	"log"
	"os"
)

/// Log (Kayıt)
/*
* Log paketi standart Golang paketleri içerisinde gelir ve programdaki olayları
* kaydetmemizi yarayacak bir altyapı sunar. Log programcının gözü kulağıdır.
* Bize hataları (bugs) bulmamız için kolaylık sağlar.
 */

func main() {
	/*
	* SetPrefix() fonksiyonu ile log çıktımızın satırının başında ne yazacağını
	* belirleyebiliyoruz.
	 */

	log.SetPrefix("KAYITIM: ")

	/*
	* SetFlags () fonksiyonu ile log çıktımızın görünüşünü ayarlıyoruz. log.Ldate
	* bize zamanını gösteriyor. log.Lmicroseconds mikrosaniyeyi ve log.Llongfile
	* ise dosyanın bulunduğu dizin ile ismini ve yapılan işlem ile ilgili satırı
	* gösteriyor.
	 */
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	/*
		Log kayıtlarımızı bir dosyaya kaytediyoruz
	*/
	logDosyam, _ := os.Create("kayit.log")
	defer logDosyam.Close()

	/*
	* multiwriter oluşturduk ve hem komut satırına hem log dosyamıza
	* yazdırması için ayarladık.
	 */
	mw := io.MultiWriter(os.Stdout, logDosyam)

	/*
	* log.SetOutput() belirdediğimiz dosyaya loglarımısı yazmak için kullanıyoruz
	* Artık loglarımız komut satırında gözükmek yerine belirttiğimiz
	* dosyaya yazılacaktır. `log.SetOutput(logDosyam)`
	 */
	log.SetOutput(mw) // log.SetOutput(logDosyam)

	/*
	* log.Println() fonksiyonu ile klasik log çıktılama işlemini yapıyoruz.
	* Fonksiyonun sonundaki ln bir alt satıra geçildğini gösteriyor.
	 */
	log.Println("Bir sıkıntı yaşadık ama mühim değil :/")

	/*
	* log.Fatalln() fonksiyonu ile kritik hataları bildirir. log.Println()
	* fonksiyonundan farkı program 1 çıkış kodu ile biter. Bu da programın hatalı
	* bittiği anlamına gelir. Normalde sağlıklı çalışan bir program 0 çıkış kodu
	* ile biter. 0 çıkış kodunu Golang programlama da kullanmamıza gerek kalmaz.
	 */
	log.Fatalln("Bu sefer çok kötü birşey oldu :<")
}
