package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "résumé"
	// Indexing gives the BYTE value (114 for 'r') because we are acutally indexing the utf8 byte array for each char
	print(s[0])
	print(s[1]) // Will give 195 because é has 2 bytes but this prints the first one only
	print(s[2])
	// 2 methods to see the char rather than the number
	print(string(s[0]))
	fmt.Printf("%c\n", s[0])

	print(len(s)) // gives the no of bytes (2 for é in utf8 encoding)

	// Casting to rune slice gives the actual characters
	r := []rune(s)
	print(r[1]) // Prints 233 (the code for 'é')

	// Efficient string building
	var sb strings.Builder
	sb.WriteString("Hello ")
	sb.WriteString("World")
	fmt.Println(sb.String())

	// Go's for range loop is actually "rune-aware." When you range over a string, Go automatically decodes the bytes into runes for you
	p := "résumé"
	for index, char := range p {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// Double quotes ("a") create a string.
	// Single quotes ('a') create a rune.

	myString := "A" // This is a string (slice of bytes)
	myRune := 'A'   // This is a rune (an int32 with the value 65)
	fmt.Printf("%T\n", myString)
	fmt.Printf("%T", myRune) // Output: int32

}

func print(v any) {
	fmt.Println(v)
}

/*
In Go, a string is a read-only slice of bytes. Because Go uses UTF-8 encoding, a single character (like an emoji or an accented letter é) might take up 1 to 4 bytes.
Bytes: The raw data. len("string") returns the number of bytes, not characters.
Runes: An alias for int32. It represents a single Unicode "Code Point."
Immutability: Strings cannot be changed once created. To build strings efficiently, use strings.Builder.
*/

/*

 */
