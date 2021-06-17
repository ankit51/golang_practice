package main

import (
	"fmt"
	"strings"
)

func findian() {
	var x string
	fmt.Printf("Enter a String: \n")
	fmt.Scan(&x)
	x = strings.ToLower(x)
	if strings.HasPrefix(x, "i") && strings.HasSuffix(x, "n") && strings.Contains(x, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}

func main() {
	findian()
}
