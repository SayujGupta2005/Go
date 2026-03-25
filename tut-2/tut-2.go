package main

import "fmt"

func main() {
	var intNum int = 64
	fmt.Println(intNum)

	var floatNum float32 = 312.3413456787543456789
	var floatNum2 float64 = 312.341345678754345678
	fmt.Println(floatNum)
	fmt.Println((floatNum2))
	// float32 has about 7 decimal digits of precision.
	// float64 has about 15 decimal digits of precision.
	//  In strings use  "" to write in one line, `` to write in multiple lines
	var a string = "sayuj"
	var b string = `nice
to
meet
you `
	fmt.Println(a)
	fmt.Println(b)
	println(len(a))

	var v = 1   // we dont need to define variable type if initialised the value then and there
	s := "name" // Adding : before = removes the necessity to add var keyword

}
