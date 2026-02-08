package main

import "fmt"

func main(){
	var a[5] int
	fmt.Println(a)

	var b = [5]int{1,2,3,4,5}
	fmt.Println(b)

	var c = [5]int{1,2}
	fmt.Println(c)

	d :=[5]int{5, 2:10, 50}
	fmt.Println(d)

	e:=[...]int{5,2,3,4,5}
	fmt.Println(e)

	fmt.Println(len(e))

	//Accessing array elements
	fmt.Println(a[0])

	//Modifying array elements
	a[0] = 10
	fmt.Println(a)

	//Bounds check
	fmt.Println(a[3])


	f:=[2][2]int{{1,2},{3,4}}
	fmt.Println(f)
	fmt.Println(len(f))
	
}