package main

import "fmt"

func main(){
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
}