package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main()  {
	fmt.Println("Please enter the name of the file:")
	var filename string
	_, err := fmt.Scanln(&filename)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	var names []Name
	for fileScanner.Scan() {
		nameArray := strings.Split(fileScanner.Text()," ")
		var name Name
		name.fname = nameArray[0]
		name.lname = nameArray[1]
		names = append(names, name)
	}
	fmt.Println()
	for i := 0; i < len(names); i++ {
		fmt.Printf("%s %s\n", names[i].fname, names[i].lname)
	}
}
