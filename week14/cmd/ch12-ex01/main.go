package main

import (
	"fmt"
)

func main() {
	//LIFO
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")
	fmt.Println("main")
}
