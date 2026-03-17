package main

import "fmt"

func main() {
fmt.Println("Hello World!");
// nil map
var m map[string]int
fmt.Println("Nil map:", m) // prints: Nil map: map[]
// m["key"] = 100 // ❌ panic: assignment to entry in nil map


// empty map
emptyMap := make(map[string]int)
fmt.Println("Empty map:", emptyMap) // prints: Empty map: map[]
emptyMap["key"] = 100 // works fine
fmt.Println("Empty map after adding an entry:", emptyMap) // prints: Empty map after adding an entry: map[key:100]

// map with initial capacity
capacityMap := make(map[string]int, 100)
fmt.Println("Map with initial capacity:", capacityMap) // prints: Map with initial capacity: map[]


// map literal
literalMap := map[string]int{
	"apple":  10,
	"banana": 20,
}
fmt.Println("Map literal:", literalMap) // prints: Map literal: map[apple:10 banana:20]


// Reading from a nil map
var nilMap map[string]int
value, exists := nilMap["key"]
fmt.Println("Value from nil map:", value) // prints: Value from nil map: 0
fmt.Println("Does key exist in nil map?", exists) // prints: Does key exist in nil map? false


// Reading from an empty map
emptyMapValue, emptyMapExists := emptyMap["key"]
fmt.Println("Value from empty map:", emptyMapValue) // prints: Value from empty map: 100
fmt.Println("Does key exist in empty map?", emptyMapExists) // prints: Does key exist in empty map? true


// Reading from a map with initial capacity comma,ok idiom is used to check if the key exists in the map or not
capacityMapValue, capacityMapExists := capacityMap["key"]
fmt.Println("Value from map with initial capacity:", capacityMapValue) // prints: Value from map with initial capacity: 0
fmt.Println("Does key exist in map with initial capacity?", capacityMapExists) // prints: Does key exist in map with initial capacity? false

// Reading from a map literal
literalMapValue, literalMapExists := literalMap["apple"]
fmt.Println("Value from map literal:", literalMapValue) // prints: Value from map literal: 10
fmt.Println("Does key exist in map literal?", literalMapExists) // prints: Does key exist in map literal? true


// Writing to an empty map
emptyMap["newKey"] = 200
fmt.Println("Empty map after adding another entry:", emptyMap) // prints: Empty map after adding another entry: map[key:100 newKey:200]


// Writing to a map with initial capacity
capacityMap["newKey"] = 300
fmt.Println("Map with initial capacity after adding an entry:", capacityMap) // prints: Map with initial capacity after adding an entry: map[newKey:300]


// Writing to a map literal
literalMap["newKey"] = 400
fmt.Println("Map literal after adding an entry:", literalMap) // prints: Map literal after adding an entry: map[apple:10 banana:20 newKey:400]


//map key should be comparable, if we try to use a slice as a key in a map, it will give an error because slices are not comparable in Go
// invalidMap := map[[]int]string{ // ❌ this will give an error
// 	{1, 2, 3}: "value",
// }

//valid map is
validMap:= map[string][]int{
	"key": {1, 2, 3},
}
fmt.Println("Valid map with slice as value:", validMap)
}