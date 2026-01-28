package main

import "fmt"

const x=10 // global constant with untyped constant
const y int32= 15

const(
	pi=3.14
	e=2.718
	character='A'
	isRunning=false
	isTrue=1<2
)
func main(){
	fmt.Println("-----------Working with constants---------")
	var a int
	a=x
	fmt.Println("The value of a is ",a)

	var b float64
	b=x
	fmt.Println("The value of b is ",b)

	var c int
	c=int(y)
	fmt.Println("The value of int is ",c)
	fmt.Printf("The value of the pi and e is %f and %f", pi,e)

}

func sum(a int,b int) int{
	return a+b
}