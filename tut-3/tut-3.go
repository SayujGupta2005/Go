package main

import "fmt"

func main() {
	numerator := 5
	denominator := 2
	printme(numerator)
	printme(denominator)
	var result, remainder int = intdiv(numerator, denominator)
	printme(result)
	printme(remainder)

}

func printme(printvalue any) {
	fmt.Println(printvalue)
}

func intdiv(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}

// || - bitwise or
// && - btiwise and
