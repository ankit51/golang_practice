package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func makejson() {
	fmt.Printf("Please enter a name: ")
	name := bufio.NewReader(os.Stdin)
	namestr, _ := name.ReadString('\n')
	fmt.Printf("Please enter an address: ")
	addr := bufio.NewReader(os.Stdin)
	addrstr, _ := addr.ReadString('\n')
	idMap := map[string]string{
		namestr[:len(namestr)-1]: addrstr[:len(addrstr)-1]}
	barr, _ := json.Marshal(idMap)
	fmt.Println(string(barr))
}

func main() {
	makejson()
}
