package main

import "fmt"

func main(){
	fmt.Println("------------------Using Strings Data Type-------------------")
	 
	var firstString string="Hello world!"
	fmt.Println("Value of firstString is:",firstString)

	//for multi-line strings we can use \n escape character
	var multiLineStringWithEscape string="This is a multi-line string.\nIt preserves the formatting,\n\tincluding line breaks and spaces."
	fmt.Println("Value of multi line string with escape is:",multiLineStringWithEscape)

	// for multi-line strings we use backticks `` also called raw string literals
	var multiLineString string=`This is a multi-line string.
It preserves the formatting,
	including line breaks and spaces.`
	fmt.Println("Value of multiLineString is:",multiLineString)

	//concatenation of strings
	var string1 string="Hello, "
	var string2 string="Gophers!"
	var concatenatedString string=string1 + string2
	fmt.Println("Value of concatenatedString is:",concatenatedString)

	//string formatting
	name:="Alice"
	age:=30
	formattedString:=fmt.Sprintf("My name is %s and I am %d years old.",name,age)
	fmt.Println("Value of formattedString is:",formattedString)
}