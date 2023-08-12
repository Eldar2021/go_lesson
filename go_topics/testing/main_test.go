package main

import "testing"

func TestMerhaba(t *testing.T) {
	if sayHello("Eldiiar", "Almazbek") != "Eldiiar Almazbek" {
		t.Error("Merhaba Fonksiyonunda bir sıkıntı var!")
	}
}

/*
* `go test` ------> İçerisinde bulunduğu projenin tüm test fonksiyonlarını test eder.
* `go test -v` ----> Her test için ayrı bilgi verir.
* `go test -timeout 30s` ----> 30 saniye zaman aşımı ile test eder.
* `go test -run TestMerhaba` ----> Sadece belirli bir fonksiyonu test eder.
 */
