package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/// Dinamik JSON Parsing Yöntemleri
/*
* Bazı durumlarda struct ile JSON parse etme uğraştırıcı olabiliyor.
* Özellikle json verisi olarak hangi tipte veri geleceğini bilmediğimiz durumlarda.
 */

func main() {
	fileName := "example.json"

	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("ReadFile Error: ", err)
	}

	var mapData []map[string]interface{}

	json.Unmarshal(data, &mapData)

	for i, data := range mapData {
		fmt.Printf("%d: data is %v\n", i, data)
	}
}
