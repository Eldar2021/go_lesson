package main

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

/// `.ini` Dosyası Okuma ve Düzenleme
/*
* `.ini` dosyası, yapılandırma verilerini depolamak ve okumak için kullanılan
* bir dosya formatıdır. Genellikle ayarlar, yapılandırmalar, parametreler
* gibi verilerin saklandığı metin dosyalarıdır. Go projelerinde `.ini` dosyaları,
* uygulamanın davranışını veya özelliklerini yapılandırmak için kullanılabilir.
*
* `.ini` dosyası kullanmanın bazı avantajları şunlar olabilir:
*
* 1) Kolay Okunabilirlik: `.ini` dosyaları genellikle basit bir yapıya sahiptir
*    ve insanlar tarafından kolayca okunabilir.
* 2) Dinamik Yapılandırma: Uygulamanın çalışma zamanında yapılandırma
*    ayarlarını değiştirmenizi sağlar, bu da uygulamayı farklı senaryolara
*    uygun hale getirir.
* 3) Dışarıdan Düzenlenebilir: `.ini` dosyalarını metin düzenleyicilerle açıp
*    düzenleyebilirsiniz, bu da ayarları hızlı bir şekilde değiştirmenizi sağlar.
*
* Go dilinde .ini dosyası oluşturmak ve okumak için çeşitli kütüphaneler bulunur.
* Örnek olarak, `go-ini/ini` ve `go-ini/ini/v1` kütüphaneleri kullanılabilir.
* Aşağıda örnek bir `.ini` dosyasının oluşturulması ve okunması için kullanılan
* `go-ini/ini` kütüphanesinin kullanımı gösterilmiştir:
 */

func CreateIniFile(fileName string) {
	config := ini.Empty()
	fmt.Println(config)

	// Bölüm ve ayarları ekliyoruz
	section, err := config.NewSection("app")

	if err != nil {
		log.Fatal(err)
	}

	section.NewKey("name", "MyApp")
	section.NewKey("version", "1.0")
	section.NewKey("debug", "true")

	// Dosyaya kaydediyoruz
	err = config.SaveTo(fileName + ".ini")
	if err != nil {
		log.Fatal(err)
	}
}

func ReadIniFile(fileName string) {
	config, err := ini.Load(fileName + ".ini")

	if err != nil {
		log.Fatal(err)
	}

	section := config.Section("app")
	fmt.Println("Name:", section.Key("name").String())
	fmt.Println("Version:", section.Key("version").String())
	fmt.Println("Debug:", section.Key("debug").MustBool(false))
}

func main() {
	fileName := "config"

	// CreateIniFile(fileName)

	ReadIniFile(fileName)
}
