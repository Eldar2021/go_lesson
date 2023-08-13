package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

/// os/exec (Komut Satırına Erişim)
/*
* os/exec paketi komut satırına (cmd, powershell, terminal) komut göndermemizi
* sağlayan Golang ile bütünleşik gelen bir pakettir. Bu paket sayesinde oluşturacağımız
* programa sistem işlerini yaptırabiliriz. Örnek olarak dosya/klasör
* taşıma/silme/oluşturma/kopyalama gibi işlemleri yaptırabilir. Daha doğrusu
* komut satırı/terminal üzerinden yapabildiğimiz her işlemi yaptırabiliriz.
* Tabi kullandığımız işletim sistemine göre terminal komutları değiştiği için ona
* göre örnek vermeye çalışacağım.
 */

func Exmaple() {
	// Komut Satırına Komut Gönderme
	/*
	* “mkdir klasörüm” komutu programın çalıştırıldığı dizinde “klasörüm” adında
	* bir klasör oluşturur. Komut girerken dikkat etmeniz geren çok önemi bir detay
	* var. Yazacağınız komut birden fazla kelimeden oluşuyorsa mutlaka ayrı ayrı
	* girmelisiniz. Eğer exec.Command() fonksiyonuna direkt olarak “mkdir klasörüm”
	* olarak girseydik, komutu tek kelime olarak algılayacaktı. Yani string dizisi
	* mantığında çalışıyor bu olay. Sonuç olarak yukarıdaki gibi basit bir şekilde
	* komut satırına komut yollayabilirsiniz
	 */
	cmd := exec.Command("mkdir", "my_folder")
	cmd.Stdout = os.Stdout

	cmd.Run()

	// Örnek 2: Komut Satırına Komut Gönderip Çıktısını Okuma
	/*
	* Yukarıda çok kolay bir şekilde komut göndermeyi gördük. Fakat iş
	* komutun çıktısını okumaya gelince işler biraz karışıyor. Yavaştan
	* vaziyetinizi alın
	*
	* Aslında korkulacak bir olay yok. Yeter ki mantığını anlayalım. Şimdi
	* yapacağımız işlemleri 4 ana parçaya bölelim.
	*
	* 1) Komutun tanımlanması
	* 2) Çıktı okuyucusunun tanımlanması
	* 3) Komutun başlatılması
	* 4) Komutun çalışması
	 */

	//komutun tanımlanması
	gorVersionCmd := exec.Command("go", "version")
	// gönderdiğimiz komutun çıktılarını alabiliyoruz.
	cmdReader, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Fprintln(os.Stderr, "cmd reader error: ", err)
		/*
		* hata değişkeninin içi boş değilse ekrana bastırmasını ve 1 numaralı
		* çıkış kodunu vermesini istedik. Bu arada 1 numaralı çıkış kodu hatalar
		* için kullanılır. Golang programlarında görmüyoruz ama 0 numaralı çıkış
		* kod da işler yolunda gittiği zaman kullanılır. C dili kodlayan
		* arkadaşlarımız bilir, int main fonksiyonunun sonuna return 0 ibaresi
		* girilir. Buraya kadar olan işlemlerimiz komutun tanımlanması ile ilgiliydi.
		 */
		os.Exit(1)
	}

	//çıktı okuyucusunun tanımlanması
	cmdOut := bufio.NewScanner(cmdReader)
	go func() {
		for cmdOut.Scan() {
			fmt.Println(cmdOut.Text())
		}
	}()

	//komutun başlatılması
	/*
	* hata değişkenimize cmd.Start() fonksiyonunu atayarak komut başlatma işleminde
	* hata oluşursa veriyi çekmesini sağladık. Hata var ise error tipindeki hata
	* mesajımızı ekrana ve 1 numaralı hatayı ekrana bastıracak
	 */
	err = gorVersionCmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Komut başlatılamadı:", err)
		os.Exit(1)
	}

	//komutun çalışması
	/*
	* Son işlemimiz ise komutun sonuçlanmasının beklenmesi. hata değişkenimize
	* `cmd.Wait()` fonksiyonunu ekleyerek bekleme işleminde oluşabilecek hatanın
	* mesajını çekmiş olduk. Aşağısında eğer hata var ise ekrana bastırması için
	* gerekli kodlarımızı girdik. Son olarak 1 numaralı çıkış işlemini yaptık.
	 */
	err = gorVersionCmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Komut çalışırken hata oluştu:", err)
		os.Exit(1)
	}

}

func main() {
	// Örnek 3: Hata Detayı Çekmeden Komut Çıktısı Alma

	goVCmd := exec.Command("go", "version")
	outPut, error := goVCmd.CombinedOutput()
	if error != nil {
		log.Fatalf("Komut hatası: %s\n", error)
	}

	fmt.Println("go version output: ", string(outPut))
}
