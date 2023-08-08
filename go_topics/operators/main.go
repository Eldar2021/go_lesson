package main

import "fmt"

/// Operators https://www.golangprograms.com/golang-if-else-statements.html

// Arithmetic Operators
/*
* Operator	Description	     Example	Result
* +	        Addition	     x + y	    Sum of x and y
* -	        Subtraction	     x - y	    Subtracts one value from another
* *	        Multiplication	 x * y	    Multiplies two values
* /	        Division	     x / y	    Quotient of x and y
* %	        Modulus	         x % y	    Remainder of x divided by y
* ++	    Increment	     x++	    Increases the value of a variable by 1
* --	    Decrement	     x--	    Decreases the value of a variable by 1
 */

func main1() {
	var x, y = 35, 7

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)
}

// Assignment Operators
/*
* Assignment	Description	                  Example
* x = y	        Assign	                      x = y
* x += y	    Add and assign	              x = x + y
* x -= y	    Subtract and assign	          x = x - y
* x *= y	    Multiply and assign	          x = x * y
* x /= y	    Divide and assign quotient	  x = x / y
* x %= y	    Divide and assign modulus	  x = x % y
 */
func main2() {
	var x, y = 15, 25
	x = y
	fmt.Println("= ", x)

	x = 15
	x += y
	fmt.Println("+=", x)

	x = 50
	x -= y
	fmt.Println("-=", x)

	x = 2
	x *= y
	fmt.Println("*=", x)

	x = 100
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)
}

// Comparison Operators
/*
* Operator	 Name	                   Example	  Result
* ==	     Equal	                   x == y	  True if x is equal to y
* !=	     Not equal	               x != y	  True if x is not equal to y
* <	         Less than	               x < y	  True if x is less than y
* <=	     Less than or equal to	   x <= y	  True if x is less than or equal to y
* >	         Greater than	           x > y	  True if x is greater than y
* >=	     Greater than or equal to  x >= y	  True if x is greater than or equal to y
 */

func main3() {
	var x, y = 15, 25

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
}

// Logical Operators
/*
* Operator	 Name	        Description	                                              Example
* &&	     Logical And	Returns true if both statements are true	              x < y && x > z
* ||	     Logical Or	    Returns true if one of the statements is true	          x < y || x > z
* !	         Logical Not	Reverse the result, returns false if the result is true	  !(x == y && x > z)
 */
func main4() {
	var x, y, z = 10, 20, 30

	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x == y && x > z))
}

// Bitwise Operators
/*
* Operator	Name	               Description
* &	        AND	                   Sets each bit to 1 if both bits are 1
* |	        OR	                   Sets each bit to 1 if one of two bits is 1
* ^	        XOR	                   Sets each bit to 1 if only one of two bits is 1
* <<	    Zero fill left shift   Shift left by pushing zeros in from the right and let the leftmost bits fall off
* >>	    Signed right shift	   Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off
 */

func main5() {
	var x uint = 9  //0000 1001
	var y uint = 65 //0100 0001
	var z uint

	z = x & y
	fmt.Println("x & y  =", z)

	z = x | y
	fmt.Println("x | y  =", z)

	z = x ^ y
	fmt.Println("x ^ y  =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)

	z = x >> 1
	fmt.Println("x >> 1 =", z)
}

/*
AND (&): Her iki sayının ilgili bitleri 1 ise sonuç 1 olur, aksi durumda sonuç 0 olur.

OR (|): Her iki sayının ilgili bitlerinden en az biri 1 ise sonuç 1 olur, aksi durumda sonuç 0 olur.

XOR (^): Her iki sayının ilgili bitleri farklı ise sonuç 1 olur, aynı ise sonuç 0 olur.

NOT (~): Verinin bütün bitlerini tersine çevirir (1'leri 0 yapar, 0'ları 1 yapar).

Shift Left (<<): Belirli sayıda biti sola kaydırır.

Shift Right (>>): Belirli sayıda biti sağa kaydırır.
*/
func main6() {
	a := 5 // 101
	b := 3 // 011

	// Bitwise AND
	resultAnd := a & b // 001
	fmt.Printf("AND Sonucu: %d\n", resultAnd)

	// Bitwise OR
	resultOr := a | b // 111
	fmt.Printf("OR Sonucu: %d\n", resultOr)

	// Bitwise XOR
	resultXor := a ^ b // 110
	fmt.Printf("XOR Sonucu: %d\n", resultXor)

	// Bitwise NOT
	resultNot := ^a // 11111111111111111111111111111010
	fmt.Printf("NOT Sonucu: %d\n", resultNot)

	// Bitwise Shift Left
	resultShiftLeft := a << 2 // 10100
	fmt.Printf("Shift Left Sonucu: %d\n", resultShiftLeft)

	// Bitwise Shift Right
	resultShiftRight := a >> 1 // 10
	fmt.Printf("Shift Right Sonucu: %d\n", resultShiftRight)
}

// run all code
func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
}
