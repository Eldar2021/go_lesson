package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

/// XML Parsing (Ayrıştırma)
/*
* Bu yazımıza Golang üzerinde XML dosyalarını işlemeyi öğreneceğiz.
* Bu işlemin yapabileceğimiz hali hazırda standart Golang paketleri
* ile gelen “encoding/xml” paketi vardır. Örneğimize geçelim.
*
* data.xml isminde aşağıdaki gibi bir belgemiz olduğunu varsayalım.
 */

type ModelData struct {
	Alan    xml.Name `xml:"persons"`
	Persons []Person `xml:"person"`
}

type Person struct {
	Alan    xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Socials []Social `xml:"social"`
}

type Social struct {
	Alan     xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	finalName := "data.xml"

	// XML dosyamızı okuyoruz (byte olarak geliyor)
	data, err := os.ReadFile(finalName)

	if err != nil {
		fmt.Println("ReadFile error: ", err)
	}

	// Yerleştirme işlemi için değişken oluşturuyoruz.
	var modelData ModelData

	err = xml.Unmarshal(data, &modelData)

	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}

	fmt.Println("Data success: ", modelData)
	fmt.Println(modelData.Alan)
	fmt.Println(modelData.Persons)

	for _, p := range modelData.Persons {
		fmt.Println(p.Alan)
		fmt.Println(p.Name)
		for _, s := range p.Socials {
			fmt.Println(s.Alan)
			fmt.Println(s.Facebook)
			fmt.Println(s.Twitter)
			fmt.Println(s.Youtube)
		}
	}
}
