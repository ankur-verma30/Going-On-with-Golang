package main

import "fmt"

func main(){
	var a,b,c int
	a=10
	b=20
	c=a+b

	fmt.Println(c) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 3 (integer division)
	fmt.Println(a % b) // 1

	// Relational Operators
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b) // false
	fmt.Println(a < b) // true

	// Logical Operators
	age :=20
	isAdult := age > 18
	fmt.Println(isAdult) // false
	fmt.Println(age > 18 && age < 30) // true
	fmt.Println(age > 18 || age < 30) // true

	// Bitwise Operators
	x := 5
	y := 3
	fmt.Println(x & y) // 1
	fmt.Println(x | y) // 7
	fmt.Println(x ^ y) // 6
	fmt.Println(x &^ y) // 4

	// Shift Operators
	fmt.Println(x << 1) // 10
	fmt.Println(x >> 1) // 2


	// Unary Operators
	fmt.Println(+x) // 5
	fmt.Println(-x) // -5


	// Assignment Operators
	var z int
	z = 10
	z += 5 // z = z + 5
	fmt.Println(z) // 15


	// Pointer operators
	t :=10
	p := &t
	fmt.Println(*p) // 10
	*p = 20
	fmt.Println(t) // 20
}