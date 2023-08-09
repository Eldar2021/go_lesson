package main

import (
	"fmt"
	"sort"
)

/// Maps

func main1() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee)
}

// Empty Map declaration
func main2() {
	var employee = map[string]int{}
	fmt.Println(employee)        // map[]
	fmt.Printf("%T\n", employee) // map[string]int
}

// Map declaration using make function
func main3() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)
}

// Map Length
func main4() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20

	// Empty Map
	employeeList := make(map[string]int)

	fmt.Println(len(employee))     // 2
	fmt.Println(len(employeeList)) // 0
}

// Accessing Items
func main5() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}

	fmt.Println(employee["Mark"])
}

// Adding Items
func main6() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee) // Initial Map

	employee["Rocky"] = 30 // Add element
	employee["Josef"] = 40

	fmt.Println(employee)
}

// Update Values
func main7() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee) // Initial Map

	employee["Mark"] = 50 // Edit item
	fmt.Println(employee)
}

// Delete Items
func main8() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	employee["Rocky"] = 30
	employee["Josef"] = 40

	fmt.Println(employee)

	delete(employee, "Mark")
	fmt.Println(employee)
}

// Iterate over a Map
func main9() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}
	for key, element := range employee {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}

// Truncate Map

func main10() {
	var employee1 = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}
	var employee2 = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}

	fmt.Println("1) employee1 : ", employee1)
	fmt.Println("1) employee2 : ", employee2)

	// Method - I
	for k := range employee1 {
		delete(employee1, k)
	}

	// Method - II
	employee2 = make(map[string]int)

	fmt.Println("2) employee1 : ", employee1)
	fmt.Println("2) employee2 : ", employee2)
}

// Sort Map Keys
func main11() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	keys := make([]string, 0, len(unSortedMap))

	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}
}

// Sort Map Values
func main12() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	// Int slice to store values of map.
	values := make([]int, 0, len(unSortedMap))

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	// Sort slice values.
	sort.Ints(values)

	// Print values of sorted Slice.
	for _, v := range values {
		fmt.Println(v)
	}
}

// Merge Maps

func main13() {
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first)
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
	main10()
	main11()
	main12()
	main13()
}
