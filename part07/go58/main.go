package main

import (
	"fmt"
	"os"
)

/// Çapraz Platform Dosya Yolları
/*
* Bir işletim sisteminde dosyanın veya dizinin yolunu belirtmek için
* taksim veya ters-taksim işaretleri kullanırız. Fakat yazağımız program
* çapraz-platformsa bu durumda ne yapmamız gerekir?
*
* Ya kendimiz bunun için bir fonksiyon oluşturacağız ya da kısa yoldan
* `os.PathSeperator`'ı kullanabiliriz.
 */

func main() {
	s := string(os.PathSeparator)

	path := "projects" + s + "modules"

	fmt.Println(path)
}

// projects\modulesr --> windows
// projects/modules --> mac and linux
