package main

import "fmt"

/*
Arrays in GO are:
Fixed length
Same type
Indexable
Contiguous in memory - Because the elements sit right next to each other in memory, the CPU can access them incredibly fast
*/

func main() {
	fmt.Println("Start")
	// use ... to let the compiler count for you
	intarray := [...]int32{1, 2, 3}
	fmt.Println(intarray) //can print whole array

	/*
		Length (len): How many elements are currently in the slice.
		Capacity (cap): How many elements the underlying array can hold before it needs to grow.
		Why use make? Growing an array is expensive (it takes time to copy data). If you know you need to store 1,000 items, use make([]int, 0, 1000) to set the capacity upfront and save your CPU the extra work.
	*/

	var array []int32 = []int32{1, 2, 3}
	fmt.Printf("The length is %v with capacity %v", len(array), cap(array))
	array = append(array, 5)
	fmt.Printf("The length is %v with capacity %v", len(array), cap(array))

	var myMap map[string]uint8 = make(map[string]uint8)
	val, ok := myMap["Adam"]
	if ok {
		fmt.Println("Adam's age is", val)
	} else {
		fmt.Println("Adam is not in the map")
	}
	// Always use make or a literal: m := make(map[string]int).
	// The "Comma Ok" Idiom: This is a very specific Go pattern.
	// Since maps return a "zero value" if a key is missing, you use this to check if the key actually existed:
	// Go intentionally randomizes map iteration order. So it won't be the same twice!

	// Loops : There is no while or do-while. There is only for

	/*
		Loop TypeSyntax ExampleUse Case
		Classic	=> for i := 0; i < 10; i++ (Standard counting)
		"While" style => for condition { ... } (Loop until something is true)
		Infinitefor => { ... } (Loop forever (use break to exit))
		Range => for index, value := range slice (Iterating through collections)

	*/

}
