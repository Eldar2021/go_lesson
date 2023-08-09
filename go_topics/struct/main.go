package main

import (
	"fmt"
	"reflect"
)

// / Struct
type rectangle1 struct {
	length  float64
	breadth float64
	color   string
}

func main1() {
	fmt.Println(rectangle1{10.5, 25.10, "red"})
}

// Creating Instances of Struct Types
type rectangle2 struct {
	length  int
	breadth int
	color   string

	geometry struct {
		area      int
		perimeter int
	}
}

func main2() {
	var rect rectangle2 // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
}

type rectangle3 struct {
	length  int
	breadth int
	color   string
}

// Creating a Struct Instance Using a Struct Literal
func main3() {
	var rect1 = rectangle3{10, 20, "Green"}
	fmt.Println(rect1)

	var rect2 = rectangle3{length: 10, color: "Green"} // breadth value skipped
	fmt.Println(rect2)

	rect3 := rectangle3{10, 20, "Green"}
	fmt.Println(rect3)

	rect4 := rectangle3{length: 10, breadth: 20, color: "Green"}
	fmt.Println(rect4)

	rect5 := rectangle3{breadth: 20, color: "Green"} // length value skipped
	fmt.Println(rect5)
}

// Struct Instantiation Using new Keyword
type rectangle4 struct {
	length  int
	breadth int
	color   string
}

func main4() {
	rect1 := new(rectangle4) // rect1 is a pointer to an instance of rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)

	var rect2 = new(rectangle4) // rect2 is an instance of rectangle
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)
}

// Struct Instantiation Using Pointer Address Operator
type rectangle5 struct {
	length  int
	breadth int
	color   string
}

func main5() {
	var rect1 = &rectangle5{10, 20, "Green"} // Can't skip any value
	fmt.Println(rect1)

	var rect2 = &rectangle5{}
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2) // breadth skipped

	var rect3 = &rectangle5{}
	(*rect3).breadth = 10
	(*rect3).color = "Blue"
	fmt.Println(rect3) // length skipped
}

// Nested Struct Type
type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func main6() {
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])
}

// Assign Default Value for Struct Field
type Employee6 struct {
	Name string
	Age  int
}

func (obj *Employee6) Info() {
	if obj.Name == "" {
		obj.Name = "John Doe"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}

func main7() {
	emp1 := Employee6{Name: "Mr. Fred"}
	emp1.Info()
	fmt.Println(emp1)

	emp2 := Employee6{Age: 26}
	emp2.Info()
	fmt.Println(emp2)
}

// Find Type of Struct
type rectangle7 struct {
	length  float64
	breadth float64
	color   string
}

func main8() {
	var rect1 = rectangle7{10, 20, "Green"}
	fmt.Println(reflect.TypeOf(rect1))         // main.rectangle
	fmt.Println(reflect.ValueOf(rect1).Kind()) // struct

	rect2 := rectangle7{length: 10, breadth: 20, color: "Green"}
	fmt.Println(reflect.TypeOf(rect2))         // main.rectangle
	fmt.Println(reflect.ValueOf(rect2).Kind()) // struct

	rect3 := new(rectangle7)
	fmt.Println(reflect.TypeOf(rect3))         // *main.rectangle
	fmt.Println(reflect.ValueOf(rect3).Kind()) // ptr

	var rect4 = &rectangle7{}
	fmt.Println(reflect.TypeOf(rect4))         // *main.rectangle
	fmt.Println(reflect.ValueOf(rect4).Kind()) // ptr
}

// Comparing Structs with the Different Values Assigned to Data Fields
type rectangle9 struct {
	length  float64
	breadth float64
	color   string
}

func main9() {
	var rect1 = rectangle9{10, 20, "Green"}
	rect2 := rectangle9{length: 20, breadth: 10, color: "Red"}

	if rect1 == rect2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	rect3 := new(rectangle9)
	var rect4 = &rectangle9{}

	if rect3 == rect4 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

// Copy Struct Type Using Value and Pointer Reference
type rectangle10 struct {
	length  float64
	breadth float64
	color   string
}

func main10() {
	r1 := rectangle10{10, 20, "Green"}
	fmt.Println(r1)

	r2 := r1
	r2.color = "Pink"
	fmt.Println(r2)

	r3 := &r1
	r3.color = "Red"
	fmt.Println(r3)

	fmt.Println(r1)
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
}
