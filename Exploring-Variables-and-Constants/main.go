package main

import "fmt"
var myInt int=10


//grouping variable: Used to group similar variables
var(
	myString="Hello"
	myInt2=20
)
func main(){
	// variable in golang
	fmt.Println(myInt)
	doSomething()

	//group variables can be written inside the function as well -> rarely used
	var(
		myBool bool=true
		myFloat float64=10.5
	)

	fmt.Println("Outside the main method group variables ",myString)
	fmt.Println("Outside the main method group variables ",myInt2)
	fmt.Println("Inside the main method group variables ",myBool)
	fmt.Println("Inside the main method group variables ",myFloat)

	//using short hand notation ----------> only used inside the function
	a:=10
	b:=20
	fmt.Println("These are shorthand variables ",a,b)
	// Atmost one new variable shoule be declared on the left side for redeclaration of pre-existing variables
	c,a:=10,20
	fmt.Println("These are shorthand variables with one new variables and other pre-existing variables ",c,a)
}

func doSomething(){
	fmt.Println("Do something: ",myInt)
}