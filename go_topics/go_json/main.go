package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/// Working with json
/*
* Go dilinde JSON (JavaScript Object Notation) verileriyle çalışmak oldukça
* yaygın bir gerekliliktir. JSON, verileri insanlar ve programlar arasında
* kolayca paylaşılabilir ve okunabilir bir formatta temsil etmek için kullanılır.
* Go dilinde JSON verileri ile çalışmak için encoding/json paketi kullanılır.
*
* JSON verileri, Go dilindeki veri türlerine dönüştürülerek veya Go veri
* türlerinden JSON formatına dönüştürülerek işlenebilir. Bu işlem "marshalling"
* ve "unmarshalling" olarak adlandırılır.
 */

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

func main() {
	jsondata := `{"id": 7, "name": "Eldiiar", "lastname": "Almazbek", "age": 22}`
	var eldi Person

	eldiError := json.Unmarshal([]byte(jsondata), &eldi)

	if eldiError != nil {
		log.Fatal("Eror Unmarshal", eldiError)
	}

	fmt.Println(eldi)
	fmt.Println(eldi.Id)
	fmt.Println(eldi.Name)
	fmt.Println(eldi.Lastname)
	fmt.Println(eldi.Age)

	ruslan := Person{Id: 8, Name: "Ruslan", Lastname: "Asanov", Age: 22}

	ruslanJson, errorRuslan := json.Marshal(ruslan)

	if errorRuslan != nil {
		log.Fatal("Eror Marshal", errorRuslan)
	}

	fmt.Println(ruslanJson)
	fmt.Println(string(ruslanJson))
}
