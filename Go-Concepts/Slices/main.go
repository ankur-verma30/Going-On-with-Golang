package main

import "fmt"

func main() {
	var a []int //nil slice
	fmt.Println(a == nil)

	a = []int{} //empty slice
	fmt.Println(a == nil)

	//different ways of declaring and initializing slice

	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := []int{5, 2: 10, 50}
	fmt.Println(c)

	//make function

	d := make([]int, 5)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	e := make([]int, 5, 10)
	fmt.Println(e)
	fmt.Println(len(e))
	fmt.Println(cap(e))

	//append function
	f := make([]int, 0)
	f = append(f, 1)
	fmt.Println(f)
	f = append(f, 2)
	g := append(f, 3, 4, 5)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))
	g = append(g, f...)
	fmt.Println(g)

}
