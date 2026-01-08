package main

import "fmt"

func main(){
	var firstString string="This is a string"
	fmt.Println(firstString)

	var myString string 
	fmt.Println(myString) // Generates the zero value of the type i.e. ""
	myString="Hello World"
	fmt.Println(myString)

	var myInt int // go gives an error if the variable is not used 
	fmt.Println(myInt)
	
	var myBool bool
	fmt.Println(myBool)
}