package main

import (
	"fmt"
	"reflect"
)

/// Arrays
/*
* Dizi, tek bir türdeki öğelerin bir koleksiyonundan oluşan bir veri
* yapısıdır veya basitçe, aynı anda birden fazla değeri tutabilen özel
* bir değişken diyebilirsiniz. Bir dizinin tuttuğu değerlere elemanları
* veya öğeleri denir. Bir dizi belirli sayıda öğe tutar ve büyüyüp küçülemez.
* Int, String, Boolean ve diğerleri gibi farklı veri türleri dizilerde öğe
* olarak ele alınabilir. Bir dizinin herhangi bir boyutunun ilk elemanının
* indisi 0, herhangi bir dizi boyutunun ikinci elemanının indisi 1'dir ve
* bu böyle devam eder
 */

func main1() {
	var intArray [5]int
	var strArray [5]string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
}

// Assign and Access Values

func main2() {
	var theArray [3]string
	theArray[0] = "India"  // Assign a value to the first element
	theArray[1] = "Canada" // Assign a value to the second element
	theArray[2] = "Japan"  // Assign a value to the third element

	fmt.Println(theArray[0]) // Access the first element value
	fmt.Println(theArray[1]) // Access the second element valu
	fmt.Println(theArray[2]) // Access the third element valu
}

// Initializing an Array with an Array Literal
func main3() {
	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(x)
	fmt.Println(y)
}

// Initializing an Array with ellipses...
/*
* Uzunluğu belirtmek yerine ... kullandığımızda. Derleyici,
* dizi bildiriminde belirtilen elemanlara dayanarak bir dizinin uzunluğunu
* belirleyebilir.
 */

func main4() {
	x := [...]int{10, 20, 30}

	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(len(x)) // 3
}

// Initializing Values for Specific Elements
func main5() {
	x := [5]int{1: 10, 3: 30}
	fmt.Println(x) // [0 10 0 30 0]
}

// Loop Through an Indexed Array
/*
* Bir for döngüsü kullanarak bir dizi elemanları arasında döngü oluşturabilirsiniz.
 */
func main6() {
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("---------------Example 1--------------------")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("---------------Example 2--------------------")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	fmt.Println("---------------Example 3--------------------")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("---------------Example 4--------------------")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}

// Copy Array
/*
* Bir diziyi değer ya da referans olarak yeni bir değişkene atayarak
* dizinin kopyasını oluşturabilirsiniz.
 */
func main7() {

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by refrence

	fmt.Printf("strArray1: %v\n", strArray1) // [Japan Australia Germany]
	fmt.Printf("strArray2: %v\n", strArray2) // [Japan Australia Germany]

	strArray1[0] = "Canada"

	fmt.Printf("strArray1: %v\n", strArray1)   // [Canada Australia Germany]
	fmt.Printf("strArray2: %v\n", strArray2)   // [Japan Australia Germany]
	fmt.Printf("*strArray3: %v\n", *strArray3) // [Canada Australia Germany]
}

// Check if Element Exists
func main8() {
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strArray, "Canada"))
	fmt.Println(itemExists(strArray, "Africa"))
}

func itemExists(arrayType interface{}, item interface{}) bool {
	/*
	* reflect.ValueOf işlevi ile arrayType'ın yansımasını elde ediyorsunuz.
	* Bu, yansıma üzerinden dizinin elemanlarına erişmeyi sağlar.
	 */
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

// Filter Array Elements
/*
* You can filter array element using : as shown below
 */

func main9() {
	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[len(countries)-1])

	fmt.Printf("All elements: %v\n", countries[0:])
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:])

	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:])
}

// run all code
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
}
