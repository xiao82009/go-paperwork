package main

import (
	"fmt"
	"os"
)

func main() {
	 
	for s, arg := range os.Args[0:] {
		fmt.Println(s,arg)
	}
	
}