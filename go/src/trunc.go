package main

import (
	"fmt"
	"strconv"
)

func trunc() {
	var x string
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&x)
	i, err := strconv.ParseFloat(x, 64)
	_ = err
	fmt.Printf(strconv.Itoa(int(i)) + "\n")
}
