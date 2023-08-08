package main

import (
	"fmt"
	"time"
)

/// Switch…Case Statements
/*
* Bu eğitimde Golang'da farklı koşullara göre farklı eylemler gerçekleştirmek
* için switch-case deyimini nasıl kullanacağınızı öğreneceksiniz.
*
* Golang ayrıca Php veya Java gibi diğer dillerde bulunana benzer
* bir switch deyimini de destekler. Switch deyimleri, bir değişkenin
* durumuna bağlı olarak uzun if else karşılaştırmalarını daha okunabilir
* kodlarla ifade etmenin alternatif bir yoludur.
 */

// switch Statement
func main1() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// switch multiple cases Statement
func main2() {
	today := time.Now()
	var t int = today.Day()

	switch t {
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// switch fallthrough(devam et) case Statement
func main3() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// swith conditional cases Statement

func main4() {
	today := time.Now()

	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// switch initializer Statement
func main5() {
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// run all code
func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
}
