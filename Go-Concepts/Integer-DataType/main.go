package main

import "fmt"

func main(){
	fmt.Println("------------------Using Integer Data Type-------------------")

	var smallPositiveValue uint8
	smallPositiveValue=10
	fmt.Println(smallPositiveValue)
	
	//creating a variable of type int8
	var smallPositiveNegativeValue int8
	smallPositiveNegativeValue=-20
	fmt.Println(smallPositiveNegativeValue)

	var newInt int=2345543
	fmt.Println("Value of newInt is ",newInt)

	//conversion from int8 to int because int and int8 are treated as different data types
	newInt=int(smallPositiveNegativeValue)
	fmt.Println("Modified value of newInt is",newInt)

	// TypeCasting rule types case the right side value into the left side value type using the syntax
	// <variable_name>=<target_data_type>(<value_to_be_converted>)

	var myByte byte='A' // byte is an alias for uint8
	fmt.Println("Value of myByte is",myByte)

	var myRune rune='😁' // rune is an alias for int32 for unicode characters
	fmt.Println("Value of myRune is",myRune)

}