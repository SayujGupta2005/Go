// Pointers
package main

import "fmt"

//When working with structs, Go provides "syntactic sugar" to make pointers easier to use.
// You don't have to write (*myStruct).Field; you can just write myStruct.Field
type Player struct {
	Health int
}

func (p *Player) TakeDamage(amount int) {
	p.Health -= amount
}

func main() {
	i := 4
	p := &i         // Or we could've wriiten var p *int32= new(int32)
	fmt.Println(&p) // Address of pointer p
	fmt.Println(p)
	fmt.Println(*p)
	*p = 3426
	fmt.Println(i)
	// p = 3456
	fmt.Println(p)
	s := new(int) //Allocates the memory for an int and returns the pointer
	fmt.Println(s)
	fmt.Println(*s) // value will be nil
	*s = 21435
	fmt.Println(s)
	fmt.Println(*s)

	var slice = []int32{1, 2, 3}
	var slice_copy = slice
	fmt.Println(slice)
	fmt.Println(slice_copy)
	slice_copy[2] = 4
	fmt.Println(slice) // We can see that on changing the variable in slice_copy,
	// the same index in slice also point to underlying data, rather than copying
	fmt.Println(slice_copy)

}

/*

Important Exceptions: Reference TypesIn Go, some types are "secretly" pointers.
You don't need to use the * symbol with them because their internal structure already points to an underlying data source:
Slices: A slice is a descriptor that points to an underlying array.Maps: Maps are pointers to a hash table.Channels: These are also pointers to an internal structure.

Type = Passing Behavior => Need * for mutation?
int, float, string, struct = Pass by Value (Copy) => Yes
Slice, Map, Channel = Pass by Reference (Internal pointer) => No

*/

/*
Go does not allow Pointer Arithmetic. * In C++, you can do p++ to move to the next memory address.
In Go, this is forbidden. This makes Go much safer because you can't accidentally wander into a "forbidden" area of memory and crash the system.
*/
