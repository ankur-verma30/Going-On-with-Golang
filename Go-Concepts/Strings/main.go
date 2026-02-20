package main

import "fmt"

func main(){
	fmt.Println("Hello World!")
	name:= "Go Concepts" // stored as sequence of bytes
	fmt.Println("Welcome to", name)
	firsstChar := name[0] // accessing first character of the string
	fmt.Println("First character of the string is:", firsstChar) // returns the ASCII value of the character
	fmt.Println("First character of the string is:", string(firsstChar)) // converting ASCII value back to character
	fmt.Println("Length of the string is:", len(name)) // length of the string in bytes


	lastChar := name[len(name)-1] // accessing last character of the string
	fmt.Println("Last character of the string is:", lastChar)
	fmt.Println("Last character of the string is:", string(lastChar)) // converting ASCII value back to character


	// string are immutable in Go, we cannot change the value of a string once it is created
	// name[0] = 'g' // ❌ this will give an error

	// to change the value of a string, we need to create a new string
	newName := "g" + name[1:] // creating a new string by concatenating 'g' with the substring of name starting from index 1
	fmt.Println("New string is:", newName)


	//converting byte or rune to the string
	byteValueOfA := byte(65) // ASCII value of 'A'
	runeValueOfB := rune(66) // ASCII value of 'B'
	fmt.Println("Byte value as string:", string(byteValueOfA)) // converting byte to string
	fmt.Println("Rune value as string:", string(runeValueOfB)) // converting rune to string

	//single quotes is used for byte and rune literals, while double quotes and backtick is used for string literals

	//convert string to byte slice
	byteSlice := []byte(name) // converting string to byte slice
	fmt.Println("Byte slice:", byteSlice)

	//convert string to rune slice
	runeSlice := []rune(name) // converting string to rune slice
	fmt.Println("Rune slice:", runeSlice)

	// Slice of string
	sliceOfString := name[0:2] // slicing the string from index 0 to 2 (exclusive)
	fmt.Println("Slice of string:", sliceOfString)

	// String concatenation
	fullName := name + " - A Comprehensive Guide" // concatenating two strings
	fmt.Println("Full name:", fullName)

	// String formatting
	formattedString := fmt.Sprintf("Welcome to %s!", name) // formatting a string using Sprintf
	fmt.Println(formattedString)

	//accessing unicode characters in a string
	unicodeString := "Hello, 世界" // string containing unicode characters
	fmt.Println("Unicode string:", unicodeString)
	fmt.Println("Length of unicode string in bytes:", len(unicodeString)) // length in bytes
	fmt.Println("Length of unicode string in runes:", len([]rune(unicodeString))) // length in runes
	
	runeSliceOfUnicodeString := []rune(unicodeString) // converting unicode string to rune slice
	fmt.Println("Rune slice of unicode string:", string(runeSliceOfUnicodeString[7]));

	// iterating over a string using range
	fmt.Println("Iterating over the string:")
	for index, char := range name {
		fmt.Printf("Index: %d, Character: %c\n", index, char) // char is of type rune
	}
}