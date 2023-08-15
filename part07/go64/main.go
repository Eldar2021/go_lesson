package main

import (
	"encoding/json"
	"fmt"
)

/// JSON Parsing (Ayrıştırma)
/*
* Bugünkü yazımızda Golang ile JSON parse etmeye bakacağız. Hepimizin bildiği gibi
* günümüzde bir API (application programming interface) a veri göndermede ya da veri
* çekmede en sık kullanılan veri formatı JSON (javascript object notation) dur.
*
* Golang ile de kendi oluşturduğumuz verimizi (Golang struct) JSON’a dönüştürüp bir
* API’a request olarak gönderebilir ya da bir API’dan gelen JSON verisini Go programımızda
* kullanabiliriz.  O halde çok uzatmadan Go programımızdaki verileri nasıl JSON’a
* dönüştürüz hemen bakalım:
 */

type Admin struct {
	Name     string
	LastName string
	Age      int8
}

func main() {
	/// Resouces
	admin1 := Admin{Name: "Eldiiar", LastName: "Almazbek", Age: 22}
	jsonData := []byte(`{"Name": "Almazbek", "LastName": "Kaparov", "Age": 48}`)
	fileNane := "example.json"

	fmt.Println(fileNane)

	// Marhsalling (to json)
	adminToMarshal(admin1)
	adminToMarshalIndent(admin1)

	// Unmarshal (from json)
	var admin2 Admin
	adminFromJson(jsonData, &admin2)
	fmt.Println(admin2)
}

func adminToMarshal(admin Admin) ([]byte, error) {
	data, err := json.Marshal(admin)

	if err != nil {
		fmt.Println("Marshal Error: ", err)
		return nil, err
	}

	fmt.Println("Marshal Data: ", string(data))

	return data, nil
}

func adminToMarshalIndent(admin Admin) ([]byte, error) {
	data, err := json.MarshalIndent(admin, "", "    ")

	if err != nil {
		fmt.Println("MarshalIndent Error: ", err)
		return nil, err
	}

	fmt.Println("MarshalIndent Data: ", string(data))

	return data, nil
}

func adminFromJson(data []byte, admin *Admin) error {
	err := json.Unmarshal(data, admin)

	if err != nil {
		fmt.Println("Unmarshal Error: ", err)
		return nil
	}

	return nil
}


