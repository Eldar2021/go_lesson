package main

import (
	"flag"
	"fmt"
)

/// Komut Satırı Bayrakları (Flags)
/*
* Sondaki `-h` bir flag(bayrak)’dir. Örnek bir program yazalım.
 */

// go run . -name=Eldi -age=21 -isregistered=false
// go run . -h
func main() {
	name := flag.String("name", "Eldiiar", "string type")
	age := flag.Int("age", 22, "int type")
	isRegistered := flag.Bool("isregistered", true, "bool type")

	flag.Parse()

	fmt.Println("name -> ", *name)
	fmt.Println("age -> ", *age)
	fmt.Println("bool -> ", *isRegistered)
}
