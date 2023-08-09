package main

import (
	"fmt"
	"reflect"
)

/// Slices/Dilim
/*
* Dilim, veri koleksiyonlarını uygulamak ve yönetmek için esnek ve
* genişletilebilir bir veri yapısıdır. Dilimler, hepsi aynı türde olan
* birden fazla öğeden oluşur. Dilim, uygun gördüğünüz şekilde büyüyebilen
* ve küçülebilen dinamik dizilerin bir parçasıdır. Diziler gibi dilimler
* de indekslenebilir ve bir uzunluğa sahiptir. Dilimlerin kapasite ve
* uzunluk özelliği vardır.
 */

func main1() {
	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}

// Declare Slice using Make
func main2() {
	var intSlice = make([]int, 10)        // when length and capacity is same
	var strSlice = make([]string, 10, 20) // when length and capacity is different

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}

// Initialize Slice with values using a Slice Literal
func main3() {
	var intSlice = []int{10, 20, 30, 40}
	var strSlice = []string{"India", "Canada", "Japan"}

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
}

// Declare Slice using new Keyword
func main4() {
	var intSlice = new([50]int)[0:10]

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
}

// Add Items
func main5() {
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	a = append(a, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println("Slice A after appending data:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
}

// Access Items
func main6() {
	var intSlice = []int{10, 20, 30, 40}

	fmt.Println(intSlice[0])
	fmt.Println(intSlice[1])
	fmt.Println(intSlice[0:4])
}

// Change Item Value
func main7() {
	var strSlice = []string{"India", "Canada", "Japan"}
	fmt.Println(strSlice)

	strSlice[2] = "Germany"
	fmt.Println(strSlice)
}

// Remove Item from Slice
func main8() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(strSlice)

	strSlice = RemoveIndex(strSlice, 3)
	fmt.Println(strSlice)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// Copy a Slice
func main9() {
	a := []int{5, 6, 7} // Create a smaller slice
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))

	b := make([]int, 5, 10) // Create a bigger slice
	copy(b, a)              // Copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

	fmt.Println("Slice B after copying:", b)
	b[3] = 8
	b[4] = 9
	fmt.Println("Slice B after adding elements:", b)
}

// Tricks of Slicing
func main10() {
	var countries = []string{"india", "japan", "canada", "australia", "russia"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[4])
	fmt.Printf("Last element: %v\n", countries[len(countries)-1])
	fmt.Printf("Last element: %v\n", countries[4:])

	fmt.Printf("All elements: %v\n", countries[0:])

	fmt.Printf("Last two elements: %v\n", countries[3:])
	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:])

	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:])
}

// Loop Through a Slice
func main11() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Println("---------------Example 1 --------------------")
	for index, element := range strSlice {
		fmt.Println(index, "--", element)
	}

	fmt.Println("---------------Example 2 --------------------")
	for _, value := range strSlice {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("---------------Example 3 --------------------")
	for range strSlice {
		fmt.Println(strSlice[j])
		j++
	}
}

// Append a slice to an existing slice
func main12() {
	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}

	slice2 = append(slice2, slice1...)

	fmt.Println(slice2)
}

// Check if Item Exists

func main13() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strSlice, "Canada"))
	fmt.Println(itemExists(strSlice, "Africa"))
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
	main7()
	main8()
	main9()
	main10()
	main11()
	main12()
	main13()
}
