package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func slice() {
	sli := make([]int, 0, 3)
	var i string
	for true {
		fmt.Printf("Enter an integer or X to Exit:\n")
		fmt.Scan(&i)
		if strings.Contains(i, "X") {
			break
		}
		j, _ := strconv.Atoi(i)
		sli = append(sli, j)
		sort.Ints(sli)
		fmt.Println(sli)
	}

}

func main() {
	x := []int{4, 8, 5}
	y := -1
	for _, elt := range x {
		if elt > y {
			y = elt
		}
	}
	fmt.Print(y)
}
