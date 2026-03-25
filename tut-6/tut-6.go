package main

import "fmt"

type GasEngine struct {
	mpg     uint8
	gallons uint8
}

type EceEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e GasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e EceEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println(("Yes"))
	} else {
		fmt.Println("No")
	}
}

func main() {
	// 1. Literal initialization (named fields)
	var myEngine = GasEngine{mpg: 25, gallons: 10}
	fmt.Println(myEngine)
	// 2. Short-hand (order must match definition)
	engine2 := GasEngine{30, 15}
	fmt.Println(engine2)
	// 3. Zero-value (all fields become 0, "", or false)
	var emptyEngine GasEngine
	fmt.Println(emptyEngine)
}

// (e GasEngine) is the receiver
