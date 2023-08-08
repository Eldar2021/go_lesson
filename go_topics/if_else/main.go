package main

import "fmt"

/// If Else
/*
The if statement - executes some code if one condition is true
The if...else statement - executes some code if a condition is true and another code if that condition is false
The if...else if....else statement - executes different codes for more than two conditions
The switch...case statement - selects one of many blocks of code to be executed
*/

// if Statement
func main1() {
	x := true
	if x {
		fmt.Println("Japan")
	}
}

// if...else Statement
func main2() {
	x := 100

	if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}

// if...else if...else Statement
func main3() {
	x := 100

	if x == 50 {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}

// if statement initialization
func main4() {
	if x := 100; x == 100 {
		fmt.Println("Japan")
	}
}

// run all code
func main() {
	main1()
	main2()
	main3()
	main4()
}
