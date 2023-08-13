package main

import (
	"fmt"
	"sort"
)

/// Sort (Sıralama)
/*
* Golang’ta dizilerin içeriğini sıralaya bileceğimiz bütünleşik olarak gelen
* “sort” isminde bir paket mevcuttur. Bu paketin kullanımı oldukça kolaydır.
 */

func main() {
	names := []string{"eldiiar", "ruslan", "baktiyar"}
	fmt.Println(names)
	ages := []int{23, 24, 27}

	sort.Strings(names)
	fmt.Println(names)

	sort.Ints(ages)
	fmt.Println(ages)

	namesIsSorted := sort.StringsAreSorted(names)
	fmt.Println("names is sorted: ", namesIsSorted)

	agesIsSorted := sort.IntsAreSorted(ages)
	fmt.Println("ages is sorted", agesIsSorted)
}
