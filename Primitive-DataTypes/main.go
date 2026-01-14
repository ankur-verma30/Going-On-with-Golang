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


	fmt.Println("------------------Using Float and Complex Data Type-------------------")

	var smallFloat float32= 3.14
	fmt.Println("Value of smallFloat is",smallFloat)

	var defaultFloat float64= 3.14159265358979323846
	fmt.Println("Value of defaultFloat is",defaultFloat) // The overflow digits are truncated in the output

	var complexNumber complex64= 2+3i
	fmt.Println("Value of complexNumber is",complexNumber)

	var defaultComplexNumber complex128= 5+7i
	fmt.Println("Value of defaultComplexNumber is",defaultComplexNumber)

	//using complex number built-in functions
	var realPart float64=real(defaultComplexNumber)
	var imagPart float64=imag(defaultComplexNumber)
	fmt.Println("Real part of defaultComplexNumber is",realPart)
	fmt.Println("Imaginary part of defaultComplexNumber is",imagPart)

	var myComplex complex64
	myComplex=complex(4,5) // using complex function to create complex number
	fmt.Println("Value of myComplex using complex function is",myComplex)

	fmt.Println("------------------Using Strings Data Type-------------------")

	

}