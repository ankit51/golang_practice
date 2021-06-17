package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findian() {
	fmt.Printf("Enter a String: \n")
	in := bufio.NewReader(os.Stdin)
	x, _ := in.ReadString('\n')
	x = strings.ToLower(x[:len(x)-1])
	if strings.HasPrefix(x, "i") && strings.HasSuffix(x, "n") && strings.Contains(x, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}

func main() {
	findian()
}
