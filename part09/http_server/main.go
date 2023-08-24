package main

import (
	"fmt"
	"net/http"
	"os"
)

/// net/http ile Web Server Oluşturma
/*
* Golang’ta web sunucusu oluşturma çok basit bir işlemdir.
* İlk örneğimizde localhost:5555 üzerinde çalışacak olan bir
*
* web sunucusu oluşturacağız.
 */

// func handler(w http.ResponseWriter, r *http.Response) {
// 	fmt.Fprintf(w, "Merhaba %s", r.Request.URL.Path[1:])
// }

func main() {
	/*
	* Tarayıcınız üzerinden localhost:5555‘e girdiğinizde sayfada sadece
	* Merhaba yazdığını göreceksiniz. Daha sonra adrese /eldi yazıp girdiğiniz
	* zaman yazının Merhaba eldi olarak değiştiğini göreceksiniz.
	 */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := loadFile("index.html")

		fmt.Fprintf(w, body, "Hi :)")
	})

	http.ListenAndServe(":5555", nil)
	fmt.Println("Web Sunucu")
}

func loadFile(fileName string) (string, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
