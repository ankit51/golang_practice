package main

import "fmt"

func trunc() {
	var flt float32
	fmt.Printf("Enter a floating point number:\n")
	fmt.Scan(&flt)
	fmt.Printf("%d", int(flt))
}

func main() {
	trunc()
}
