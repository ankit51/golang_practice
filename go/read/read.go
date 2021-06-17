package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please enter the file to be read with full path and extension: ")
	in := bufio.NewReader(os.Stdin)
	x, _ := in.ReadString('\n')
	file, _ := os.Open(x[:len(x)-1])

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	type Person struct {
		fname string
		lname string
	}

	sli := []Person{}

	for _, eachline := range txtlines {
		split := strings.Split(eachline, " ")
		var p1 Person
		p1.fname = split[0]
		p1.lname = split[1]
		sli = append(sli, p1)
	}

	for _, v := range sli {
		fmt.Println(v.fname, v.lname)
	}

}
